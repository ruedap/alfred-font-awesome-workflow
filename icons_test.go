package main

import "testing"

func TestIconsFind_countAll(t *testing.T) {
	fi := findIcons([]string{""})

	ex := 519
	ac := len(fi)
	if ex != ac {
		t.Errorf("failed to find icon: expected %v, got %v", ex, ac)
	}
}

func TestIconsFind_firstID(t *testing.T) {
	fi := findIcons([]string{""})

	ex := "adjust"
	ac := fi[0].ID
	if ex != ac {
		t.Errorf("failed to find icon: expected %v, got %v", ex, ac)
	}
}

func TestIconsFind_lastID(t *testing.T) {
	fi := findIcons([]string{""})

	ex := "youtube-square"
	ac := fi[len(fi)-1].ID
	if ex != ac {
		t.Errorf("failed to find icon: expected %v, got %v", ex, ac)
	}
}

func findIcons(terms []string) Icons {
	return NewIcons().Find(terms)
}
