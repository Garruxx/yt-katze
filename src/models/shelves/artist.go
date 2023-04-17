package shelves

import (
	"katze/src/models/artist"
	"katze/src/models/external"
)

type Artist struct {
	Name        string                      `json:"name"`
	Thumbnails  []external.ThumbnailElement `json:"thumbnails"`
	MusicList   artist.MusicList            `json:"musicList"`
	AlbumsList  artist.AlbumList            `json:"albumsList"`
	SinglesList artist.SingleList           `json:"singlesList"`
	VisitorID   string                      `json:"visitorID,omitempty"`
}
