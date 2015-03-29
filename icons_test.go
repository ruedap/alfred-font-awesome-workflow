package main

import "testing"

func TestIconsFind_countAll(t *testing.T) {
	fi := iconsFindHelper([]string{""})
	actual := len(fi)

	expected := 519
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestIconsFind_countZero(t *testing.T) {
	fi := iconsFindHelper([]string{"foo-bar-baz"})
	actual := len(fi)

	expected := 0
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestIconsFind_countOne(t *testing.T) {
	fi := iconsFindHelper([]string{"github-square"})
	actual := len(fi)

	expected := 1
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestIconsFind_countTwo(t *testing.T) {
	fi := iconsFindHelper([]string{"github-"})
	actual := len(fi)

	expected := 2
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestIconsFind_firstID(t *testing.T) {
	fi := iconsFindHelper([]string{""})
	actual := fi[0].ID

	expected := "adjust"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestIconsFind_lastID(t *testing.T) {
	fi := iconsFindHelper([]string{""})
	actual := fi[len(fi)-1].ID

	expected := "youtube-square"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestIconsFind_taxi(t *testing.T) {
	fi := iconsFindHelper([]string{"taxi"})
	actual := fi[0].Name

	expected := "Taxi"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}

	actual = fi[0].ID
	expected = "taxi"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}

	actual = fi[0].Unicode
	expected = "f1ba"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}

	actual = fi[0].Created
	expected = "4.1"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}

	actual = fi[0].Aliases[0]
	expected = "cab"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}

	actual = fi[0].Filter[0]
	expected = "vehicle"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}

	actual = fi[0].Categories[0]
	expected = "Web Application Icons"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestIconsFind_aliases(t *testing.T) {
	fi := iconsFindHelper([]string{"navicon"})
	actual := fi[0].ID

	expected := "bars"
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}

	if len(fi) != 1 {
		t.Errorf("expected %v to eq %v", len(fi), 1)
	}
}

func iconsFindHelper(terms []string) Icons {
	return NewIcons().Find(terms)
}
