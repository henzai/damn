package cmd

import "github.com/urfave/cli"

// Flags is
type Flags []cli.Flag

// NewFlags is
func NewFlags() *Flags {

	u := Flags{
		cli.StringFlag{
			Name:  "host, H",
			Usage: "target `host`",
			Value: "localhost",
		},
		cli.IntFlag{
			Name:  "port, P",
			Usage: "target `port`",
			Value: 5432,
		},
		cli.StringFlag{
			Name:  "dbname, D",
			Usage: "target `dbname`",
			Value: "postgres",
		},
		cli.StringFlag{
			Name:  "username, U",
			Usage: "PostgreSQL `username`",
			Value: "postgres",
		},
		cli.StringFlag{
			Name:  "password, W",
			Usage: "PostgreSQL `password`",
			Value: "admin",
		},
		cli.BoolFlag{
			Name:  "compress, c",
			Usage: "execute with compress (gunzip)",
		},
		cli.BoolFlag{
			Name:  "use-name, n",
			Usage: "use Container Name arg",
		},
		cli.StringFlag{
			Name:  "out, o",
			Usage: "`file`name",
		},
	}
	return &u
}
