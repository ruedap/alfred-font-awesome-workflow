package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
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

func LoadIcons() []Icon {
	b, err := ioutil.ReadFile("icons.yml")
	if err != nil {
		panic(err) // FIXME
	}

	icons := Icons{}
	err = yaml.Unmarshal([]byte(b), &icons)
	if err != nil {
		panic(err) // FIXME
	}

	return icons.Icons
}
