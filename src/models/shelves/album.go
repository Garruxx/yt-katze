package shelves

import (
	"katze/src/models/external"
	"katze/src/models/music"
)

type Album struct {
	Title      string                      `json:"title"`
	Artists    []music.Artist              `json:"artists"`
	Thumbnails []external.ThumbnailElement `json:"thumbnails"`
	Tracks     []music.Song                `json:"tracks"`
	Duration   string                      `json:"duration"`
	Year       int                         `json:"year,omitempty"`
	VisitorID  string                      `json:"visitor_id,omitempty"`
}
