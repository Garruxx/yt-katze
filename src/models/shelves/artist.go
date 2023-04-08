package shelves

import (
	"katze/src/models"
	"katze/src/models/music"
)

type Artist struct {
	Name       string             `json:"name"`
	ID         string             `json:"id"`
	Thumbnails []models.Thumbnail `json:"thumbnails"`
	Tracks     []music.Song       `json:"tracks"`
	Albums     []Album            `json:"albums"`
	Singles    []music.Single     `json:"singles"`
	VisitorID  string             `json:"visitor_id,omitempty"`
}
