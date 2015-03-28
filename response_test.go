package main

import (
	"testing"
)

func TestResponseInitTerms(t *testing.T) {
	terms := []string{"Foo-Foo", "BAR*BAR?", "バズ"}
	InitTerms(terms)
	if terms[0] != "foo-foo" {
		t.Error("failed to initialize terms")
	}

	if terms[1] != "bar*bar?" {
		t.Error("failed to initialize terms")
	}

	if terms[2] != "バズ" {
		t.Error("failed to initialize terms")
	}
}

func TestResponseContainTerms(t *testing.T) {
	if !ContainTerms([]string{"foo-bar"}, "FOO-Bar") {
		t.Error("failed to match terms")
	}

	if !ContainTerms([]string{"foo-bar"}, "1000foo-bar2000") {
		t.Error("failed to match terms")
	}

	if !ContainTerms([]string{"foo", "oops"}, "foops") {
		t.Error("failed to match terms")
	}

	if ContainTerms([]string{"foo-bar"}, "foo--bar") {
		t.Error("failed to match terms")
	}
}
