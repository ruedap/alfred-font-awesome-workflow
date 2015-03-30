package main

import "strings"

type Icon struct {
	Name       string
	ID         string
	Unicode    string
	Created    string
	Filter     []string
	Aliases    []string
	Categories []string
}

func (ic *Icon) containID(terms []string) bool {
	for _, term := range terms {
		if !strings.Contains(ic.ID, term) {
			return false
		}
	}

	return true
}
