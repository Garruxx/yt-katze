package artist

import "katze/src/models/external"

type CardItem struct {
	Thumbnails []external.ThumbnailElement `json:"thumbnails,omitempty"`
	Title      string                      `json:"title,omitempty"`
	Year       int                         `json:"year,omitempty"`
	BrowseID   string                      `json:"browseId,omitempty"`
	Params     string                      `json:"params,omitempty"`
	Explicit   bool                        `json:"explicit,omitempty"`
}
