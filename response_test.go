package main

import (
	"reflect"
	"testing"
)

func TestResponse_newResponse(t *testing.T) {
	terms := []string{"Foo-Foo", "!BAR*BAR?", "バズ"}

	actual := newResponse(terms).Terms
	expected := []string{"foo-foo", "!bar*bar?", "バズ"}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestResponse_AddItem(t *testing.T) {
	r := newResponse([]string{})

	actual := r.AddItem(&responseItem{Title: "title-foo"}).Items
	expected := []responseItem{responseItem{Title: "title-foo"}}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestResponse_ToXML(t *testing.T) {
	r := newResponse([]string{})
	item := responseItem{
		Valid:    true,
		UID:      "f000-uid",
		Title:    "title-foo",
		Subtitle: "Subtitle foo.",
		Arg:      "arg-foo",
		Icon:     "./icons/title-foo.png",
		Unicode:  "f000-unicode",
	}
	r.AddItem(&item)

	actual, err := r.ToXML()
	if err != nil {
		t.Error("failed to convert to XML")
		return
	}

	expected := `<?xml version="1.0" encoding="UTF-8"?>
<items><item valid="true" arg="arg-foo" uid="f000-uid" unicode="f000-unicode"><title>title-foo</title><subtitle>Subtitle foo.</subtitle><icon>./icons/title-foo.png</icon></item></items>`
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestResponse_ToXML_Blank(t *testing.T) {
	r := newResponse([]string{})

	actual, err := r.ToXML()
	if err != nil {
		t.Error("failed to convert to XML")
		return
	}

	expected := `<?xml version="1.0" encoding="UTF-8"?>
<items></items>`
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}
