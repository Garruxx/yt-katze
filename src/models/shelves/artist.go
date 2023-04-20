package shelves

import (
	"katze/src/models/artist"
	"katze/src/models/external"
)

type Artist struct {
	Name        string                      `json:"name"`
	Thumbnails  []external.ThumbnailElement `json:"thumbnails"`
	MusicsList  artist.MusicsList           `json:"musicsList"`
	AlbumsList  artist.AlbumsList           `json:"albumsList"`
	SinglesList artist.SinglesList          `json:"singlesList"`
	VisitorID   string                      `json:"visitorId,omitempty"`
}
