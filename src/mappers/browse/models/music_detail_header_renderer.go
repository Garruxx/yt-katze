package models

import (
	"katze/src/models/external"
	"katze/src/models/music"
)

type MusicDetailHeaderRenderer struct {
	Thumbnails []external.ThumbnailElement
	Title      string
	Artists    []music.Artist
	Year       int
	Duration   string
}
