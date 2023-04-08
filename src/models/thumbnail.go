package models

type Thumbnail struct {
	URL    string `json:"thumbnails"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
