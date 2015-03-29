package main

import (
	"encoding/xml"
	"strings"
)

type Response struct {
	Items   []ResponseItem
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

func NewResponse() *Response {
	return new(Response).init()
}

func (r *Response) init() *Response {
	r.Items = []ResponseItem{}
	return r
}

func (r *Response) AddItem(item *ResponseItem) {
	r.Items = append(r.Items, *item)
}

func (r *Response) GetXMLString() string {
	var x, err = xml.Marshal(r)
	if err != nil {
		panic(err) // FIXME
	}

	return xmlHeader + string(x)
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
