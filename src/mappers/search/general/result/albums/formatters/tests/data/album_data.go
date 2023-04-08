package data

import (
	mappersModels "katze/src/mappers/models"
	"katze/src/mappers/search/models"
	globalModels "katze/src/models"
)

var ITEM_ALBUM_VALID_SINGLE = models.MusicResponsiveListItemRenderer{
	PageType: "MUSIC_PAGE_TYPE_ALBUM",
	ID:       "AlbumTestID",
	Explicit: true,
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
					Text: "TitleTest",
				},
			},
		},
		{
			Items: mappersModels.MultiColumn{
				{
					Text: "Single",
				},
				{
					Text:     "ArtistTest",
					BrowseID: "ArtistTestID",
					PageType: "MUSIC_PAGE_TYPE_ARTIST",
				},
				{
					Text:     "ArtistTest",
					BrowseID: "ArtistTestID",
					PageType: "MUSIC_PAGE_TYPE_ARTIST",
				},
				{
					Text: "2022",
				},
			},
		},
	},
}

var ITEM_ALBUM_VALID_EP = models.MusicResponsiveListItemRenderer{
	PageType: "MUSIC_PAGE_TYPE_ALBUM",
	ID:       "AlbumTestID",
	Explicit: true,
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
					Text: "TitleTest",
				},
			},
		},
		{
			Items: mappersModels.MultiColumn{
				{
					Text: "EP",
				},
				{
					Text:     "ArtistTest",
					BrowseID: "ArtistTestID",
					PageType: "MUSIC_PAGE_TYPE_ARTIST",
				},
				{
					Text:     "ArtistTest",
					BrowseID: "ArtistTestID",
					PageType: "MUSIC_PAGE_TYPE_ARTIST",
				},
				{
					Text: "2022",
				},
			},
		},
	},
}

var ITEM_ALBUM_VALID = models.MusicResponsiveListItemRenderer{
	PageType: "MUSIC_PAGE_TYPE_ALBUM",
	ID:       "AlbumTestID",
	Explicit: true,
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
					Text: "TitleTest",
				},
			},
		},
		{
			Items: mappersModels.MultiColumn{
				{
					Text: "Album",
				},
				{
					Text:     "ArtistTest",
					BrowseID: "ArtistTestID",
					PageType: "MUSIC_PAGE_TYPE_ARTIST",
				},
				{
					Text:     "ArtistTest",
					BrowseID: "ArtistTestID",
					PageType: "MUSIC_PAGE_TYPE_ARTIST",
				},
				{
					Text: "2022",
				},
			},
		},
	},
}

var ITEM_ALBUM_VALID_NO_EXPLICIT = models.MusicResponsiveListItemRenderer{

	PageType: "MUSIC_PAGE_TYPE_ALBUM",
	ID:       "AlbumTestID",
	Explicit: false,
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
					Text: "TitleTest",
				},
			},
		},
		{
			Items: mappersModels.MultiColumn{
				{
					Text: "Single",
				},
				{
					Text:     "ArtistTest",
					BrowseID: "ArtistTestID",
					PageType: "MUSIC_PAGE_TYPE_ARTIST",
				},
				{
					Text:     "ArtistTest",
					BrowseID: "ArtistTestID",
					PageType: "MUSIC_PAGE_TYPE_ARTIST",
				},
				{
					Text: "2022",
				},
			},
		},
	},
}

var ITEM_ALBUM_INVALID_EMPTY = models.MusicResponsiveListItemRenderer{}

var ITEM_ALBUM_INVALID_NO_ID = models.MusicResponsiveListItemRenderer{
	PageType: "MUSIC_PAGE_TYPE_ALBUM",
	Explicit: true,
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
					Text: "TitleTest",
				},
			},
		},
		{
			Items: mappersModels.MultiColumn{
				{
					Text: "Single",
				},
				{
					Text:     "ArtistTest",
					BrowseID: "ArtistTestID",
					PageType: "MUSIC_PAGE_TYPE_ARTIST",
				},
				{
					Text:     "ArtistTest",
					BrowseID: "ArtistTestID",
					PageType: "MUSIC_PAGE_TYPE_ARTIST",
				},
				{
					Text: "2022",
				},
			},
		},
	},
}

var ITEM_ALBUM_INVALID_NO_FLEX_COLUMNS = models.MusicResponsiveListItemRenderer{
	PageType: "MUSIC_PAGE_TYPE_ALBUM",
	ID:       "AlbumTestID",
	Explicit: true,
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
}

var ITEM_ALBUM_INVALID_YEAR = models.MusicResponsiveListItemRenderer{
	PageType: "MUSIC_PAGE_TYPE_ALBUM",
	ID:       "AlbumTestID",
	Explicit: true,
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
					Text: "TitleTest",
				},
			},
		},
		{
			Items: mappersModels.MultiColumn{
				{
					Text: "Single",
				},
				{
					Text:     "ArtistTest",
					BrowseID: "ArtistTestID",
					PageType: "MUSIC_PAGE_TYPE_ARTIST",
				},
				{
					Text:     "ArtistTest",
					BrowseID: "ArtistTestID",
					PageType: "MUSIC_PAGE_TYPE_ARTIST",
				},
				{
					Text: "2022a",
				},
			},
		},
	},
}

var ITEM_ALBUM_INVALID_NO_PAGE_TYPE = models.MusicResponsiveListItemRenderer{
	ID:       "AlbumTestID",
	Explicit: true,
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
					Text: "TitleTest",
				},
			},
		},
		{
			Items: mappersModels.MultiColumn{
				{
					Text: "Single",
				},
				{
					Text:     "ArtistTest",
					BrowseID: "ArtistTestID",
					PageType: "MUSIC_PAGE_TYPE_ARTIST",
				},
				{
					Text:     "ArtistTest",
					BrowseID: "ArtistTestID",
					PageType: "MUSIC_PAGE_TYPE_ARTIST",
				},
				{
					Text: "2022",
				},
			},
		},
	},
}
