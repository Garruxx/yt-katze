package models

type PlayerBrowseID struct {
	BrowseID string `json:"browseId"`
}

type PlayerInformation struct {
	Lyrics  PlayerBrowseID `json:"lyrics"`
	Related PlayerBrowseID `json:"related"`
}
