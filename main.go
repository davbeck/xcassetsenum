package main

import "github.com/mkideal/cli"

type argT struct {
	cli.Helper
	CatalogPath   string `cli:"*c,catalog" usage:"The path to an xcassets catalog to process into a Swift enum. Required."`
	AccessControl string `cli:"a,access_control" usage:"The access to use for the generated enum, such as public, private, internal. Defaults to internal." dft:"internal"`
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)

		c := NewCatalog(argv.CatalogPath, argv.AccessControl)
		c.writeEnum()

		return nil
	})
}
