package data

import (
	"katze/src/mappers/browse/models"
	"katze/src/models/external"
)

var FIND_CAROUSEL_SHELF_VALID = []models.CarouselShelfRenderer{
	{
		Title:    "titleTest",
		BrowseID: "browseIDTest",
		Params:   "paramsTest",
		Contents: []external.MusicCarouselShelfRendererContent{
			{
				MusicTwoRowItemRenderer: external.ContentMusicTwoRowItemRenderer{},
			},
			{
				MusicTwoRowItemRenderer: external.ContentMusicTwoRowItemRenderer{},
			},
		},
	},
	{
		Title:    "titleTest1",
		BrowseID: "browseIDTest",
		Params:   "paramsTest",
		Contents: []external.MusicCarouselShelfRendererContent{
			{
				MusicTwoRowItemRenderer: external.ContentMusicTwoRowItemRenderer{},
			},
			{
				MusicTwoRowItemRenderer: external.ContentMusicTwoRowItemRenderer{},
			},
		},
	},
}

var FIND_CAROUSEL_SHELF_INVALID_EMPTY = []models.CarouselShelfRenderer{}
