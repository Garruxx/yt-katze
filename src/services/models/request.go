package models

type Request struct {
	Params        string
	Query         string
	BrowseID      string
	GoogVisitorID *string
	UrlQueries    string
	UrlPath       string
}
