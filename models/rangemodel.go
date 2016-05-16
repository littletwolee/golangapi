package models

type Rangemodel struct {
	Filename       string
	Filedata       []byte
	Contenttype    string
	Start          int64
	End            int64
}
