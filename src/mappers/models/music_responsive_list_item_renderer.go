package models

import "katze/src/models/external"

type MusicResponsiveListItemRenderer struct {
	Thumbnails  []external.ThumbnailElement
	FlexColumns []ColumnsItems
	PageType    string
	Explicit    bool
	ID          string
}
