package main

import "encoding/xml"

type response struct {
	Items   []responseItem
	XMLName struct{} `xml:"items"`
}

type responseItem struct {
	Valid    bool   `xml:"valid,attr"`
	Arg      string `xml:"arg,attr"`
	UID      string `xml:"uid,attr"`
	Unicode  string `xml:"unicode,attr"`
	Title    string `xml:"title"`
	Subtitle string `xml:"subtitle"`
	Icon     string `xml:"icon"`

	XMLName struct{} `xml:"item"`
}

func newResponse() *response {
	r := new(response)
	r.Items = []responseItem{}

	return r
}

func (r *response) addItem(item *responseItem) *response {
	r.Items = append(r.Items, *item)
	return r
}

func (r *response) toXML() (string, error) {
	var x, err = xml.Marshal(r)
	return xml.Header + string(x), err
}
