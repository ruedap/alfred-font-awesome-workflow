package main

import "testing"

func TestIconsFind_totalCount(t *testing.T) {
	ics := NewIcons()
	ai := ics.Find([]string{""})

	ex := 519
	ac := len(ai)
	if ex != ac {
		t.Errorf("failed to find icon: expected %v, got %v", ex, ac)
	}
}

func TestIconsFind_firstID(t *testing.T) {
	ics := NewIcons()
	ai := ics.Find([]string{""})

	ex := "adjust"
	ac := ai[0].ID
	if ex != ac {
		t.Errorf("failed to find icon: expected %v, got %v", ex, ac)
	}
}

func TestIconsFind_lastID(t *testing.T) {
	ics := NewIcons()
	ai := ics.Find([]string{""})

	ex := "youtube-square"
	ac := ai[len(ai)-1].ID
	if ex != ac {
		t.Errorf("failed to find icon: expected %v, got %v", ex, ac)
	}
}
