package main

import (
	"fmt"
	"html"

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
		cli.StringFlag{Name: "ref", Usage: "Convert to character reference"},
		cli.StringFlag{Name: "url", Usage: "Convert to URL of official site"},
	},
}

func doSearch(c *cli.Context) {
	terms := c.Args()

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

func doConvert(c *cli.Context) {
	name := c.String("name")
	code := c.String("code")
	ref := c.String("ref")
	url := c.String("url")
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
