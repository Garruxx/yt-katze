package music

import "katze/src/models/external"

type BestMatch struct {
	Type       string                      `json:"type,omitempty"`
	AlbumType  string                      `json:"albumType,omitempty"`
	Title      string                      `json:"title,omitempty"`
	ID         string                      `json:"id,omitempty"`
	Thumbnails []external.ThumbnailElement `json:"thumbnails,omitempty"`
	WatchID    string                      `json:"watchId,omitempty"`
	Artists    []Artist                    `json:"artists,omitempty"`
	AlbumTitle string                      `json:"album,omitempty"`
	AlbumID    string                      `json:"albumId,omitempty"`
	Duration   string                      `json:"duration,omitempty"`
	Explicit   bool                        `json:"explicit,omitempty"`
}
