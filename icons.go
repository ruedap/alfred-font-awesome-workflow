package main

import (
	"io/ioutil"
	"os"
	"sort"

	"gopkg.in/yaml.v2"
)

type IconsYaml struct {
	Icons Icons
}

type Icons []Icon

type Icon struct {
	Name       string
	ID         string
	Unicode    string
	Created    string
	Filter     []string
	Aliases    []string
	Categories []string
}

func NewIcons() Icons {
	path := os.Getenv("FAW_ICONS_YAML_PATH") // for testing
	if path == "" {
		path = "icons.yml" // default
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err) // FIXME
	}

	var y IconsYaml
	err = yaml.Unmarshal([]byte(b), &y)
	if err != nil {
		panic(err) // FIXME
	}

	return y.Icons.Sort()
}

func (ics Icons) Find(terms []string) Icons {
	var foundIcons Icons

	for _, icon := range ics {
		if ContainTerms(terms, icon.ID) {
			foundIcons = append(foundIcons, icon)
			continue
		}

		if ics.FindBy(icon.Aliases, terms) {
			foundIcons = append(foundIcons, icon)
			continue
		}
	}

	return foundIcons
}

func (ics Icons) FindBy(keywords, terms []string) bool {
	for _, k := range keywords {
		// FIXME: ContainTerms function dependency
		if ContainTerms(terms, k) {
			return true
		}
	}

	return false
}

// Len for sort
func (ics Icons) Len() int {
	return len(ics)
}

// Less for sort
func (ics Icons) Less(i, j int) bool {
	return ics[i].ID < ics[j].ID
}

// Swap for sort
func (ics Icons) Swap(i, j int) {
	ics[i], ics[j] = ics[j], ics[i]
}

func (ics Icons) Sort() Icons {
	sort.Sort(ics)
	return ics
}
