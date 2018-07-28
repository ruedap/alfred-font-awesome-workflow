package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

// Maintain tests for the `plus` icon
// https://github.com/ruedap/alfred-font-awesome-workflow/issues/74

func TestCommand_find(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cmd := &command{outStream: outStream, errStream: errStream}
	terms := []string{"plu"} // `plus`

	status := cmd.find(terms)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeOK)
	}

	actual := outStream.String()
  expected := `<?xml version="1.0" encoding="UTF-8"?>
<items><item valid="true" arg="f271" uid="f271"><title>calendar-plus</title><subtitle>Paste class name: fa-calendar-plus</subtitle><icon>./icons/calendar-plus.png</icon><unicode>f271</unicode></item><item valid="true" arg="f217" uid="f217"><title>cart-plus</title><subtitle>Paste class name: fa-cart-plus</subtitle><icon>./icons/cart-plus.png</icon><unicode>f217</unicode></item><item valid="true" arg="f4f1" uid="f4f1"><title>creative-commons-sampling-plus</title><subtitle>Paste class name: fa-creative-commons-sampling-plus</subtitle><icon>./icons/creative-commons-sampling-plus.png</icon><unicode>f4f1</unicode></item><item valid="true" arg="f2b3" uid="f2b3"><title>google-plus</title><subtitle>Paste class name: fa-google-plus</subtitle><icon>./icons/google-plus.png</icon><unicode>f2b3</unicode></item><item valid="true" arg="f0d5" uid="f0d5"><title>google-plus-g</title><subtitle>Paste class name: fa-google-plus-g</subtitle><icon>./icons/google-plus-g.png</icon><unicode>f0d5</unicode></item><item valid="true" arg="f0d4" uid="f0d4"><title>google-plus-square</title><subtitle>Paste class name: fa-google-plus-square</subtitle><icon>./icons/google-plus-square.png</icon><unicode>f0d4</unicode></item><item valid="true" arg="f1e6" uid="f1e6"><title>plug</title><subtitle>Paste class name: fa-plug</subtitle><icon>./icons/plug.png</icon><unicode>f1e6</unicode></item><item valid="true" arg="f067" uid="f067"><title>plus</title><subtitle>Paste class name: fa-plus</subtitle><icon>./icons/plus.png</icon><unicode>f067</unicode></item><item valid="true" arg="f055" uid="f055"><title>plus-circle</title><subtitle>Paste class name: fa-plus-circle</subtitle><icon>./icons/plus-circle.png</icon><unicode>f055</unicode></item><item valid="true" arg="f0fe" uid="f0fe"><title>plus-square</title><subtitle>Paste class name: fa-plus-square</subtitle><icon>./icons/plus-square.png</icon><unicode>f0fe</unicode></item><item valid="true" arg="f00e" uid="f00e"><title>search-plus</title><subtitle>Paste class name: fa-search-plus</subtitle><icon>./icons/search-plus.png</icon><unicode>f00e</unicode></item><item valid="true" arg="f234" uid="f234"><title>user-plus</title><subtitle>Paste class name: fa-user-plus</subtitle><icon>./icons/user-plus.png</icon><unicode>f234</unicode></item></items>`
	if !strings.Contains(actual, expected) {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestCommand_put_NameFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cmd := &command{outStream: outStream, errStream: errStream}
	flags := map[string]string{"name": "f067"}

	status := cmd.put(flags)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeOK)
	}

	actual := outStream.String()
	expected := "fa-plus"
	if !strings.Contains(actual, expected) {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestCommand_put_CodeFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cmd := &command{outStream: outStream, errStream: errStream}
	flags := map[string]string{"code": "f067"}

	status := cmd.put(flags)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeOK)
	}

	actual := outStream.String()
	expected := "f067"
	if !strings.Contains(actual, expected) {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestCommand_put_RefFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cmd := &command{outStream: outStream, errStream: errStream}
	flags := map[string]string{"ref": "f067"}

	status := cmd.put(flags)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeOK)
	}

	actual := outStream.String()
	expected := "plus"
	if !strings.Contains(actual, expected) {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestCommand_put_URLFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cmd := &command{outStream: outStream, errStream: errStream}
	flags := map[string]string{"url": "f067"}

	status := cmd.put(flags)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeOK)
	}

	actual := outStream.String()
	expected := "https://fontawesome.com/icons/plus"
	if !strings.Contains(actual, expected) {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestCommand_errorXML(t *testing.T) {
	err := errors.New("foo")
	actual := errorXML(err)
	expected := `<?xml version="1.0" encoding="UTF-8"?>
<items><item valid="false" arg="Error: foo" uid="error"><title>Error: foo</title><subtitle>Font Awesome Workflow</subtitle><icon>/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/AlertStopIcon.icns</icon></item></items>`
	if !strings.Contains(actual, expected) {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}
