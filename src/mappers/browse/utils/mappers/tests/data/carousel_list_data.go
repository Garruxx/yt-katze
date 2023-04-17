package data

import (
	"katze/src/mappers/browse/models"
	"katze/src/models/external"
)

var params = "paramsTest"
var CAROUSEL_LIST_VALID = models.CarouselShelfRenderer{
	Title:    "titleTest",
	BrowseID: "browseIDTest",
	Params:   "paramsTest",
	Contents: []external.MusicCarouselShelfRendererContent{
		{
			MusicTwoRowItemRenderer: external.ContentMusicTwoRowItemRenderer{
				Title: external.MusicCarouselShelfBasicHeaderRendererTitle{
					Runs: []external.StickyRun{
						{
							Text: "titleTest",
						},
					},
				},
				Subtitle: external.Subtitle{
					Runs: []external.TentacledRun{
						{
							Text: "subtitleTest",
						},
						{
							Text: "2024",
						},
					},
				},
				NavigationEndpoint: external.MusicTwoRowItemRendererNavigationEndpoint{
					BrowseEndpoint: &external.BottomEndpointBrowseEndpoint{
						BrowseID: "browseIDTest",
						Params:   &params,
					},
				},
			},
		},{
			MusicTwoRowItemRenderer: external.ContentMusicTwoRowItemRenderer{
				Title: external.MusicCarouselShelfBasicHeaderRendererTitle{
					Runs: []external.StickyRun{
						{
							Text: "titleTest",
						},
					},
				},
				Subtitle: external.Subtitle{
					Runs: []external.TentacledRun{
						{
							Text: "subtitleTest",
						},
						{
							Text: "2024",
						},
					},
				},
				NavigationEndpoint: external.MusicTwoRowItemRendererNavigationEndpoint{
					BrowseEndpoint: &external.BottomEndpointBrowseEndpoint{
						BrowseID: "browseIDTest",
						Params:   &params,
					},
				},
			},
		},
	},
}

var CAROUSEL_LIST_VALID_EMPTY = models.CarouselShelfRenderer{}
