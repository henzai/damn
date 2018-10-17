package cmd

import (
	"fmt"
	"os"
	"os/exec"

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
		cmd := Command{context: c}

		// 引数で得た文字列のコンテナidがあるか確認
		cmd.checkFirstArg()

		// 引数のコンテナid/名が起動中か確認
		cmd.isActiveContainer()

		d := &restoreScript{}

		// コマンド実行
		cmdstr := cmd.getCommand(d)
		fmt.Printf("%v\n", cmdstr)

		_, err := exec.Command("sh", "-c", cmdstr).Output()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return nil
	},
}

// Restore用
const restoreCmdCompress = "docker exec -i %v sh -c \"gunzip -c | psql -U %v -d %v\" < %v"
const restoreCmd = "docker exec -i %v psql -U %v -d %v < %v"

type restoreScript struct {
}

func (d *restoreScript) getCommand() string {
	return restoreCmd
}

func (d *restoreScript) getCommandCompress() string {
	return restoreCmdCompress
}
