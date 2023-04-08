package data

import (
	mappersModels "katze/src/mappers/models"
	"katze/src/mappers/search/models"
	globalModels "katze/src/models"
)

var ITEM_ARTIST_VALID = models.MusicResponsiveListItemRenderer{
	PageType: "MUSIC_PAGE_TYPE_ARTIST",
	ID:       "ArtistIdTest",
	Thumbnails: []globalModels.Thumbnail{
		{
			URL:    "example.com",
			Width:  100,
			Height: 100,
		},
		{
			URL:    "example.com",
			Width:  100,
			Height: 100,
		},
	},
	FlexColumns: []mappersModels.ColumnsItems{
		{
			Items: mappersModels.MultiColumn{
				{
					Text: "ArtistNameTest",
				},
			},
		},
		{
			Items: mappersModels.MultiColumn{
				{
					Text: "example",
				},
				{
					Text: "example",
				},
			},
		},
	},
}

var ITEM_ARTIST_INVALID_EMPTY = models.MusicResponsiveListItemRenderer{}

var ITEM_ARTIST_INVALID_PAGE_TYPE = models.MusicResponsiveListItemRenderer{
	PageType: "MUSIC_PAGE_INVALID",
	ID:       "AlbumTestID",
	Thumbnails: []globalModels.Thumbnail{
		{
			URL:    "example.com",
			Width:  100,
			Height: 100,
		},
		{
			URL:    "example.com",
			Width:  100,
			Height: 100,
		},
	},
	FlexColumns: []mappersModels.ColumnsItems{
		{
			Items: mappersModels.MultiColumn{
				{
					Text: "ArtistNameTest",
				},
			},
		},
		{
			Items: mappersModels.MultiColumn{
				{
					Text: "example",
				},
				{
					Text: "example",
				},
			},
		},
	},
}

var ITEM_ARTIST_INVALID_FLEX_COLUMN_LENGTH = models.MusicResponsiveListItemRenderer{
	PageType: "MUSIC_PAGE_TYPE_ARTIST",
	ID:       "AlbumTestID",
	Thumbnails: []globalModels.Thumbnail{
		{
			URL:    "example.com",
			Width:  100,
			Height: 100,
		},
		{
			URL:    "example.com",
			Width:  100,
			Height: 100,
		},
	},
	FlexColumns: []mappersModels.ColumnsItems{
		{
			Items: mappersModels.MultiColumn{
				{
					Text: "ArtistNameTest",
				},
			},
		},
	},
}

var ITEM_ARTIST_INVALID_NO_ID = models.MusicResponsiveListItemRenderer{
	PageType: "MUSIC_PAGE_TYPE_ARTIST",
	Thumbnails: []globalModels.Thumbnail{
		{
			URL:    "example.com",
			Width:  100,
			Height: 100,
		},
		{
			URL:    "example.com",
			Width:  100,
			Height: 100,
		},
	},
	FlexColumns: []mappersModels.ColumnsItems{
		{
			Items: mappersModels.MultiColumn{
				{
					Text: "ArtistNameTest",
				},
			},
		},
		{
			Items: mappersModels.MultiColumn{
				{
					Text: "example",
				},
				{
					Text: "example",
				},
			},
		},
	},
}
