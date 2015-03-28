package main

import (
	"fmt"
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandSearch,
	commandConvert,
}

var commandSearch = cli.Command{
	Name:   "search",
	Usage:  "Search through Font Awesome icons",
	Action: doSearch,
}

var commandConvert = cli.Command{
	Name:   "convert",
	Usage:  "Convert icon name to each format",
	Action: doConvert,
	Flags: []cli.Flag{
		cli.StringFlag{Name: "name", Usage: "Convert to CSS class name"},
		cli.StringFlag{Name: "code", Usage: "Convert to character code"},
	},
}

func doSearch(c *cli.Context) {
	terms := c.Args()

	InitTerms(terms)

	r := NewResponse()
	icons := LoadIcons()

	for _, icon := range icons {
		if !MatchTerms(terms, icon.Id) {
			continue
		}

		r.AddItem(&ResponseItem{
			Valid:    true,
			Uid:      icon.Unicode,
			Title:    icon.Id,
			Subtitle: "Paste class name: fa-" + icon.Id,
			Arg:      icon.Id,
			Icon:     "./icons/fa-" + icon.Id + ".png",
			Unicode:  icon.Unicode,
		})
	}

	r.Print()
}

func doConvert(c *cli.Context) {
	name := c.String("name")
	code := c.String("code")

	if name != "" {
		fmt.Print("fa-" + name)
	}

	if code != "" {
		icons := LoadIcons()
		var u string
		for _, icon := range icons {
			if MatchTerms([]string{code}, icon.Id) {
				u = icon.Unicode
				break
			}
		}

		fmt.Print(u)
	}
}
