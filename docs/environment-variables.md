# go-ipws environment variables

## `LIBP2P_TCP_REUSEPORT` (`IPWS_REUSEPORT`)

go-ipws tries to reuse the same source port for all connections to improve NAT
traversal. If this is an issue, you can disable it by setting
`LIBP2P_TCP_REUSEPORT` to false.

This variable was previously `IPWS_REUSEPORT`.

Default: true

## `IPWS_PATH`

Sets the location of the IPWS repo (where the config, blocks, etc.
are stored).

Default: ~/.ipws

## `IPWS_LOGGING`

Sets the log level for go-ipws. It can be set to one of:

* `CRITICAL`
* `ERROR`
* `WARNING`
* `NOTICE`
* `INFO`
* `DEBUG`

Logging can also be configured (on a subsystem by subsystem basis) at runtime
with the `ipws log` command.

Default: `ERROR`

## `IPWS_LOGGING_FMT`

Sets the log message format. Can be one of:

* `color`
* `nocolor`

Default: `color`

## `GOLOG_FILE`

Sets the file to which go-ipws logs. By default, go-ipws logs to standard error.

## `GOLOG_TRACING_FILE`

Sets the file to which go-ipws sends tracing events. By default, tracing is
disabled.

This log can be read at runtime (without writing it to a file) using the `ipws
log tail` command.

Warning: Enabling tracing will likely affect performance.

## `IPWS_FUSE_DEBUG`

Enables fuse debug logging.

Default: false

## `YAMUX_DEBUG`

Enables debug logging for the yamux stream muxer.

Default: false

## `IPWS_FD_MAX`

Sets the file descriptor limit for go-ipws. If go-ipws fails to set the file
descriptor limit, it will log an error.

Defaults: 2048

## `IPWS_DIST_PATH`

URL from which go-ipws fetches repo migrations (when the daemon is launched with
the `--migrate` flag).

Default: https://ipfs.io/ipfs/$something (depends on the IPWS version)

## `LIBP2P_MUX_PREFS`

Tells go-ipws which multiplexers to use in which order.

Default: "/yamux/1.0.0 /mplex/6.7.0"
