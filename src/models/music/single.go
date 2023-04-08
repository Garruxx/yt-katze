package music

import  "katze/src/models"

type Single struct {
	Title      string                     `json:"title"`
	Artists    []Artist                   `json:"artist"`
	ID         string                     `json:"id"`
	Year       int                        `json:"year,omitempty"`
	Thumbnails []models.Thumbnail `json:"thumbnails"`
	Explicit   bool                       `json:"explicit,omitempty"`
}
