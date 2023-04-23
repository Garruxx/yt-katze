package models

type Request struct {
	Params        string
	Query         string
	BrowseID      string
	VideoID       string
	GoogVisitorID *string
	UrlQueries    string
	UrlPath       string
}
