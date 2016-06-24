package main

import "github.com/mkideal/cli"

func main() {
	cli.Run(new(Catalog), func(ctx *cli.Context) error {
		c := ctx.Argv().(*Catalog)

		c.loadAssets()
		c.writeEnum()

		ctx.String("source")
		return nil
	})
}
