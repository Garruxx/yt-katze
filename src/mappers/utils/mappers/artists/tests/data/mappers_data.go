package data

import "katze/src/models/external"

// utils data
var flexColumData = []external.FluffyFlexColumn{
	{
		MusicResponsiveListItemFlexColumnRenderer: external.FluffyMusicResponsiveListItemFlexColumnRenderer{
			Text: external.TitleClass{
				Runs: []external.FluffyRun{
					{
						Text: "ArtistNameTest",
					},
				},
			},
		},
	},
	{
		MusicResponsiveListItemFlexColumnRenderer: external.FluffyMusicResponsiveListItemFlexColumnRenderer{
			Text: external.TitleClass{
				Runs: []external.FluffyRun{
					{
						Text: "example",
					},
					{
						Text: "example",
					},
				},
			},
		},
	},
}

var thumbnailsData = []external.ThumbnailElement{
	{
		URL:    "example.url",
		Width:  222,
		Height: 222,
	},
	{
		URL:    "example.url",
		Width:  222,
		Height: 222,
	},
}

var ITEMS_ARTISTS_VALID_ONE_ARTIST = []external.MischievousContent{
	{

		MusicResponsiveListItemRenderer: external.StickyMusicResponsiveListItemRenderer{

			Thumbnail: external.ThumbnailRendererClass{
				MusicThumbnailRenderer: external.MusicThumbnailRenderer{
					Thumbnail: external.MusicThumbnailRendererThumbnail{
						Thumbnails: thumbnailsData,
					},
				},
			},

			FlexColumns: flexColumData,
			NavigationEndpoint: &external.MusicResponsiveListItemRendererNavigationEndpoint{

				ClickTrackingParams: "CAQQybcCIhMI8f3X2J3ZAhXZ6QIVHJ0aHw5a",
				BrowseEndpoint: external.OnTapBrowseEndpoint{
					BrowseID: "testID",

					BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{

						BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
							PageType: "MUSIC_PAGE_TYPE_ARTIST",
						},
					},
				},
			},
		},
	},
}

var ITEMS_ARTISTS_VALID_TWO_ARTISTS = []external.MischievousContent{
	{

		MusicResponsiveListItemRenderer: external.StickyMusicResponsiveListItemRenderer{

			Thumbnail: external.ThumbnailRendererClass{
				MusicThumbnailRenderer: external.MusicThumbnailRenderer{
					Thumbnail: external.MusicThumbnailRendererThumbnail{
						Thumbnails: thumbnailsData,
					},
				},
			},

			FlexColumns: flexColumData,
			NavigationEndpoint: &external.MusicResponsiveListItemRendererNavigationEndpoint{

				ClickTrackingParams: "CAQQybcCIhMI8f3X2J3ZAhXZ6QIVHJ0aHw5a",
				BrowseEndpoint: external.OnTapBrowseEndpoint{
					BrowseID: "testID",

					BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{

						BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
							PageType: "MUSIC_PAGE_TYPE_ARTIST",
						},
					},
				},
			},
		},
	}, {

		MusicResponsiveListItemRenderer: external.StickyMusicResponsiveListItemRenderer{

			Thumbnail: external.ThumbnailRendererClass{
				MusicThumbnailRenderer: external.MusicThumbnailRenderer{
					Thumbnail: external.MusicThumbnailRendererThumbnail{
						Thumbnails: thumbnailsData,
					},
				},
			},

			FlexColumns: flexColumData,
			NavigationEndpoint: &external.MusicResponsiveListItemRendererNavigationEndpoint{

				ClickTrackingParams: "CAQQybcCIhMI8f3X2J3ZAhXZ6QIVHJ0aHw5a",
				BrowseEndpoint: external.OnTapBrowseEndpoint{
					BrowseID: "testID",

					BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{

						BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
							PageType: "MUSIC_PAGE_TYPE_ARTIST",
						},
					},
				},
			},
		},
	},
}

var ITEMS_ARTISTS_VALID_EMPTY = []external.MischievousContent{}
