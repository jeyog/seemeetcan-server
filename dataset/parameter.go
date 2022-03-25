package dataset

type Parameter struct {
	Id    string `xml:"id,attr"`
	Value string `xml:",chardata"`
}
