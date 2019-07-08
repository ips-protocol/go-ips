#!/bin/bash

src_dir=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
plist=io.ipfs.ipfs-daemon.plist
dest_dir="$HOME/Library/LaunchAgents"
IPWS_PATH="${IPWS_PATH:-$HOME/.ipws}"
escaped_ipws_path=$(echo $IPWS_PATH|sed 's/\//\\\//g')

IPWS_BIN=$(which ipws || echo ipws)
escaped_ipws_bin=$(echo $IPWS_BIN|sed 's/\//\\\//g')

mkdir -p "$dest_dir"

sed -e 's/{{IPWS_PATH}}/'"$escaped_ipws_path"'/g' \
  -e 's/{{IPWS_BIN}}/'"$escaped_ipws_bin"'/g' \
  "$src_dir/$plist" \
  > "$dest_dir/$plist"

launchctl list | grep ipws-daemon >/dev/null
if [ $? ]; then
  echo Unloading existing ipws-daemon
  launchctl unload "$dest_dir/$plist"
fi

echo Loading ipws-daemon
if (( `sw_vers -productVersion | cut -d'.' -f2` > 9 )); then
  sudo chown root "$dest_dir/$plist"
  sudo launchctl bootstrap system "$dest_dir/$plist"
else
  launchctl load "$dest_dir/$plist"
fi
launchctl list | grep ipws-daemon
