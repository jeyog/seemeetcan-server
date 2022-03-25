package dataset

type Column struct {
	Id   string `xml:"id,attr"`
	Type string `xml:"type,attr"`
	Size string `xml:"size,attr"`
}
