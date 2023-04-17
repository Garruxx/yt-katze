package artist

import "katze/src/models/external"

type CardItem struct {
	Thumbnails []external.ThumbnailElement
	Title      string
	Year       int
	BrowseID   string
	Params     string
	Explicit   bool
}
