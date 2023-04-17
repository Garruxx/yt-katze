package music

import "katze/src/models/external"

type Artist struct {
	Name       string                      `json:"name"`
	ID         string                      `json:"id"`
	Thumbnails []external.ThumbnailElement `json:"thumbnails,omitempty"`
}
