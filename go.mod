module github.com/ipfs/go-ipfs

require (
	bazil.org/fuse v0.0.0-20180421153158-65cc252bf669
	github.com/Kubuxu/go-os-helper v0.0.1
	github.com/Kubuxu/gocovmerge v0.0.0-20161216165753-7ecaa51963cd
	github.com/allegro/bigcache v1.2.1 // indirect
	github.com/aristanetworks/goarista v0.0.0-20190628180533-8e7d5b18fe7a // indirect
	github.com/blang/semver v3.5.1+incompatible
	github.com/bren2010/proquint v0.0.0-20160323162903-38337c27106d
	github.com/cenkalti/backoff v2.1.1+incompatible
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/dustin/go-humanize v1.0.0
	github.com/elgris/jsondiff v0.0.0-20160530203242-765b5c24c302
	github.com/ethereum/go-ethereum v1.8.7
	github.com/fatih/color v1.7.0 // indirect
	github.com/fsnotify/fsnotify v1.4.7
	github.com/go-critic/go-critic v0.0.0-20181204210945-ee9bf5809ead // indirect
	github.com/gogo/protobuf v1.2.1
	github.com/golangci/golangci-lint v1.16.1-0.20190425135923-692dacb773b7
	github.com/hashicorp/go-multierror v1.0.0
	github.com/hashicorp/golang-lru v0.5.1
	github.com/hsanjuan/go-libp2p-http v0.0.2
	github.com/ipfs/dir-index-html v1.0.3
	github.com/ipfs/go-bitswap v0.0.7
	github.com/ipfs/go-block-format v0.0.2
	github.com/ipfs/go-blockservice v0.0.3
	github.com/ipfs/go-cid v0.0.2
	github.com/ipfs/go-cidutil v0.0.2
	github.com/ipfs/go-datastore v0.0.5
	github.com/ipfs/go-detect-race v0.0.1
	github.com/ipfs/go-ds-badger v0.0.4
	github.com/ipfs/go-ds-flatfs v0.0.2
	github.com/ipfs/go-ds-leveldb v0.0.2
	github.com/ipfs/go-ds-measure v0.0.1
	github.com/ipfs/go-fs-lock v0.0.1
	github.com/ipfs/go-ipfs-addr v0.0.1
	github.com/ipfs/go-ipfs-blockstore v0.0.1
	github.com/ipfs/go-ipfs-blocksutil v0.0.1
	github.com/ipfs/go-ipfs-chunker v0.0.1
	github.com/ipfs/go-ipfs-cmds v0.0.8
	github.com/ipfs/go-ipfs-ds-help v0.0.1
	github.com/ipfs/go-ipfs-exchange-interface v0.0.1
	github.com/ipfs/go-ipfs-exchange-offline v0.0.1
	github.com/ipfs/go-ipfs-files v0.0.3
	github.com/ipfs/go-ipfs-posinfo v0.0.1
	github.com/ipfs/go-ipfs-routing v0.0.1
	github.com/ipfs/go-ipfs-util v0.0.1
	github.com/ipfs/go-ipld-cbor v0.0.2
	github.com/ipfs/go-ipld-format v0.0.2
	github.com/ipfs/go-ipld-git v0.0.2
	github.com/ipfs/go-ipns v0.0.1
	github.com/ipfs/go-log v0.0.1
	github.com/ipfs/go-merkledag v0.0.3
	github.com/ipfs/go-metrics-interface v0.0.1
	github.com/ipfs/go-metrics-prometheus v0.0.2
	github.com/ipfs/go-mfs v0.0.7
	github.com/ipfs/go-path v0.0.4
	github.com/ipfs/go-unixfs v0.0.6
	github.com/ipfs/go-verifcid v0.0.1
	github.com/ipfs/hang-fds v0.0.1
	github.com/ipfs/iptb v1.4.0
	github.com/ipfs/iptb-plugins v0.0.2
	github.com/ipweb-group/go-ipws-config v0.9.1
	github.com/ipweb-group/interface-go-ipws-core v0.9.0
	github.com/jbenet/go-is-domain v1.0.2
	github.com/jbenet/go-random v0.0.0-20190219211222-123a90aedc0c
	github.com/jbenet/go-random-files v0.0.0-20190219210431-31b3f20ebded
	github.com/jbenet/go-temp-err-catcher v0.0.0-20150120210811-aac704a3f4f2
	github.com/jbenet/goprocess v0.1.3
	github.com/libp2p/go-libp2p v0.0.30
	github.com/libp2p/go-libp2p-autonat-svc v0.0.5
	github.com/libp2p/go-libp2p-circuit v0.0.9
	github.com/libp2p/go-libp2p-connmgr v0.0.6
	github.com/libp2p/go-libp2p-crypto v0.0.2
	github.com/libp2p/go-libp2p-host v0.0.3
	github.com/libp2p/go-libp2p-interface-connmgr v0.0.5
	github.com/libp2p/go-libp2p-kad-dht v0.0.13
	github.com/libp2p/go-libp2p-kbucket v0.1.1
	github.com/libp2p/go-libp2p-loggables v0.0.1
	github.com/libp2p/go-libp2p-metrics v0.0.1
	github.com/libp2p/go-libp2p-mplex v0.1.1
	github.com/libp2p/go-libp2p-net v0.0.2
	github.com/libp2p/go-libp2p-peer v0.1.1
	github.com/libp2p/go-libp2p-peerstore v0.0.6
	github.com/libp2p/go-libp2p-pnet v0.0.1
	github.com/libp2p/go-libp2p-protocol v0.0.1
	github.com/libp2p/go-libp2p-pubsub v0.0.3
	github.com/libp2p/go-libp2p-pubsub-router v0.0.3
	github.com/libp2p/go-libp2p-quic-transport v0.0.3
	github.com/libp2p/go-libp2p-record v0.0.1
	github.com/libp2p/go-libp2p-routing v0.0.1
	github.com/libp2p/go-libp2p-routing-helpers v0.0.2
	github.com/libp2p/go-libp2p-secio v0.0.3
	github.com/libp2p/go-libp2p-swarm v0.0.6
	github.com/libp2p/go-libp2p-tls v0.0.1
	github.com/libp2p/go-libp2p-yamux v0.1.3
	github.com/libp2p/go-maddr-filter v0.0.4
	github.com/libp2p/go-mplex v0.0.4 // indirect
	github.com/libp2p/go-stream-muxer v0.0.1
	github.com/libp2p/go-testutil v0.0.1
	github.com/libp2p/go-yamux v1.2.3 // indirect
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/mattn/go-runewidth v0.0.4 // indirect
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mr-tron/base58 v1.1.2
	github.com/multiformats/go-multiaddr v0.0.4
	github.com/multiformats/go-multiaddr-dns v0.0.2
	github.com/multiformats/go-multiaddr-net v0.0.1
	github.com/multiformats/go-multibase v0.0.1
	github.com/multiformats/go-multihash v0.0.5
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v0.9.3
	github.com/prometheus/procfs v0.0.0-20190519111021-9935e8e0588d // indirect
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/shirou/gopsutil v0.0.0-20190627142359-4c8b404ee5c5
	github.com/syndtr/goleveldb v1.0.0
	github.com/whyrusleeping/base32 v0.0.0-20170828182744-c30ac30633cc
	github.com/whyrusleeping/go-sysinfo v0.0.0-20190219211824-4a357d4b90b1
	github.com/whyrusleeping/multiaddr-filter v0.0.0-20160516205228-e903e4adabd7
	github.com/whyrusleeping/tar-utils v0.0.0-20180509141711-8c6c8ba81d5c
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/dig v1.7.0 // indirect
	go.uber.org/fx v1.9.0
	go.uber.org/goleak v0.10.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go4.org v0.0.0-20190313082347-94abd6928b1d // indirect
	golang.org/x/sys v0.0.0-20190614160838-b47fdc937951
	google.golang.org/appengine v1.4.0 // indirect
	gopkg.in/cheggaaa/pb.v1 v1.0.28
	gopkg.in/fatih/set.v0 v0.1.0 // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/karalabe/cookiejar.v2 v2.0.0-20150724131613-8dcd6a7f4951 // indirect
	gotest.tools/gotestsum v0.3.4
)

go 1.12

replace github.com/ethereum/go-ethereum v1.8.7 => github.com/ipweb-group/ipw v1.9.0
