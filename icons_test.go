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

func TestIconsFind_countZero(t *testing.T) {
	fi := findIcons([]string{"foo-bar-baz"})

	ex := 0
	ac := len(fi)
	if ex != ac {
		t.Errorf("failed to find icon: expected %v, got %v", ex, ac)
	}
}

func TestIconsFind_countOne(t *testing.T) {
	fi := findIcons([]string{"github-square"})

	ex := 1
	ac := len(fi)
	if ex != ac {
		t.Errorf("failed to find icon: expected %v, got %v", ex, ac)
	}
}

func TestIconsFind_countTwo(t *testing.T) {
	fi := findIcons([]string{"github-"})

	ex := 2
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

func TestIconsFind_taxi(t *testing.T) {
	fi := findIcons([]string{"taxi"})

	ex := "Taxi"
	ac := fi[0].Name
	if ex != ac {
		t.Errorf("failed to find icon: expected %v, got %v", ex, ac)
	}

	ex = "taxi"
	ac = fi[0].ID
	if ex != ac {
		t.Errorf("failed to find icon: expected %v, got %v", ex, ac)
	}

	ex = "f1ba"
	ac = fi[0].Unicode
	if ex != ac {
		t.Errorf("failed to find icon: expected %v, got %v", ex, ac)
	}

	ex = "4.1"
	ac = fi[0].Created
	if ex != ac {
		t.Errorf("failed to find icon: expected %v, got %v", ex, ac)
	}

	ex = "cab"
	ac = fi[0].Aliases[0]
	if ex != ac {
		t.Errorf("failed to find icon: expected %v, got %v", ex, ac)
	}

	ex = "vehicle"
	ac = fi[0].Filter[0]
	if ex != ac {
		t.Errorf("failed to find icon: expected %v, got %v", ex, ac)
	}

	ex = "Web Application Icons"
	ac = fi[0].Categories[0]
	if ex != ac {
		t.Errorf("failed to find icon: expected %v, got %v", ex, ac)
	}
}

func findIcons(terms []string) Icons {
	return NewIcons().Find(terms)
}
