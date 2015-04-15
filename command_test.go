package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestCommand_find(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cmd := &command{outStream: outStream, errStream: errStream}
	terms := []string{"app"}

	status := cmd.find(terms)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeOK)
	}

	actual := outStream.String()
	expected := `<?xml version="1.0" encoding="UTF-8"?>
<items><item valid="true" arg="apple" uid="f179"><title>apple</title><subtitle>Paste class name: fa-apple</subtitle><icon>./icons/fa-apple.png</icon><unicode>f179</unicode></item><item valid="true" arg="whatsapp" uid="f232"><title>whatsapp</title><subtitle>Paste class name: fa-whatsapp</subtitle><icon>./icons/fa-whatsapp.png</icon><unicode>f232</unicode></item></items>`
	if !strings.Contains(actual, expected) {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestCommand_put_NameFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cmd := &command{outStream: outStream, errStream: errStream}
	flags := map[string]string{"name": "apple"}

	status := cmd.put(flags)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeOK)
	}

	actual := outStream.String()
	expected := "fa-apple"
	if !strings.Contains(actual, expected) {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestCommand_put_CodeFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cmd := &command{outStream: outStream, errStream: errStream}
	flags := map[string]string{"code": "apple"}

	status := cmd.put(flags)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeOK)
	}

	actual := outStream.String()
	expected := "f179"
	if !strings.Contains(actual, expected) {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestCommand_put_RefFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cmd := &command{outStream: outStream, errStream: errStream}
	flags := map[string]string{"ref": "apple"}

	status := cmd.put(flags)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeOK)
	}

	actual := outStream.String()
	expected := ""
	if !strings.Contains(actual, expected) {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestCommand_put_URLFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cmd := &command{outStream: outStream, errStream: errStream}
	flags := map[string]string{"url": "apple"}

	status := cmd.put(flags)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeOK)
	}

	actual := outStream.String()
	expected := "http://fontawesome.io/icon/apple/"
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
