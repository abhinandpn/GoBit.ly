package model

type Goly struct {
	Id       uint64
	ReDirect string
	Goly     string
	Clicked  uint64
	Random   string
}

type URLShortener struct {
    urls map[string]string
}
