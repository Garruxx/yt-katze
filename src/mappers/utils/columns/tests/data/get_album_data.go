package data

import "katze/src/models"

var GET_ALBUM_DATATEST_VALID = []models.FlexColumn{
	{
		Text:     "testText",
		BrowseID: "testID",
		PageType: "MUSIC_PAGE_TYPE_ALBUM",
	},
	{
		Text:     "testText1",
		BrowseID: "testID1",
		PageType: "MUSIC_PAGE_TYPE_ALBUM",
	},
	{
		Text:     "testText",
		BrowseID: "testID",
		PageType: "MUSIC_PAGE_TYPE_ARTIST",
	},
}

var GET_ALBUM_DATATEST_EMPTY = []models.FlexColumn{}

var GET_ALBUM_DATATEST_INVALID_JUST_TEXT = []models.FlexColumn{
	{
		Text:     "testText",
		PageType: "MUSIC_PAGE_TYPE_ALBUM",
	},
}

var GET_ALBUM_DATATEST_INVALID_BROWSE_ID_EMPTY = []models.FlexColumn{
	{
		Text:     "testText",
		BrowseID: "",
		PageType: "MUSIC_PAGE_TYPE_ALBUM",
	},
}

var GET_ALBUM_DATATEST_INVALID_WATCH_ID_NOT_EMPTY = []models.FlexColumn{
	{
		Text:     "testText",
		BrowseID: "testID",
		WatchID:  "testWatchID",
		PageType: "MUSIC_PAGE_TYPE_ALBUM",
	},
	{
		Text:     "testText1",
		BrowseID: "testID1",
		PageType: "MUSIC_PAGE_TYPE_ALBUM1",
	},
}

var GET_ALBUM_DATATEST_NO_ALBUMS = []models.FlexColumn{
	{
		Text:     "testText",
		BrowseID: "testID",
		PageType: "MUSIC_PAGE_TYPE_ARTIST",
	},
	{
		Text:     "testText1",
		BrowseID: "testID1",
		PageType: "MUSIC_PAGE_TYPE_ARTIST",
	},
}
