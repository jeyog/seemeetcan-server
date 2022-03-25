package dataset

type Col struct {
	Id    string `xml:"id,attr"`
	Value string `xml:",chardata"`
}
