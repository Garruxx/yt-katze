package data

import "katze/src/models/external"

// util data
var flexColumData = []external.FluffyFlexColumn{
	{
		MusicResponsiveListItemFlexColumnRenderer: external.FluffyMusicResponsiveListItemFlexColumnRenderer{
			Text: external.TitleClass{
				Runs: []external.FluffyRun{
					{
						Text: "artistName",
					},
					{
						Text: "artistName",
					},
					{
						Text: "noArtistText",
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
						Text: "artistName",
					},
					{
						Text: "artistName",
					},
					{
						Text: "noArtistText",
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

// MusicResponsiveListItemRenderer data
var MUSIC_RESPONSIVE_LIST_ITEM_RENDERER_ARTIST = external.MischievousContent{

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

var MUSIC_RESPONSIVE_LIST_ITEM_RENDERER_ALBUM = external.MischievousContent{

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
				BrowseID: "UC-9-kyTW8ZkZNDHQJ6FgpwQ",

				BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{

					BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
						PageType: "MUSIC_PAGE_TYPE_ALBUM",
					},
				},
			},
		},

		Badges: &[]external.Badge{},
	},
}

var MUSIC_RESPONSIVE_LIST_ITEM_RENDERER_SONG = external.MischievousContent{

	MusicResponsiveListItemRenderer: external.StickyMusicResponsiveListItemRenderer{

		Thumbnail: external.ThumbnailRendererClass{
			MusicThumbnailRenderer: external.MusicThumbnailRenderer{
				Thumbnail: external.MusicThumbnailRendererThumbnail{
					Thumbnails: thumbnailsData,
				},
			},
		},
		FlexColumns: flexColumData,
		PlaylistItemData: &external.PlaylistItemData{
			VideoID: "example.id",
		},

		Badges: &[]external.Badge{},
	},
}

var MUSIC_RESPONSIVE_LIST_ITEM_RENDERER_INVALID = external.MischievousContent{

	MusicResponsiveListItemRenderer: external.StickyMusicResponsiveListItemRenderer{

		FlexColumns: []external.FluffyFlexColumn{},
		PlaylistItemData: &external.PlaylistItemData{
			VideoID: "",
		},

		Badges: &[]external.Badge{},
	},
}

var MUSIC_RESPONSIVE_LIST_ITEM_RENDERER_EMPTY = external.MischievousContent{}
