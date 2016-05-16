package models

type Filemodel struct {
	Filename       string
	Filedata       []byte
	Contenttype    string
	Filetype       string
	Currentchunk   int
	Maxchunks      int
	Size           int64
}
