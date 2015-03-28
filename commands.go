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
	icons := searchIcon(terms)

	for _, icon := range icons {
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

	s := r.GetXMLString()
	fmt.Print(s)
}

func doConvert(c *cli.Context) {
	name := c.String("name")
	code := c.String("code")
	ref := c.String("ref")
	url := c.String("url")

	if name != "" {
		fmt.Print("fa-" + name)
	}

	if code != "" {
		icons := searchIcon([]string{code})
		fmt.Print(icons[0].Unicode)
	}

	if ref != "" {
		icons := searchIcon([]string{ref})
		str := html.UnescapeString("&#x" + icons[0].Unicode + ";")
		fmt.Print(str)
	}

	if url != "" {
		fmt.Print("http://fontawesome.io/icon/" + url + "/")
	}
}

func searchIcon(terms []string) []Icon {
	icons := LoadIcons()
	var r []Icon
	for _, icon := range icons {
		if ContainTerms(terms, icon.Id) {
			r = append(r, icon)
		}
	}

	return r
}
