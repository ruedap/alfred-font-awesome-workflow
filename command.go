package main

import (
	"errors"
	"fmt"
	"html"
	"io"
)

// Exit codes.
const (
	ExitCodeOK int = 0

	ExitCodeError = 10 + iota
)

type Command struct {
	outStream, errStream io.Writer
}

func (cmd *Command) Find(terms []string) int {
	r := NewResponse(terms)
	icons := NewIcons().Find(r.Terms)

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

	xml := r.ToXML()
	_, err := fmt.Fprint(cmd.outStream, xml)

	if xml == "" || err != nil {
		return ExitCodeError
	}
	return ExitCodeOK
}

func (cmd *Command) Put(flags map[string]string) int {
	name := flags["name"]
	code := flags["code"]
	ref := flags["ref"]
	url := flags["url"]
	ics := NewIcons()
	ost := cmd.outStream
	var err error

	switch {
	case name != "":
		_, err = fmt.Fprint(ost, "fa-"+name)
	case code != "":
		icons := ics.Find([]string{code})
		_, err = fmt.Fprint(ost, icons[0].Unicode)
	case ref != "":
		icons := ics.Find([]string{ref})
		str := html.UnescapeString("&#x" + icons[0].Unicode + ";")
		_, err = fmt.Fprint(ost, str)
	case url != "":
		_, err = fmt.Fprint(ost, "http://fontawesome.io/icon/"+url+"/")
	default:
		err = errors.New("invalid flag argument")
	}

	if err != nil {
		return ExitCodeError
	}
	return ExitCodeOK
}
