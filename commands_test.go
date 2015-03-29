package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCommandsExecFind(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cmd := &Command{outStream: outStream, errStream: errStream}
	terms := []string{"app"}

	status := cmd.execFind(terms)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeOK)
	}

	actual := outStream.String()
	expected := `<?xml version="1.0" encoding="UTF-8"?><items><item valid="true" arg="apple" uid="f179" unicode="f179"><title>apple</title><subtitle>Paste class name: fa-apple</subtitle><icon>./icons/fa-apple.png</icon></item><item valid="true" arg="whatsapp" uid="f232" unicode="f232"><title>whatsapp</title><subtitle>Paste class name: fa-whatsapp</subtitle><icon>./icons/fa-whatsapp.png</icon></item></items>`
	if !strings.Contains(actual, expected) {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}
