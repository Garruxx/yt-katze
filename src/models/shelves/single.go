package shelves

import (
	"katze/src/models"
	"katze/src/models/music"
)

type Single struct {
	Title      string             `json:"title"`
	Artists    []music.Artist     `json:"artist"`
	Year       int                `json:"year"`
	Thumbnails []models.Thumbnail `json:"thumbnails"`
	ID         string             `json:"id"`
	VisitorID  string             `json:"visitor_id,omitempty"`
}
