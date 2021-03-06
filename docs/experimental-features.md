# Experimental features of go-ipws

This document contains a list of experimental features in go-ipws.
These features, commands, and APIs aren't mature, and you shouldn't rely on them.
Once they reach maturity, there's going to be mention in the changelog and
release posts. If they don't reach maturity, the same applies, and their code is
removed.

Subscribe to https://github.com/ipfs/go-ipfs/issues/3397 to get updates.

When you add a new experimental feature to go-ipws, or change an experimental
feature, you MUST please make a PR updating this document, and link the PR in
the above issue.

- [ipws pubsub](#ipws-pubsub)
- [Client mode DHT routing](#client-mode-dht-routing)
- [go-multiplex stream muxer](#go-multiplex-stream-muxer)
- [Raw leaves for unixfs files](#raw-leaves-for-unixfs-files)
- [ipws filestore](#ipws-filestore)
- [ipws urlstore](#ipws-urlstore)
- [BadgerDB datastore](#badger-datastore)
- [Private Networks](#private-networks)
- [ipws p2p](#ipws-p2p)
- [p2p http proxy](#p2p-http-proxy)
- [Circuit Relay](#circuit-relay)
- [Plugins](#plugins)
- [Directory Sharding / HAMT](#directory-sharding--hamt)
- [IPNS PubSub](#ipns-pubsub)
- [QUIC](#quic)
- [AutoRelay](#autorelay)
- [TLS 1.3 Handshake](#tls-13-as-default-handshake-protocol)

---

## ipws pubsub

### State

experimental, default-disabled.

### In Version

0.4.5

### How to enable

run your daemon with the `--enable-pubsub-experiment` flag. Then use the
`ipws pubsub` commands.

### gossipsub

Gossipsub is a new, experimental routing protocol for pubsub that
should waste less bandwidth than floodsub, the current pubsub
protocol. It's backwards compatible with floodsub so enabling this
feature shouldn't break compatibility with existing IPWS nodes.

You can enable gossipsub via configuration:
`ipws config Pubsub.Router gossipsub`

### Message Signing

As of 0.4.18, go-ipws signs all pubsub messages by default. For now, it doesn't
*reject* unsigned messages but it will in the future.

You can turn off message signing (not recommended unless you're using a private
network) by running:
`ipws config Pubsub.DisableSigning true`

You can turn on strict signature verification (require that all messages be
signed) by running:
`ipws config Pubsub.StrictSignatureVerification true`

(this last option will be set to true by default and eventually removed entirely)

### Road to being a real feature
- [ ] Needs more people to use and report on how well it works
- [ ] Needs authenticated modes to be implemented
- [ ] needs performance analyses to be done

---

## Client mode DHT routing
Allows the dht to be run in a mode that doesn't serve requests to the network,
saving bandwidth.

### State
experimental.

### In Version
0.4.5

### How to enable
run your daemon with the `--routing=dhtclient` flag.

### Road to being a real feature
- [ ] Needs more people to use and report on how well it works.
- [ ] Needs analysis of effect it has on the network as a whole.

---

## go-multiplex stream muxer
Adds support for using the go-multiplex stream muxer alongside (or instead of)
yamux and spdy. This multiplexer is far simpler, and uses less memory and
bandwidth than the others, but is lacking on congestion control and backpressure
logic. It is available to try out and experiment with.

### State
Stable

### In Version
0.4.5

### How to enable

To make it the default stream muxer, set the environment variable
`LIBP2P_MUX_PREFS` as follows:
```
export LIBP2P_MUX_PREFS="/mplex/6.7.0 /yamux/1.0.0 /spdy/3.1.0"
```

---

## Raw Leaves for unixfs files
Allows files to be added with no formatting in the leaf nodes of the graph.

### State
experimental.

### In Version
master, 0.4.5

### How to enable
Use `--raw-leaves` flag when calling `ipws add`.

### Road to being a real feature
- [ ] Needs more people to use and report on how well it works.

---

## ipws filestore
Allows files to be added without duplicating the space they take up on disk.

### State
experimental.

### In Version
master, 0.4.7

### How to enable
Modify your ipws config:
```
ipws config --json Experimental.FilestoreEnabled true
```

Then restart your IPWS node to reload your config.

Finally, when adding files with ipws add, pass the --nocopy flag to use the
filestore instead of copying the files into your local IPWS repo.

### Road to being a real feature
- [ ] Needs more people to use and report on how well it works.
- [ ] Need to address error states and failure conditions
- [ ] Need to write docs on usage, advantages, disadvantages
- [ ] Need to merge utility commands to aid in maintenance and repair of filestore

---

## ipws urlstore
Allows ipws to retrieve blocks contents via a url instead of storing it in the datastore

### State
experimental.

### In Version
master, v0.4.17

### How to enable
Modify your ipws config:
```
ipws config --json Experimental.UrlstoreEnabled true
```

And then add a file at a specific URL using `ipws urlstore add <url>`

### Road to being a real feature
- [ ] Needs more people to use and report on how well it works.
- [ ] Need to address error states and failure conditions
- [ ] Need to write docs on usage, advantages, disadvantages
- [ ] Need to implement caching
- [ ] Need to add metrics to monitor performance

---

## Private Networks

Allows ipws to only connect to other peers who have a shared secret key.

### State
Experimental

### In Version
master, 0.4.7

### How to enable
Generate a pre-shared-key using [ipws-swarm-key-gen](https://github.com/Kubuxu/go-ipfs-swarm-key-gen)):
```
go get github.com/Kubuxu/go-ipfs-swarm-key-gen/ipfs-swarm-key-gen
ipws-swarm-key-gen > ~/.ipws/swarm.key
```

To join a given private network, get the key file from someone in the network
and save it to `~/.ipws/swarm.key` (If you are using a custom `$IPWS_PATH`, put
it in there instead).

When using this feature, you will not be able to connect to the default bootstrap
nodes (Since we aren't part of your private network) so you will need to set up
your own bootstrap nodes.

First, to prevent your node from even trying to connect to the default bootstrap nodes, run:
```bash
ipws bootstrap rm --all
```

Then add your own bootstrap peers with:
```bash
ipws bootstrap add <multiaddr>
```

For example:
```
ipws bootstrap add /ip4/104.236.76.40/tcp/4001/ipfs/QmSoLV4Bbm51jM9C4gDYZQ9Cy3U6aXMJDAbzgu2fzaDs64
```

Bootstrap nodes are no different from all other nodes in the network apart from
the function they serve.

To be extra cautious, You can also set the `LIBP2P_FORCE_PNET` environment
variable to `1` to force the usage of private networks. If no private network is
configured, the daemon will fail to start.

### Road to being a real feature
- [ ] Needs more people to use and report on how well it works
- [ ] More documentation

---

## ipws p2p

Allows tunneling of TCP connections through Libp2p streams. If you've ever used
port forwarding with SSH (the `-L` option in openssh), this feature is quite
similar.

### State

Experimental

### In Version

master, 0.4.10

### How to enable

The `p2p` command needs to be enabled in config:

```sh
> ipws config --json Experimental.Libp2pStreamMounting true
```

### How to use

**Netcat example:**

First, pick a protocol name for your application. Think of the protocol name as
a port number, just significantly more user-friendly. In this example, we're
going to use `/x/kickass/1.0`.

***Setup:***

1. A "server" node with peer ID `$SERVER_ID`
2. A "client" node.

***On the "server" node:***

First, start your application and have it listen for TCP connections on
port `$APP_PORT`.

Then, configure the p2p listener by running:

```sh
> ipws p2p listen /x/kickass/1.0 /ip4/127.0.0.1/tcp/$APP_PORT
```

This will configure IPWS to forward all incoming `/x/kickass/1.0` streams to
`127.0.0.1:$APP_PORT` (opening a new connection to `127.0.0.1:$APP_PORT` per
incoming stream.

***On the "client" node:***

First, configure the client p2p dialer, so that it forwards all inbound
connections on `127.0.0.1:SOME_PORT` to the server node listening
on `/x/kickass/1.0`.

```sh
> ipws p2p forward /x/kickass/1.0 /ip4/127.0.0.1/tcp/$SOME_PORT /ipfs/$SERVER_ID
```

Next, have your application open a connection to `127.0.0.1:$SOME_PORT`. This
connection will be forwarded to the service running on `127.0.0.1:$APP_PORT` on
the remote machine. You can test it with netcat:

***On "server" node:***
```sh
> nc -v -l -p $APP_PORT
```

***On "client" node:***
```sh
> nc -v 127.0.0.1 $SOME_PORT
```

You should now see that a connection has been established and be able to
exchange messages between netcat instances.

(note that depending on your netcat version you may need to drop the `-v` flag)

**SSH example**

**Setup:**

1. A "server" node with peer ID `$SERVER_ID` and running ssh server on the
   default port.
2. A "client" node.

_you can get `$SERVER_ID` by running `ipws id -f "<id>\n"`_

***First, on the "server" node:***

```sh
ipws p2p listen /x/ssh /ip4/127.0.0.1/tcp/22
```

***Then, on "client" node:***

```sh
ipws p2p forward /x/ssh /ip4/127.0.0.1/tcp/2222 /ipfs/$SERVER_ID
```

You should now be able to connect to your ssh server through a libp2p connection
with `ssh [user]@127.0.0.1 -p 2222`.


### Road to being a real feature
- [ ] Needs more people to use and report on how well it works / fits use cases
- [ ] More documentation
- [ ] Support other protocols (e.g, unix domain sockets, websockets, etc.)

---

## p2p http proxy

Allows proxying of HTTP requests over p2p streams. This allows serving any standard http app over p2p streams.

### State

Experimental

### In Version

master, 0.4.19

### How to enable

The `p2p` command needs to be enabled in config:

```sh
> ipws config --json Experimental.Libp2pStreamMounting true
```

On the client, the p2p http proxy needs to be enabled in the config:

```sh
> ipws config --json Experimental.P2pHttpProxy true
```

### How to use

**Netcat example:**

First, pick a protocol name for your application. Think of the protocol name as
a port number, just significantly more user-friendly. In this example, we're
going to use `/http`.

***Setup:***

1. A "server" node with peer ID `$SERVER_ID`
2. A "client" node.

***On the "server" node:***

First, start your application and have it listen for TCP connections on
port `$APP_PORT`.

Then, configure the p2p listener by running:

```sh
> ipws p2p listen --allow-custom-protocol /http /ip4/127.0.0.1/tcp/$APP_PORT
```

This will configure IPWS to forward all incoming `/http` streams to
`127.0.0.1:$APP_PORT` (opening a new connection to `127.0.0.1:$APP_PORT` per incoming stream.

***On the "client" node:***

Next, have your application make a http request to `127.0.0.1:8080/p2p/$SERVER_ID/http/$FORWARDED_PATH`. This
connection will be forwarded to the service running on `127.0.0.1:$APP_PORT` on
the remote machine (which needs to be a http server!) with path `$FORWARDED_PATH`. You can test it with netcat:

***On "server" node:***
```sh
> echo -e "HTTP/1.1 200\nContent-length: 11\n\nIPWS rocks!" | nc -l -p $APP_PORT
```

***On "client" node:***
```sh
> curl http://localhost:8080/p2p/$SERVER_ID/http/
```

You should now see the resulting http response: IPWS rocks!

### Custom protocol names
We also support use of protocol names of the form /x/$NAME/http where $NAME doesn't contain any "/"'s

### Road to being a real feature
- [ ] Needs p2p streams to graduate from experiments
- [ ] Needs more people to use and report on how well it works / fits use cases
- [ ] More documentation

---

## Circuit Relay

Allows peers to connect through an intermediate relay node when there
is no direct connectivity.

### State
Experimental

### In Version
master, 0.4.11

### How to enable

The relay transport is enabled by default, which allows peers to dial through
relay and listens for incoming relay connections. The transport can be disabled
by setting `Swarm.DisableRelay = true` in the configuration.

By default, peers don't act as intermediate nodes (relays). This can be enabled
by setting `Swarm.EnableRelayHop = true` in the configuration. Note that the
option needs to be set before online services are started to have an effect; an
already online node would have to be restarted.

### Basic Usage:

In order to connect peers QmA and QmB through a relay node QmRelay:

- Both peers should connect to the relay:
`ipws swarm connect /transport/address/ipfs/QmRelay`
- Peer QmA can then connect to peer QmB using the relay:
`ipws swarm connect /ipws/QmRelay/p2p-circuit/ipfs/QmB`

Peers can also connect with an unspecific relay address, which will
try to dial through known relays:
`ipws swarm connect /p2p-circuit/ipfs/QmB`

Peers can see their (unspecific) relay address in the output of
`ipws swarm addrs listen`

### Road to being a real feature

- [ ] Needs more people to use it and report on how well it works.
- [ ] Advertise relay addresses to the DHT for NATed or otherwise unreachable
      peers.
- [ ] Active relay discovery for specific relay address advertisement. We would
      like advertised relay addresses to designate specific relays for efficient
      dialing.
- [ ] Dialing priorities for relay addresses; arguably, relay addresses should
      have lower priority than direct dials.

## Plugins

### In Version
0.4.11

### State
Experimental

Plugins allow to add functionality without the need to recompile the daemon.

### Basic Usage:

See [Plugin docs](./plugins.md)

### Road to being a real feature

- [ ] Better support for platforms other than Linux
- [ ] More plugins and plugin types
- [ ] Feedback on stability

 ## Badger datastore

 ### In Version
 0.4.11

 Badger-ds is new datastore implementation based on
 https://github.com/dgraph-io/badger

 ### Basic Usage

 ```
 $ ipws init --profile=badgerds
 ```
 or install https://github.com/ipfs/ipfs-ds-convert/ and
 ```
 [BACKUP ~/.ipws]
 $ ipws config profile apply badgerds
 $ ipws-ds-convert convert
 ```

###

### Road to being a real feature

- [ ] Needs more testing
- [ ] Make sure there are no unknown major problems

## Directory Sharding / HAMT

### In Version
0.4.8

### State
Experimental

Allows to create directories with unlimited number of entries - currently
size of unixfs directories is limited by the maximum block size

### Basic Usage:

```
ipws config --json Experimental.ShardingEnabled true
```

### Road to being a real feature

- [ ] Make sure that objects that don't have to be sharded aren't
- [ ] Generalize sharding and define a new layer between IPLD and IPWS

---

## IPNS pubsub

### In Version

0.4.14

### State

Experimental, default-disabled.

Utilizes pubsub for publishing ipns records in real time.

When it is enabled:
- IPNS publishers push records to a name-specific pubsub topic,
  in addition to publishing to the DHT.
- IPNS resolvers subscribe to the name-specific topic on first
  resolution and receive subsequently published records through pubsub
  in real time. This makes subsequent resolutions instant, as they
  are resolved through the local cache. Note that the initial
  resolution still goes through the DHT, as there is no message
  history in pubsub.

Both the publisher and the resolver nodes need to have the feature enabled for it
to work effectively.

### How to enable

run your daemon with the `--enable-namesys-pubsub` flag; enables pubsub.

### Road to being a real feature

- [ ] Needs more people to use and report on how well it works
- [ ] Add a mechanism for last record distribution on subscription,
  so that we don't have to hit the DHT for the initial resolution.
  Alternatively, we could republish the last record periodically.



## QUIC

### In Version

0.4.18

### State

Experiment, disabled by default

### How to enable

Modify your ipws config:

```
ipws config --json Experimental.QUIC true
```

For listening on a QUIC address, add it the swarm addresses, e.g. `/ip4/0.0.0.0/udp/4001/quic`.


### Road to being a real feature

- [ ] The IETF QUIC specification needs to be finalised.
- [ ] Make sure QUIC connections work reliably
- [ ] Make sure QUIC connection offer equal or better performance than TCP connections on real world networks
- [ ] Finalize libp2p-TLS handshake spec.


## AutoRelay

### In Version

0.4.19-dev

### State

Experimental, disabled by default.

Automatically discovers relays and advertises relay addresses when the node is behind an impenetrable NAT.

### How to enable

Modify your ipws config:

```
ipws config --json Swarm.EnableAutoRelay true
```

Bootstrappers (and other public nodes) need to also enable the AutoNATService:
```
ipws config --json Swarm.EnableAutoNATService true
```

### Road to being a real feature

- [ ] needs testing


## TLS 1.3 as default handshake protocol

### State

Every go-ipws node (>=0.4.21) accepts secio and TLS 1.3 connections but prefers
secio over TLS when dialing. To prefer TLS when dialing, you'll have to enable
this feature.

### How to enable

Modify your ipws config:

```
ipws config --json Experimental.PreferTLS true
```

### Road to being a real feature

- [ ] needs testing
- [ ] needs adoption
