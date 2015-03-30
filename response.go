package main

import (
	"encoding/xml"
	"strings"
)

type Response struct {
	Items   []ResponseItem
	Terms   []string `xml:"-"`
	XMLName struct{} `xml:"items"`
}

type ResponseItem struct {
	Valid    bool   `xml:"valid,attr"`
	Arg      string `xml:"arg,attr"`
	UID      string `xml:"uid,attr"`
	Unicode  string `xml:"unicode,attr"`
	Title    string `xml:"title"`
	Subtitle string `xml:"subtitle"`
	Icon     string `xml:"icon"`

	XMLName struct{} `xml:"item"`
}

const xmlHeader = `<?xml version="1.0" encoding="UTF-8"?>`

func NewResponse(terms []string) *Response {
	r := new(Response)
	r.Items = []ResponseItem{}

	for i, t := range terms {
		terms[i] = strings.ToLower(t)
	}
	r.Terms = terms

	return r
}

func (r *Response) AddItem(item *ResponseItem) {
	r.Items = append(r.Items, *item)
}

func (r *Response) ToXML() string {
	var x, err = xml.Marshal(r)
	if err != nil {
		panic(err) // FIXME
	}

	return xmlHeader + string(x)
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
