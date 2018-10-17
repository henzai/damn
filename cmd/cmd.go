package cmd

import (
	"fmt"
	"os"

	"github.com/henzai/damn/docker"
	"github.com/urfave/cli"
)

type Command struct {
	context    *cli.Context
	id         string
	user       string
	db         string
	file       string
	isCompress bool
}

func (c *Command) checkFirstArg() string {
	arg := c.context.Args().First()
	if arg == "" {
		fmt.Println("引数がありません。")
		os.Exit(1)
	}
	return arg
}

func (c *Command) isActiveContainer() {
	firstArg := c.context.Args().First()
	var id string
	if c.context.Bool("use-name") {
		// コンテナ名を確認
		result, _ := docker.HasContainerByName(firstArg)
		if id != "" {
			fmt.Println("起動中のコンテナが存在しません。")
			os.Exit(1)
		}
		id = result
	} else {
		// コンテナidを確認
		result, _ := docker.HasContainerByID(firstArg)
		if id != "" {
			fmt.Println("起動中のコンテナが存在しません。")
			os.Exit(1)
		}
		id = result
	}
	fmt.Printf("target Container: %v\n", id)
	c.id = id
}

func (c *Command) getCommand(s scriptFormat) string {
	c.user = c.context.String("username")
	c.db = c.context.String("dbname")
	if i := c.context.String("out"); i != "" {
		c.file = i
	} else {
		c.file = c.db
	}
	c.isCompress = c.context.Bool("compress")
	if c.isCompress {
		return fmt.Sprintf(s.getCommandCompress(), c.id, c.user, c.db, c.file)
	}
	return fmt.Sprintf(s.getCommand(), c.id, c.user, c.db, c.file)

}

type scriptFormat interface {
	getCommand() string
	getCommandCompress() string
}
