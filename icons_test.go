package main

import "testing"

func TestIconsFind_totalCount(t *testing.T) {
	ics := NewIcons()

	ex := 519
	ac := len(ics.Find([]string{""}))
	if ex != ac {
		t.Errorf("failed to find icon: expected %v, got %v", ex, ac)
	}
}

func TestIconsFind_firstID(t *testing.T) {
	ics := NewIcons()

	ex := "adjust"
	ac := ics.Find([]string{""})[0].Id
	if ex != ac {
		t.Errorf("failed to find icon: expected %v, got %v", ex, ac)
	}
}
