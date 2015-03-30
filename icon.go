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

// TODO: Refactoring
func (ic *Icon) containAlias(terms []string) bool {
	if len(ic.Aliases) == 0 {
		return false
	}

	b := true
	for _, alias := range ic.Aliases {
		b = true
		for _, term := range terms {
			if !strings.Contains(alias, term) {
				b = false
			}
		}

		if b {
			break
		}
	}

	return b
}
