package data

import (
	mappersModels "katze/src/mappers/models"
	"katze/src/mappers/search/models"
	globalModels "katze/src/models"
)

var ITEM_SONG_VALID = models.MusicResponsiveListItemRenderer{
	PageType: "MUSIC_TYPE",
	ID:       "SongTestID",
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
					Text: "Song",
				},
				{
					Text:     "AlbumTest",
					BrowseID: "AlbumTestID",
					PageType: "MUSIC_PAGE_TYPE_ALBUM",
				},
				{
					Text:     "ArtistTest",
					BrowseID: "ArtistTestID",
					PageType: "MUSIC_PAGE_TYPE_ARTIST",
				},
				{
					Text: "22:22",
				},
			},
		},
	},
}

var ITEM_SONG_VALID_NO_EXPLICIT = models.MusicResponsiveListItemRenderer{

	PageType: "MUSIC_TYPE",
	ID:       "SongTestID",
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
					Text: "Song",
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
					Text:     "AlbumTest",
					BrowseID: "AlbumTestID",
					PageType: "MUSIC_PAGE_TYPE_ALBUM",
				},
				{
					Text: "22:22",
				},
			},
		},
	},
}

var ITEM_SONG_INVALID_EMPTY = models.MusicResponsiveListItemRenderer{}

var ITEM_SONG_INVALID_NO_ID = models.MusicResponsiveListItemRenderer{
	PageType: "MUSIC_TYPE",
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
					Text: "Song",
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
					Text: "22:22",
				},
			},
		},
	},
}

var ITEM_SONG_INVALID_NO_FLEX_COLUMNS = models.MusicResponsiveListItemRenderer{
	PageType: "MUSIC_TYPE",
	ID:       "SongTestID",
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

var ITEM_SONG_INVALID_DURATION = models.MusicResponsiveListItemRenderer{
	PageType: "MUSIC_TYPE",
	ID:       "SongTestID",
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
					Text: "Song",
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

var ITEM_SONG_INVALID_NO_PAGE_TYPE = models.MusicResponsiveListItemRenderer{
	ID:       "SongTestID",
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
					Text: "Song",
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
					Text: "22:22",
				},
			},
		},
	},
}
