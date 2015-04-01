package main

import "testing"

func TestIcon_containID(t *testing.T) {
	ic := Icon{
		ID:      "caret-square-o-left",
		Aliases: []string{"foo-left", "bar-baz-left"},
	}

	terms := []string{"left", "square"}
	actual := ic.containID(terms)
	expected := true
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}

	terms = []string{"sq", "qua", "are"}
	actual = ic.containID(terms)
	expected = true
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}

	terms = []string{"left", "square", "foo"}
	actual = ic.containID(terms)
	expected = false
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestIcon_containAlias(t *testing.T) {
	ic := Icon{
		ID:      "caret-square-o-left",
		Aliases: []string{"foo-left", "bar-baz-left"},
	}

	terms := []string{"left", "baz"}
	actual := ic.containAlias(terms)
	expected := true
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}

	terms = []string{"for", "bar"}
	actual = ic.containAlias(terms)
	expected = false
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}

	terms = []string{"foobar"}
	actual = ic.containAlias(terms)
	expected = false
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}

	ic = Icon{
		ID:      "caret-square-o-left",
		Aliases: []string{},
	}

	terms = []string{"a"}
	actual = ic.containAlias(terms)
	expected = false
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}
