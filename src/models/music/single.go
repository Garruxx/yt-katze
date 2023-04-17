package music

import "katze/src/models/external"

type Single struct {
	Title      string                      `json:"title"`
	Artists    []Artist                    `json:"artist"`
	ID         string                      `json:"id"`
	Year       int                         `json:"year,omitempty"`
	Thumbnails []external.ThumbnailElement `json:"thumbnails"`
	Explicit   bool                        `json:"explicit,omitempty"`
}
