package main

import (
	"errors"
	"fmt"
	"html"
	"io"
	"strings"

	"github.com/ruedap/go-alfred"
)

// Exit codes.
const (
	ExitCodeOK int = 0

	ExitCodeError = 10 + iota
)

type command struct {
	outStream, errStream io.Writer
}

func (cmd *command) find(terms []string) int {
	for i, t := range terms {
		terms[i] = strings.ToLower(t)
	}

	r := alfred.NewResponse()
	icons := newIcons().find(terms)

	for _, icon := range icons {
		r.AddItem(&alfred.ResponseItem{
			Valid:    true,
			UID:      icon.Unicode,
			Title:    icon.ID,
			Subtitle: "Paste class name: fa-" + icon.ID,
			Arg:      icon.ID,
			Icon:     "./icons/fa-" + icon.ID + ".png",
			Extra: map[string]string{
				"Unicode": icon.Unicode,
			},
		})
	}

	xml, err := r.ToXML()
	if err != nil {
		return ExitCodeError
	}

	_, err = fmt.Fprint(cmd.outStream, xml)
	if err != nil {
		return ExitCodeError
	}

	return ExitCodeOK
}

func (cmd *command) put(flags map[string]string) int {
	name := flags["name"]
	code := flags["code"]
	ref := flags["ref"]
	url := flags["url"]
	ics := newIcons()
	ost := cmd.outStream
	var err error

	switch {
	case name != "":
		_, err = fmt.Fprint(ost, "fa-"+name)
	case code != "":
		icons := ics.find([]string{code})
		_, err = fmt.Fprint(ost, icons[0].Unicode)
	case ref != "":
		icons := ics.find([]string{ref})
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
