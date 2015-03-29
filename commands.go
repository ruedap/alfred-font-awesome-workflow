package main

import (
	"fmt"
	"html"
	"io"
	"os"

	"github.com/codegangsta/cli"
)

// Exit codes.
const (
	ExitCodeOK int = 0

	ExitCodeError = 10 + iota
)

type Command struct {
	outStream, errStream io.Writer
}

var Commands = []cli.Command{
	commandFind,
	commandPut,
}

var commandFind = cli.Command{
	Name:  "find",
	Usage: "Search through Font Awesome icons",
	Action: func(c *cli.Context) {
		cmd := &Command{outStream: os.Stdout, errStream: os.Stderr}
		cmd.execFind(c.Args())
	},
}

var commandPut = cli.Command{
	Name:  "put",
	Usage: "Put the icon into each format",
	Action: func(c *cli.Context) {
		flags := map[string]string{
			"name": c.String("name"),
			"code": c.String("code"),
			"ref":  c.String("ref"),
			"url":  c.String("url"),
		}
		cmd := &Command{outStream: os.Stdout, errStream: os.Stderr}
		cmd.execPut(flags)
	},
	Flags: []cli.Flag{
		cli.StringFlag{Name: "name", Usage: "CSS class name"},
		cli.StringFlag{Name: "code", Usage: "Character code"},
		cli.StringFlag{Name: "ref", Usage: "Character reference"},
		cli.StringFlag{Name: "url", Usage: "URL of official site"},
	},
}

func (cmd *Command) execFind(terms []string) int {
	InitTerms(terms)

	r := NewResponse()
	icons := NewIcons().Find(terms)

	for _, icon := range icons {
		r.AddItem(&ResponseItem{
			Valid:    true,
			UID:      icon.Unicode,
			Title:    icon.ID,
			Subtitle: "Paste class name: fa-" + icon.ID,
			Arg:      icon.ID,
			Icon:     "./icons/fa-" + icon.ID + ".png",
			Unicode:  icon.Unicode,
		})
	}

	s := r.GetXMLString()
	_, err := fmt.Fprint(cmd.outStream, s)
	if s == "" || err != nil {
		return ExitCodeError
	}
	return ExitCodeOK
}

func (cmd *Command) execPut(flags map[string]string) int {
	name := flags["name"]
	code := flags["code"]
	ref := flags["ref"]
	url := flags["url"]
	ics := NewIcons()
	ost := cmd.outStream
	var err error

	if name != "" {
		_, err = fmt.Fprint(ost, "fa-"+name)
	}

	if code != "" {
		icons := ics.Find([]string{code})
		_, err = fmt.Fprint(ost, icons[0].Unicode)
	}

	if ref != "" {
		icons := ics.Find([]string{ref})
		str := html.UnescapeString("&#x" + icons[0].Unicode + ";")
		_, err = fmt.Fprint(ost, str)
	}

	if url != "" {
		_, err = fmt.Fprint(ost, "http://fontawesome.io/icon/"+url+"/")
	}

	if err != nil {
		return ExitCodeError
	}
	return ExitCodeOK
}
