package dataset

type Dataset struct {
	Id         string     `xml:"id,attr"`
	ColumnInfo ColumnInfo `xml:"ColumnInfo"`
	Rows       Rows       `xml:"Rows"`
}

type ColumnInfo struct {
	Columns []Column `xml:"Column"`
}

type Rows struct {
	Rows []Row `xml:"Row"`
}
