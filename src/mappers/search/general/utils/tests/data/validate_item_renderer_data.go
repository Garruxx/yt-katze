package data

import (
	"katze/src/mappers/models"
)

var browseID = "browseId"
var VALIDATE_ITEM_RENDERER_VALID_A_FLEX_COLUMN = models.MusicResponsiveListItemRenderer{

	PageType: "MUSIC_PAGE_TYPE_ARTIST",
	FlexColumns: []models.ColumnsItems{
		{
			Items: models.MultiColumn{
				{
					Text:     "ArtisNametest",
					BrowseID: &browseID,
				},
			},
		},
	},
	ID: "id",
}

var VALIDATE_ITEM_RENDERER_VALID_TWO_FLEX_COLUMN = models.MusicResponsiveListItemRenderer{

	PageType: "MUSIC_PAGE_TYPE_ARTIST",
	FlexColumns: []models.ColumnsItems{
		{
			Items: models.MultiColumn{
				{
					Text:     "ArtisNametest",
					BrowseID: &browseID,
				},
			},
		},
		{
			Items: models.MultiColumn{
				{
					Text:     "ArtisNametest",
					BrowseID: &browseID,
				},
			},
		},
	},
	ID: "id",
}

var VALIDATE_ITEM_RENDERER_INVALID_PAGE_TYPE = models.MusicResponsiveListItemRenderer{

	PageType: "ERROR",
	FlexColumns: []models.ColumnsItems{
		{
			Items: models.MultiColumn{
				{
					Text:     "ArtisNametest",
					BrowseID: &browseID,
				},
			},
		},
		{
			Items: models.MultiColumn{
				{
					Text:     "ArtisNametest",
					BrowseID: &browseID,
				},
			},
		},
	},
	ID: "id",
}

var VALIDATE_ITEM_RENDERER_INVALID_NO_FLEX_COLUMN = models.MusicResponsiveListItemRenderer{

	PageType:    "MUSIC_PAGE_TYPE_ARTIST",
	FlexColumns: []models.ColumnsItems{},
	ID:          "id",
}

var VALIDATE_ITEM_RENDERER_INVALID_NO_ID = models.MusicResponsiveListItemRenderer{

	PageType: "MUSIC_PAGE_TYPE_ARTIST",
	FlexColumns: []models.ColumnsItems{
		{
			Items: models.MultiColumn{
				{
					Text:     "ArtisNametest",
					BrowseID: &browseID,
				},
			},
		},
		{
			Items: models.MultiColumn{
				{
					Text:     "ArtisNametest",
					BrowseID: &browseID,
				},
			},
		},
	},
}
