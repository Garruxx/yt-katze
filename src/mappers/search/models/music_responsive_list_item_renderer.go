package models

import (
	mappersModels "katze/src/mappers/models"
	"katze/src/models"
)

type MusicResponsiveListItemRenderer struct {
	Thumbnails  []models.Thumbnail
	FlexColumns []mappersModels.ColumnsItems
	PageType    string
	Explicit    bool
	ID          string
}
