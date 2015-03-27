package main

import (
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandSearch,
}

var commandSearch = cli.Command{
	Name:   "search",
	Usage:  "Search through Font Awesome icons",
	Action: doSearch,
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
			Uid:      icon.Id,
			Title:    icon.Id,
			Subtitle: "Paste class name: fa-" + icon.Id,
			Arg:      icon.Id,
			Icon:     "./icons/fa-" + icon.Id + ".png",
		})
	}

	r.Print()
}
