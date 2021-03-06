package commands

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	core "github.com/ipfs/go-ipfs/core"
	cmdenv "github.com/ipfs/go-ipfs/core/commands/cmdenv"

	cmds "github.com/ipfs/go-ipfs-cmds"
	ic "github.com/libp2p/go-libp2p-crypto"
	kb "github.com/libp2p/go-libp2p-kbucket"
	peer "github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	identify "github.com/libp2p/go-libp2p/p2p/protocol/identify"
)

const offlineIdErrorMessage = `'ipws id' currently cannot query information on remote
peers without a running daemon; we are working to fix this.
In the meantime, if you want to query remote peers using 'ipws id',
please run the daemon:

    ipws daemon &
    ipws id QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ
`

type IdOutput struct {
	ID              string
	PublicKey       string
	Addresses       []string
	AgentVersion    string
	ProtocolVersion string
	Beneficiary     string
}

const (
	formatOptionName = "format"
)

var IDCmd = &cmds.Command{
	Helptext: cmds.HelpText{
		Tagline: "Show ipws node id info.",
		ShortDescription: `
Prints out information about the specified peer.
If no peer is specified, prints out information for local peers.

'ipws id' supports the format option for output with the following keys:
<id> : The peers id.
<aver>: Agent version.
<pver>: Protocol version.
<bene>: Beneficiary.
<pubkey>: Public key.
<addrs>: Addresses (newline delimited).

EXAMPLE:

    ipws id Qmece2RkXhsKe5CRooNisBTh4SK119KrXXGmoK6V3kb8aH -f="<addrs>\n"
`,
	},
	Arguments: []cmds.Argument{
		cmds.StringArg("peerid", false, false, "Peer.ID of node to look up."),
	},
	Options: []cmds.Option{
		cmds.StringOption(formatOptionName, "f", "Optional output format."),
	},
	Run: func(req *cmds.Request, res cmds.ResponseEmitter, env cmds.Environment) error {
		n, err := cmdenv.GetNode(env)
		if err != nil {
			return err
		}

		var id peer.ID
		if len(req.Arguments) > 0 {
			var err error
			id, err = peer.IDB58Decode(req.Arguments[0])
			if err != nil {
				return fmt.Errorf("invalid peer id")
			}
		} else {
			id = n.Identity
		}

		if id == n.Identity {
			output, err := printSelf(n)
			if err != nil {
				return err
			}
			return cmds.EmitOnce(res, output)
		}

		// TODO handle offline mode with polymorphism instead of conditionals
		if !n.IsOnline {
			return errors.New(offlineIdErrorMessage)
		}

		p, err := n.Routing.FindPeer(req.Context, id)
		if err == kb.ErrLookupFailure {
			return errors.New(offlineIdErrorMessage)
		}
		if err != nil {
			return err
		}

		output, err := printPeer(n.Peerstore, p.ID)
		if err != nil {
			return err
		}
		return cmds.EmitOnce(res, output)
	},
	Encoders: cmds.EncoderMap{
		cmds.Text: cmds.MakeTypedEncoder(func(req *cmds.Request, w io.Writer, out *IdOutput) error {
			format, found := req.Options[formatOptionName].(string)
			if found {
				output := format
				output = strings.Replace(output, "<id>", out.ID, -1)
				output = strings.Replace(output, "<aver>", out.AgentVersion, -1)
				output = strings.Replace(output, "<pver>", out.ProtocolVersion, -1)
				output = strings.Replace(output, "<pubkey>", out.PublicKey, -1)
				output = strings.Replace(output, "<bene>", out.Beneficiary, -1)
				output = strings.Replace(output, "<addrs>", strings.Join(out.Addresses, "\n"), -1)
				output = strings.Replace(output, "\\n", "\n", -1)
				output = strings.Replace(output, "\\t", "\t", -1)
				fmt.Fprint(w, output)
			} else {
				marshaled, err := json.MarshalIndent(out, "", "\t")
				if err != nil {
					return err
				}
				marshaled = append(marshaled, byte('\n'))
				fmt.Fprintln(w, string(marshaled))
			}
			return nil
		}),
	},
	Type: IdOutput{},
}

func printPeer(ps pstore.Peerstore, p peer.ID) (interface{}, error) {
	if p == "" {
		return nil, errors.New("attempted to print nil peer")
	}

	info := new(IdOutput)
	info.ID = p.Pretty()

	if pk := ps.PubKey(p); pk != nil {
		pkb, err := ic.MarshalPublicKey(pk)
		if err != nil {
			return nil, err
		}
		info.PublicKey = base64.StdEncoding.EncodeToString(pkb)
	}

	for _, a := range ps.Addrs(p) {
		info.Addresses = append(info.Addresses, a.String())
	}

	// if v, err := ps.Get(p, "ProtocolVersion"); err == nil {
	// 	if vs, ok := v.(string); ok {
	// 		info.ProtocolVersion = vs
	// 	}
	// }
	info.ProtocolVersion = core.LibP2PVersion
	if v, err := ps.Get(p, "AgentVersion"); err == nil {
		if vs, ok := v.(string); ok {
			if c := strings.Count(vs, "/"); c == 3 {
				pos := strings.LastIndex(vs, "/")
				info.AgentVersion = vs[0:pos]
				info.Beneficiary = vs[(pos + 1):(strings.Count(vs, "") - 1)]
			} else {
				info.AgentVersion = vs
			}
		}
	}

	return info, nil
}

// printing self is special cased as we get values differently.
func printSelf(node *core.IpfsNode) (interface{}, error) {
	info := new(IdOutput)
	info.ID = node.Identity.Pretty()

	pk := node.PrivateKey.GetPublic()
	pkb, err := ic.MarshalPublicKey(pk)
	if err != nil {
		return nil, err
	}
	info.PublicKey = base64.StdEncoding.EncodeToString(pkb)

	if node.PeerHost != nil {
		for _, a := range node.PeerHost.Addrs() {
			s := a.String() + "/ipfs/" + info.ID
			info.Addresses = append(info.Addresses, s)
		}
	}
	// info.ProtocolVersion = identify.LibP2PVersion
	// info.AgentVersion = identify.ClientVersion
	info.ProtocolVersion = core.LibP2PVersion
	if c := strings.Count(identify.ClientVersion, "/"); c == 3 {
		pos := strings.LastIndex(identify.ClientVersion, "/")
		info.AgentVersion = identify.ClientVersion[0:pos]
	} else {
		info.AgentVersion = identify.ClientVersion
	}
	info.Beneficiary = node.Beneficiary
	return info, nil
}
