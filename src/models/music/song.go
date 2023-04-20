package music

import "katze/src/models/external"

type Song struct {
	Title       string                      `json:"title"`
	ID          string                      `json:"id"`
	Artists     []Artist                    `json:"artistS"`
	AlbumTitle  *string                     `json:"album"`
	AlbumID     *string                     `json:"albumId"`
	Duration    *string                     `json:"duration,omitempty"`
	Explicit    *bool                       `json:"explicit,omitempty"`
	Thumbnails  []external.ThumbnailElement `json:"thumbnailS,omitempty"`
	TrackNumber *int                        `json:"trackNumber,omitempty"`
}
