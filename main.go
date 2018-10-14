package main

import (
	"log"
	"os"

	"github.com/henzai/damn/cmd"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()

	app.Name = "damn"
	app.Usage = "Damn!"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		cmd.Dump,
		cmd.Restore,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
