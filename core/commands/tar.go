package commands

import (
	"fmt"
	"io"

	"github.com/ipfs/go-ipfs/core/commands/cmdenv"
	"github.com/ipfs/go-ipfs/namesys/resolve"
	tar "github.com/ipfs/go-ipfs/tar"

	"github.com/ipfs/go-ipfs-cmds"
	dag "github.com/ipfs/go-merkledag"
	"github.com/ipfs/go-path"
)

var TarCmd = &cmds.Command{
	Helptext: cmds.HelpText{
		Tagline: "Utility functions for tar files in ipws.",
	},

	Subcommands: map[string]*cmds.Command{
		"add": tarAddCmd,
		"cat": tarCatCmd,
	},
}

var tarAddCmd = &cmds.Command{
	Helptext: cmds.HelpText{
		Tagline: "Import a tar file into ipws.",
		ShortDescription: `
'ipws tar add' will parse a tar file and create a merkledag structure to
represent it.
`,
	},

	Arguments: []cmds.Argument{
		cmds.FileArg("file", true, false, "Tar file to add.").EnableStdin(),
	},
	Run: func(req *cmds.Request, res cmds.ResponseEmitter, env cmds.Environment) error {
		nd, err := cmdenv.GetNode(env)
		if err != nil {
			return err
		}

		enc, err := cmdenv.GetCidEncoder(req)
		if err != nil {
			return err
		}

		it := req.Files.Entries()
		file, err := cmdenv.GetFileArg(it)
		if err != nil {
			return err
		}

		node, err := tar.ImportTar(req.Context, file, nd.DAG)
		if err != nil {
			return err
		}

		c := node.Cid()

		return cmds.EmitOnce(res, &AddEvent{
			Name: it.Name(),
			Hash: enc.Encode(c),
		})
	},
	Type: AddEvent{},
	Encoders: cmds.EncoderMap{
		cmds.Text: cmds.MakeTypedEncoder(func(req *cmds.Request, w io.Writer, out *AddEvent) error {
			fmt.Fprintln(w, out.Hash)
			return nil
		}),
	},
}

var tarCatCmd = &cmds.Command{
	Helptext: cmds.HelpText{
		Tagline: "Export a tar file from IPWS.",
		ShortDescription: `
'ipws tar cat' will export a tar file from a previously imported one in IPWS.
`,
	},

	Arguments: []cmds.Argument{
		cmds.StringArg("path", true, false, "ipws path of archive to export.").EnableStdin(),
	},
	Run: func(req *cmds.Request, res cmds.ResponseEmitter, env cmds.Environment) error {
		nd, err := cmdenv.GetNode(env)
		if err != nil {
			return err
		}

		p, err := path.ParsePath(req.Arguments[0])
		if err != nil {
			return err
		}

		root, err := resolve.Resolve(req.Context, nd.Namesys, nd.Resolver, p)
		if err != nil {
			return err
		}

		rootpb, ok := root.(*dag.ProtoNode)
		if !ok {
			return dag.ErrNotProtobuf
		}

		r, err := tar.ExportTar(req.Context, rootpb, nd.DAG)
		if err != nil {
			return err
		}

		return res.Emit(r)
	},
}
