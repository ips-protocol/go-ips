#!/usr/bin/env bash
#
# Copyright (c) 2019 Protocol Labs
# MIT Licensed; see the LICENSE file in this repository.
#

test_description="Test plugin loading"

. lib/test-lib.sh

test_init_ipfs

test_expect_success "ipfs id succeeds" '
  ipfs id
'

test_expect_success "make a bad plugin" '
  mkdir -p "$IPWS_PATH/plugins" &&
  echo foobar > "$IPWS_PATH/plugins/foo.so" &&
  chmod +x "$IPWS_PATH/plugins/foo.so"
'

test_expect_success "ipfs id fails due to a bad plugin" '
  test_expect_code 1 ipfs id
'

test_expect_success "cleanup bad plugin" '
  rm "$IPWS_PATH/plugins/foo.so"
'

test_done
