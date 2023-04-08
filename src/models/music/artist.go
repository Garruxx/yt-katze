package music

import "katze/src/models"

type Artist struct {
	Name       string             `json:"name"`
	ID         string             `json:"id"`
	Thumbnails []models.Thumbnail `json:"thumbnails,omitempty"`
}
