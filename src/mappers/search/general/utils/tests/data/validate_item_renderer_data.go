package data

import (
	mappersModels "katze/src/mappers/models"
	searchModels "katze/src/mappers/search/models"
)

var VALIDATE_ITEM_RENDERER_VALID_A_FLEX_COLUMN = searchModels.MusicResponsiveListItemRenderer{

	PageType: "MUSIC_PAGE_TYPE_ARTIST",
	FlexColumns: []mappersModels.ColumnsItems{
		{
			Items: mappersModels.MultiColumn{
				{
					Text:     "ArtisNametest",
					BrowseID: "browseId",
				},
			},
		},
	},
	ID: "id",
}

var VALIDATE_ITEM_RENDERER_VALID_TWO_FLEX_COLUMN = searchModels.MusicResponsiveListItemRenderer{

	PageType: "MUSIC_PAGE_TYPE_ARTIST",
	FlexColumns: []mappersModels.ColumnsItems{
		{
			Items: mappersModels.MultiColumn{
				{
					Text:     "ArtisNametest",
					BrowseID: "browseId",
				},
			},
		},
		{
			Items: mappersModels.MultiColumn{
				{
					Text:     "ArtisNametest",
					BrowseID: "browseId",
				},
			},
		},
	},
	ID: "id",
}

var VALIDATE_ITEM_RENDERER_INVALID_PAGE_TYPE = searchModels.MusicResponsiveListItemRenderer{

	PageType: "ERROR",
	FlexColumns: []mappersModels.ColumnsItems{
		{
			Items: mappersModels.MultiColumn{
				{
					Text:     "ArtisNametest",
					BrowseID: "browseId",
				},
			},
		},
		{
			Items: mappersModels.MultiColumn{
				{
					Text:     "ArtisNametest",
					BrowseID: "browseId",
				},
			},
		},
	},
	ID: "id",
}

var VALIDATE_ITEM_RENDERER_INVALID_NO_FLEX_COLUMN = searchModels.MusicResponsiveListItemRenderer{

	PageType:    "MUSIC_PAGE_TYPE_ARTIST",
	FlexColumns: []mappersModels.ColumnsItems{},
	ID:          "id",
}

var VALIDATE_ITEM_RENDERER_INVALID_NO_ID = searchModels.MusicResponsiveListItemRenderer{

	PageType: "MUSIC_PAGE_TYPE_ARTIST",
	FlexColumns: []mappersModels.ColumnsItems{
		{
			Items: mappersModels.MultiColumn{
				{
					Text:     "ArtisNametest",
					BrowseID: "browseId",
				},
			},
		},
		{
			Items: mappersModels.MultiColumn{
				{
					Text:     "ArtisNametest",
					BrowseID: "browseId",
				},
			},
		},
	},
}
