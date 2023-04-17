package data

import "katze/src/models"

var GET_ARTIST_FLEX_COLUMS_WITH_TWO_ARTISTS = []models.FlexColumn{
	{
		Text:     "artistName",
		PageType: &pageTypeArtist,
		BrowseID: &testID,
	},
	{
		Text:     "artistName",
		PageType: &pageTypeArtist,
		BrowseID: &testID,
	},
	{
		Text:     "noArtistText",
		PageType: &pageTypeAlbum,
		BrowseID: &testID,
	},
}

var GET_ARTIST_FLEX_COLUMS_WITH_INVALID_DATA = []models.FlexColumn{
	{
		Text:     "artistName",
		PageType: &pageTypeArtist,
		WatchID:  &testID,
	},
	{
		Text:     "artistName",
		PageType: &pageTypeArtist,
		BrowseID: &testID,
	},
	{
		Text:     "noArtistText",
		PageType: &pageTypeAlbum,
		BrowseID: &testID,
	},
}
