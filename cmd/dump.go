package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

// Dump is
var Dump = cli.Command{
	Name:      "dump",
	Aliases:   []string{"d"},
	Usage:     "add a task to the list",
	ArgsUsage: "container-name",
	Flags:     *NewFlags(),
	Action: func(c *cli.Context) error {
		cmd := Command{context: c}

		// 引数で得た文字列のコンテナidがあるか確認
		cmd.checkFirstArg()

		// 引数のコンテナid/名が起動中か確認
		cmd.isActiveContainer()

		d := &dumpScript{}

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

// Dump用
const shCmdWithCompress = "docker exec -i %v pg_dump --no-owner -U %v %v | gzip -9 > %v"
const shCmd = "docker exec -i %v pg_dump --no-owner -U %v %v > %v"

type dumpScript struct {
}

func (d *dumpScript) getCommand() string {
	return shCmd
}

func (d *dumpScript) getCommandCompress() string {
	return shCmdWithCompress
}
