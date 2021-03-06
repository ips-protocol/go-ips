# iptb test framework
#
# Copyright (c) 2014, 2016 Jeromy Johnson, Christian Couder
# MIT Licensed; see the LICENSE file in this repository.

export IPTB_ROOT="$(pwd)/.iptb"

ipwsi() {
  dir="$1"
  shift
  IPWS_PATH="$IPTB_ROOT/testbeds/default/$dir" ipws "$@"
}

check_has_connection() {
  node="$1"
  ipwsi "$node" swarm peers >"swarm_peers_$node" &&
  grep "ipws" "swarm_peers_$node" >/dev/null
}

iptb() {
    if ! command iptb "$@"; then
        case "$1" in
            start|stop|connect)
                test_fsh command iptb logs
                ;;
        esac
        return 1
    fi
}

startup_cluster() {
  num_nodes="$1"
  shift
  other_args="$@"
  bound=$(expr "$num_nodes" - 1)

  if test -n "$other_args"; then
    test_expect_success "start up nodes with additional args" "
      iptb start -wait -- ${other_args[@]}
    "
  else
    test_expect_success "start up nodes" '
      iptb start -wait
    '
  fi

  test_expect_success "connect nodes to eachother" '
    iptb connect [1-$bound] 0
  '

  for i in $(test_seq 0 "$bound")
  do
    test_expect_success "node $i is connected" '
      check_has_connection "$i" ||
      test_fsh cat "swarm_peers_$i"
    '
  done
}

iptb_wait_stop() {
    while ! iptb run -- sh -c '! { test -e "$IPWS_PATH/repo.lock" && fuser -f "$IPWS_PATH/repo.lock" >/dev/null; }'; do
        go-sleep 10ms
    done
}
