package dataset

import "encoding/xml"

type XMain struct {
	XMLName    xml.Name   `xml:"Root"`
	XMLns      string     `xml:"xmlns,attr"`
	Parameters Parameters `xml:"Parameters"`
	Dataset    []Dataset  `xml:"Dataset"`
}

type Parameters struct {
	Parameters []Parameter `xml:"Parameter"`
}
