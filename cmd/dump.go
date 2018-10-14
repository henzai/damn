package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/henzai/damn/docker"
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
		// 引数で得た文字列のコンテナidがあるか確認
		arg := c.Args().First()
		if arg == "" {
			fmt.Println("引数がありません。")
			os.Exit(1)
		}

		// 引数のコンテナid/名が起動中か確認
		var id string
		if c.Bool("use-name") {
			// コンテナ名を確認
			result, _ := docker.HasContainerByName(arg)
			if id != "" {
				fmt.Println("起動中のコンテナが存在しません。")
				os.Exit(1)
			}
			id = result
		} else {
			// コンテナidを確認
			result, _ := docker.HasContainerByID(arg)
			if id != "" {
				fmt.Println("起動中のコンテナが存在しません。")
				os.Exit(1)
			}
			id = result
		}
		fmt.Printf("target Container: %v\n", id)

		// コマンド実行
		cmdstr := getCmd(c, id)
		fmt.Printf("%v\n", cmdstr)

		_, err := exec.Command("sh", "-c", cmdstr).Output()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return nil
	},
}

// コマンド実行用定数
const shCmdWithCompress = "docker exec -i %v pg_dump --no-owner -U %v %v | gzip -9 > %v.sql.gz"
const shCmd = "docker exec -i %v pg_dump --no-owner -U %v %v > %v.sql"

// ダンプ用コマンドを得る
func getCmd(c *cli.Context, containerID string) string {
	db := c.String("dbname")
	user := c.String("username")
	var file string
	if i := c.String("out"); i != "" {
		file = i
	} else {
		file = db
	}
	if c.Bool("compress") {
		return fmt.Sprintf(shCmdWithCompress, containerID, db, user, file)
	}
	return fmt.Sprintf(shCmd, containerID, db, user, file)
}
