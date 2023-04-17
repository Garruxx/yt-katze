package music

import "katze/src/models/external"

type Song struct {
	Title       string                      `json:"title"`
	ID          string                      `json:"id"`
	Artists     []Artist                    `json:"artist"`
	AlbumTitle  *string                     `json:"album"`
	AlbumID     *string                     `json:"album_id"`
	Duration    *string                     `json:"duration,omitempty"`
	Explicit    *bool                       `json:"explicit,omitempty"`
	Thumbnails  []external.ThumbnailElement `json:"thumbnail,omitempty"`
	TrackNumber *int                        `json:"track_number,omitempty"`
}
