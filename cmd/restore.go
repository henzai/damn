package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

// Restore is
var Restore = cli.Command{
	Name:      "restore",
	Aliases:   []string{"res"},
	Usage:     "restore Dump file",
	ArgsUsage: "container-name...",
	Flags:     *NewFlags(),
	Action: func(c *cli.Context) error {
		fmt.Println("restore")
		return nil
	},
}
