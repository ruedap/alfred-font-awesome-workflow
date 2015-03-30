package main

import "testing"

func TestResponse_NewResponse_Terms(t *testing.T) {
	terms := []string{"Foo-Foo", "BAR*BAR?", "バズ"}
	actual := NewResponse(terms).Terms

	expected := "foo-foo"
	if actual[0] != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}

	expected = "bar*bar?"
	if actual[1] != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}

	expected = "バズ"
	if actual[2] != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestResponse_AddItem(t *testing.T) {
	r := NewResponse([]string{})
	r.AddItem(&ResponseItem{Title: "title-foo"})

	if r.Items[0].Title != "title-foo" {
		t.Error("failed to add item to new response")
	}
}

func TestResponse_ToXML(t *testing.T) {
	r := NewResponse([]string{})
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

	actual := r.ToXML()

	expected := `<?xml version="1.0" encoding="UTF-8"?><items><item valid="true" arg="arg-foo" uid="f000-uid" unicode="f000-unicode"><title>title-foo</title><subtitle>Subtitle foo.</subtitle><icon>./icons/title-foo.png</icon></item></items>`
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}
