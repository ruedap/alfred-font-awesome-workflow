package main

import (
	"errors"
	"fmt"
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
			Arg:      icon.Unicode,
			Icon:     "./icons/" + icon.ID + ".png",
			Extra: map[string]string{
				"Unicode": icon.Unicode,
			},
		})
	}

	xml, err := r.ToXML()
	if err != nil {
		fmt.Fprint(cmd.outStream, errorXML(err))
		return ExitCodeError
	}

	_, err = fmt.Fprint(cmd.outStream, xml)
	if err != nil {
		fmt.Fprint(cmd.outStream, errorXML(err))
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
		icon := ics.findByUnicode(name)[0]
		_, err = fmt.Fprint(ost, "fa-"+icon.ID)
	case code != "":
		icon := ics.findByUnicode(code)[0]
		_, err = fmt.Fprint(ost, icon.Unicode)
	case ref != "":
		icon := ics.findByUnicode(ref)[0]
		_, err = fmt.Fprint(ost, icon.ID)
	case url != "":
		icon := ics.findByUnicode(url)[0]
		_, err = fmt.Fprint(ost, "https://fontawesome.com/icons/"+icon.ID)
	default:
		err = errors.New("invalid flag argument")
	}

	if err != nil {
		return ExitCodeError
	}
	return ExitCodeOK
}

func errorXML(err error) string {
	return alfred.ErrorXML(
		"Error: "+err.Error(),
		"Font Awesome Workflow",
		"Error: "+err.Error(),
	)
}
