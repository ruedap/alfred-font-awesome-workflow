package main

import (
	"encoding/xml"
	"strings"
)

type response struct {
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

func NewResponse(terms []string) *response {
	r := new(response)
	r.Items = []ResponseItem{}

	for i, t := range terms {
		terms[i] = strings.ToLower(t)
	}
	r.Terms = terms

	return r
}

func (r *response) AddItem(item *ResponseItem) *response {
	r.Items = append(r.Items, *item)
	return r
}

func (r *response) ToXML() (string, error) {
	var x, err = xml.Marshal(r)
	return xml.Header + string(x), err
}
