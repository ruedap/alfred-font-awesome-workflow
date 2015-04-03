package main

import (
	"io/ioutil"
	"os"
	"sort"

	"gopkg.in/yaml.v2"
)

// TODO: Refactoring
type IconsYaml struct {
	Icons Icons
}

type Icons []icon

func NewIcons() Icons {
	p := iconsYamlPath()
	b, _ := iconsReadYaml(p)      // FIXME: error handling
	y, _ := iconsUnmarshalYaml(b) // FIXME: error handling

	return y.Icons.Sort()
}

func iconsYamlPath() string {
	path := os.Getenv("FAW_ICONS_YAML_PATH") // for test env
	if path == "" {
		path = "icons.yml" // default
	}

	return path
}

func iconsReadYaml(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func iconsUnmarshalYaml(b []byte) (IconsYaml, error) {
	var y IconsYaml
	err := yaml.Unmarshal([]byte(b), &y)
	return y, err
}

func (ics Icons) Find(terms []string) Icons {
	var foundIcons Icons

	for _, ic := range ics {
		if ic.contains(terms) {
			foundIcons = append(foundIcons, ic)
		}
	}

	return foundIcons
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
