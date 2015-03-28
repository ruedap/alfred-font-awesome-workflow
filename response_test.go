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
