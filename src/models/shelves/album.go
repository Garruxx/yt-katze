package shelves

import (
	"katze/src/models"
	"katze/src/models/music"
)

type Album struct {
	Title      string             `json:"title"`
	Artist     string             `json:"artist"`
	ID         string             `json:"id"`
	Thumbnails []models.Thumbnail `json:"thumbnails"`
	Tracks     []music.Song       `json:"tracks"`
	Duration   string             `json:"duration"`
	Year       int                `json:"year,omitempty"`
	VisitorID  string             `json:"visitor_id,omitempty"`
}
