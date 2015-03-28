package main

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Icons struct {
	Icons []Icon
}

type Icon struct {
	Name       string
	Id         string
	Unicode    string
	Created    string
	Filter     []string
	Categories []string
}

func NewIcons() *Icons {
	return new(Icons).init()
}

func (ics *Icons) init() *Icons {
	path := os.Getenv("FAW_ICONS_YAML_PATH") // for testing
	if path == "" {
		path = "icons.yml" // default
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err) // FIXME
	}

	err = yaml.Unmarshal([]byte(b), &ics)
	if err != nil {
		panic(err) // FIXME
	}

	return ics
}

func (ics *Icons) Find(terms []string) []Icon {
	icons := NewIcons().Icons

	var r []Icon
	for _, icon := range icons {
		// FIXME: ContainTerms function dependency
		if ContainTerms(terms, icon.Id) {
			r = append(r, icon)
		}
	}

	return r
}
