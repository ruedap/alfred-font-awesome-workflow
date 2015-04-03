package main

import (
	"os"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandsFind,
	commandsPut,
}

var commandsFind = cli.Command{
	Name:  "find",
	Usage: "Search through Font Awesome icons",
	Action: func(c *cli.Context) {
		cmd := &command{outStream: os.Stdout, errStream: os.Stderr}
		cmd.Find(c.Args())
	},
}

var commandsPut = cli.Command{
	Name:  "put",
	Usage: "Put the icon into each format",
	Action: func(c *cli.Context) {
		flags := map[string]string{
			"name": c.String("name"),
			"code": c.String("code"),
			"ref":  c.String("ref"),
			"url":  c.String("url"),
		}
		cmd := &command{outStream: os.Stdout, errStream: os.Stderr}
		cmd.Put(flags)
	},
	Flags: []cli.Flag{
		cli.StringFlag{Name: "name", Usage: "CSS class name"},
		cli.StringFlag{Name: "code", Usage: "Character code"},
		cli.StringFlag{Name: "ref", Usage: "Character reference"},
		cli.StringFlag{Name: "url", Usage: "URL of official site"},
	},
}
