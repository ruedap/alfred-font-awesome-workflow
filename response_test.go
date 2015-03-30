package main

import "testing"

func TestResponse_AddItem(t *testing.T) {
	r := NewResponse()
	r.AddItem(&ResponseItem{Title: "title-foo"})

	if r.Items[0].Title != "title-foo" {
		t.Error("failed to add item to new response")
	}
}

func TestResponse_GetXMLString(t *testing.T) {
	r := NewResponse()
	item := ResponseItem{
		Valid:    true,
		UID:      "f000-uid",
		Title:    "title-foo",
		Subtitle: "Subtitle foo.",
		Arg:      "arg-foo",
		Icon:     "./icons/title-foo.png",
		Unicode:  "f000-unicode",
	}
	r.AddItem(&item)

	actual := r.GetXMLString()

	expected := `<?xml version="1.0" encoding="UTF-8"?><items><item valid="true" arg="arg-foo" uid="f000-uid" unicode="f000-unicode"><title>title-foo</title><subtitle>Subtitle foo.</subtitle><icon>./icons/title-foo.png</icon></item></items>`
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestResponse_InitTerms(t *testing.T) {
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

func TestResponse_ContainTerms(t *testing.T) {
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
