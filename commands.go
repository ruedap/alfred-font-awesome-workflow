package main

import (
	"fmt"
	"html"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandFind,
	commandPut,
}

var commandFind = cli.Command{
	Name:  "find",
	Usage: "Search through Font Awesome icons",
	Action: func(c *cli.Context) {
		commandFindExec(c.Args())
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
		commandPutExec(flags)
	},
	Flags: []cli.Flag{
		cli.StringFlag{Name: "name", Usage: "CSS class name"},
		cli.StringFlag{Name: "code", Usage: "Character code"},
		cli.StringFlag{Name: "ref", Usage: "Character reference"},
		cli.StringFlag{Name: "url", Usage: "URL of official site"},
	},
}

func commandFindExec(terms []string) {
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
	fmt.Print(s)
}

func commandPutExec(flags map[string]string) {
	name := flags["name"]
	code := flags["code"]
	ref := flags["ref"]
	url := flags["url"]
	ics := NewIcons()

	if name != "" {
		fmt.Print("fa-" + name)
	}

	if code != "" {
		icons := ics.Find([]string{code})
		fmt.Print(icons[0].Unicode)
	}

	if ref != "" {
		icons := ics.Find([]string{ref})
		str := html.UnescapeString("&#x" + icons[0].Unicode + ";")
		fmt.Print(str)
	}

	if url != "" {
		fmt.Print("http://fontawesome.io/icon/" + url + "/")
	}
}
