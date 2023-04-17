package models

import "katze/src/models/external"

type CarouselShelfRenderer struct {
	Title    string
	BrowseID string
	Params   string
	Contents []external.MusicCarouselShelfRendererContent
}
