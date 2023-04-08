package music

import "katze/src/models"

type Album struct {
	Title      string              `json:"title"`
	ID         string              `json:"id"`
	Artists    *[]Artist           `json:"artist,omitempty"`
	Thumbnails *[]models.Thumbnail `json:"thumbnails,omitempty"`
	Single     *bool               `json:"single,omitempty"`
	EP         *bool               `json:"ep,omitempty"`
	Year       *int                `json:"year,omitempty"`
	Explicit   *bool               `json:"explicit,omitempty"`
}
