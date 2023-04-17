package data

import "katze/src/models"

var GET_ALBUM_DATATEST_VALID = []models.FlexColumn{
	{
		Text:     "testText",
		BrowseID: &testID,
		PageType: &pageTypeAlbum,
	},
	{
		Text:     "testText1",
		BrowseID: &testID1,
		PageType: &pageTypeAlbum,
	},
	{
		Text:     "testText",
		BrowseID: &testID,
		PageType: &pageTypeArtist,
	},
}

var GET_ALBUM_DATATEST_EMPTY = []models.FlexColumn{}

var GET_ALBUM_DATATEST_INVALID_JUST_TEXT = []models.FlexColumn{
	{
		Text:     "testText",
		PageType: &pageTypeAlbum,
	},
}

var GET_ALBUM_DATATEST_INVALID_BROWSE_ID_EMPTY = []models.FlexColumn{
	{
		Text:     "testText",
		BrowseID: nil,
		PageType: &pageTypeAlbum,
	},
}

var GET_ALBUM_DATATEST_INVALID_WATCH_ID_NOT_EMPTY = []models.FlexColumn{
	{
		Text:     "testText",
		BrowseID: &testID,
		WatchID:  &testWatchID,
		PageType: &pageTypeAlbum,
	},
	{
		Text:     "testText1",
		BrowseID: &testID,
		PageType: &pageTypeAlbum1,
	},
}

var GET_ALBUM_DATATEST_NO_ALBUMS = []models.FlexColumn{
	{
		Text:     "testText",
		BrowseID: &testID,
		PageType: &pageTypeArtist,
	},
	{
		Text:     "testText1",
		BrowseID: &testID,
		PageType: &pageTypeArtist,
	},
}
