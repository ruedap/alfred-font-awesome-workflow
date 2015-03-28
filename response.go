package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Response struct {
	Items   []ResponseItem
	XMLName struct{} `xml:"items"`
}

type ResponseItem struct {
	Valid    bool   `xml:"valid,attr"`
	Arg      string `xml:"arg,attr"`
	Uid      string `xml:"uid,attr"`
	Title    string `xml:"title"`
	Subtitle string `xml:"subtitle"`
	Icon     string `xml:"icon"`
	Unicode  string `xml:"unicode"`

	XMLName struct{} `xml:"item"`
}

const xmlHeader = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"

func NewResponse() *Response {
	return new(Response).Init()
}

func (r *Response) Init() *Response {
	r.Items = []ResponseItem{}
	return r
}

func (r *Response) AddItem(item *ResponseItem) {
	r.Items = append(r.Items, *item)
}

func (r *Response) Print() {
	var x, _ = xml.Marshal(r)
	fmt.Print(xmlHeader, string(x))
}

func InitTerms(terms []string) {
	for i, t := range terms {
		terms[i] = strings.ToLower(t)
	}
}

func ContainTerms(terms []string, name string) bool {
	n := strings.ToLower(name)

	for _, t := range terms {
		if !strings.Contains(n, t) {
			return false
		}
	}

	return true
}
