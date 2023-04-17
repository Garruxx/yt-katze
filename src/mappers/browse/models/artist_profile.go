package models

import (
	"katze/src/models/external"
)

type ArtistProfile struct {
	Name               string `json:"name"`
	Thumbnails         []external.ThumbnailElement
	Carousels          []external.MusicCarouselShelfRenderer
	MusicShelfRenderer external.TentacledMusicShelfRenderer
	VisitorID          string `json:"visitorId"`
}
