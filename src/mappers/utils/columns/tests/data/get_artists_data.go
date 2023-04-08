package data

import "katze/src/models"

var GET_ARTIST_FLEX_COLUMS_WITH_TWO_ARTISTS = []models.FlexColumn{
	{
		Text:     "artistName",
		PageType: "MUSIC_PAGE_TYPE_ARTIST",
		BrowseID: "testID",
	},
	{
		Text:     "artistName",
		PageType: "MUSIC_PAGE_TYPE_ARTIST",
		BrowseID: "testID",
	},
	{
		Text:     "noArtistText",
		PageType: "MUSIC_PAGE_TYPE_ALBUM",
		BrowseID: "testID",
	},
}

var GET_ARTIST_FLEX_COLUMS_WITH_INVALID_DATA = []models.FlexColumn{
	{
		Text:     "artistName",
		PageType: "MUSIC_PAGE_TYPE_ARTIST",
		WatchID:  "testID",
	},
	{
		Text:     "artistName",
		PageType: "MUSIC_PAGE_TYPE_ARTIST",
		BrowseID: "testID",
	},
	{
		Text:     "noArtistText",
		PageType: "MUSIC_PAGE_TYPE_ALBUM",
		BrowseID: "testID",
	},
}

