package data

import (
	"katze/src/models/external"
)

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

var ITEM_ARTIST_VALID = external.MischievousContent{

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
}
var ITEM_ARTIST_INVALID_EMPTY = external.MischievousContent{}

var ITEM_ARTIST_INVALID_PAGE_TYPE = external.MischievousContent{

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
						PageType: "MUSIC_PAGE_INVALID",
					},
				},
			},
		},
	},
}

var ITEM_ARTIST_INVALID_FLEX_COLUMN_LENGTH = external.MischievousContent{

	MusicResponsiveListItemRenderer: external.StickyMusicResponsiveListItemRenderer{

		Thumbnail: external.ThumbnailRendererClass{
			MusicThumbnailRenderer: external.MusicThumbnailRenderer{
				Thumbnail: external.MusicThumbnailRendererThumbnail{
					Thumbnails: thumbnailsData,
				},
			},
		},

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
}

var ITEM_ARTIST_INVALID_NO_ID = external.MischievousContent{

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
				BrowseID: "",

				BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{

					BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
						PageType: "MUSIC_PAGE_TYPE_ARTIST",
					},
				},
			},
		},
	},
}
