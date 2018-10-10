package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "damn"
	app.Usage = "Damn!"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:    "dump",
			Aliases: []string{"d"},
			Usage:   "add a task to the list",
			Action: func(c *cli.Context) error {
				fmt.Println("added task: ", c.Args().First())
				fmt.Println(c.Bool("c"))
				return nil
			},
			ArgsUsage: "container-name...",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "d, dbname",
					Usage: "PostgreSQL Database",
					Value: "postgres",
				},
				cli.StringFlag{
					Name:  "U, username",
					Usage: "PostgreSQL User",
					Value: "postgres",
				},
				cli.StringFlag{
					Name:  "W, password",
					Usage: "PostgreSQL password",
					Value: "admin",
				},
				cli.BoolFlag{
					Name:  "c, compress",
					Usage: "gunzip dump file",
				},
			},
		},
		{
			Name:    "restore",
			Aliases: []string{"res"},
			Usage:   "restore Dump file",
			Action: func(c *cli.Context) error {
				fmt.Println("restore")
				return nil
			},
			ArgsUsage: "container-name...",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "d, dbname",
					Usage: "PostgreSQL Database",
					Value: "postgres",
				},
				cli.StringFlag{
					Name:  "U, username",
					Usage: "PostgreSQL User",
					Value: "postgres",
				},
				cli.StringFlag{
					Name:  "W, password",
					Usage: "PostgreSQL password",
					Value: "admin",
				},
				cli.BoolFlag{
					Name:  "c, compress",
					Usage: "gunzip dump file",
				},
			},
		},
	}

	app.Run(os.Args)
}
