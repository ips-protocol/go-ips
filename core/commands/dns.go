package commands

import (
	"fmt"
	"io"

	ncmd "github.com/ipfs/go-ipfs/core/commands/name"
	namesys "github.com/ipfs/go-ipfs/namesys"
	nsopts "github.com/ipweb-group/interface-go-ipws-core/options/namesys"

	cmds "github.com/ipfs/go-ipfs-cmds"
)

const (
	dnsRecursiveOptionName = "recursive"
)

var DNSCmd = &cmds.Command{
	Helptext: cmds.HelpText{
		Tagline: "Resolve DNS links.",
		ShortDescription: `
Multihashes are hard to remember, but domain names are usually easy to
remember.  To create memorable aliases for multihashes, DNS TXT
records can point to other DNS links, IPWS objects, IPNS keys, etc.
This command resolves those links to the referenced object.
`,
		LongDescription: `
Multihashes are hard to remember, but domain names are usually easy to
remember.  To create memorable aliases for multihashes, DNS TXT
records can point to other DNS links, IPWS objects, IPNS keys, etc.
This command resolves those links to the referenced object.

Note: This command can only recursively resolve DNS links,
it will fail to recursively resolve through IPNS keys etc.
For general-purpose recursive resolution, use ipws name resolve -r.

For example, with this DNS TXT record:

	> dig +short TXT _dnslink.ipfs.io
	dnslink=/ipfs/QmRzTuh2Lpuz7Gr39stNr6mTFdqAghsZec1JoUnfySUzcy

The resolver will give:

	> ipws dns ipfs.io
	/ipfs/QmRzTuh2Lpuz7Gr39stNr6mTFdqAghsZec1JoUnfySUzcy

The resolver can recursively resolve:

	> dig +short TXT recursive.ipfs.io
	dnslink=/ipns/ipfs.io
	> ipws dns -r recursive.ipfs.io
	/ipfs/QmRzTuh2Lpuz7Gr39stNr6mTFdqAghsZec1JoUnfySUzcy
`,
	},

	Arguments: []cmds.Argument{
		cmds.StringArg("domain-name", true, false, "The domain-name name to resolve.").EnableStdin(),
	},
	Options: []cmds.Option{
		cmds.BoolOption(dnsRecursiveOptionName, "r", "Resolve until the result is not a DNS link.").WithDefault(true),
	},
	Run: func(req *cmds.Request, res cmds.ResponseEmitter, env cmds.Environment) error {
		recursive, _ := req.Options[dnsRecursiveOptionName].(bool)
		name := req.Arguments[0]
		resolver := namesys.NewDNSResolver()

		var ropts []nsopts.ResolveOpt
		if !recursive {
			ropts = append(ropts, nsopts.Depth(1))
		}

		output, err := resolver.Resolve(req.Context, name, ropts...)
		if err != nil && (recursive || err != namesys.ErrResolveRecursion) {
			return err
		}
		return cmds.EmitOnce(res, &ncmd.ResolvedPath{Path: output})
	},
	Encoders: cmds.EncoderMap{
		cmds.Text: cmds.MakeTypedEncoder(func(req *cmds.Request, w io.Writer, out *ncmd.ResolvedPath) error {
			fmt.Fprintln(w, out.Path.String())
			return nil
		}),
	},
	Type: ncmd.ResolvedPath{},
}
