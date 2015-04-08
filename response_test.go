package main

import (
	"reflect"
	"testing"
)

func TestResponse_newResponse(t *testing.T) {
	actual := newResponse()
	expected := new(response)
	expected.Items = []responseItem{}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestResponse_addItem(t *testing.T) {
	r := newResponse()

	actual := r.addItem(&responseItem{Title: "title-foo"}).Items
	expected := []responseItem{responseItem{Title: "title-foo"}}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestResponse_toXML(t *testing.T) {
	r := newResponse()
	item := responseItem{
		Valid:    true,
		UID:      "f000-uid",
		Title:    "title-foo",
		Subtitle: "Subtitle foo.",
		Arg:      "arg-foo",
		Icon:     "./icons/title-foo.png",
		Unicode:  "f000-unicode",
	}
	r.addItem(&item)

	actual, err := r.toXML()
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

func TestResponse_toXML_Blank(t *testing.T) {
	r := newResponse()

	actual, err := r.toXML()
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
