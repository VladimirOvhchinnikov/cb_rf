package main

import (
	"bytes"
	"encoding/xml"

	"golang.org/x/net/html/charset"
)

type ValCurs struct {
	Date   string   `xml:"Date,attr"`
	Name   string   `xml:"name,attr"`
	Valute []Valute `xml:"Valute"`
}

type Valute struct {
	ID        string `xml:"ID,attr"`
	NumCode   string `xml:"NumCode"`
	CharCode  string `xml:"CharCode"`
	Nominal   int    `xml:"Nominal"`
	Name      string `xml:"Name"`
	Value     string `xml:"Value"`
	VunitRate string `xml:"VunitRate"`
}

type DateRate struct {
	Date string
	Rate float64
}

func Converter(chunks [][]byte) ([]ValCurs, error) {
	res := make([]ValCurs, len(chunks))
	for i, raw := range chunks {
		dec := xml.NewDecoder(bytes.NewReader(raw))
		dec.CharsetReader = charset.NewReaderLabel
		if err := dec.Decode(&res[i]); err != nil {
			return nil, err
		}
	}
	return res, nil
}
