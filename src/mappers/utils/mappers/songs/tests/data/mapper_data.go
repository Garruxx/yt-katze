package data

import "katze/src/models/external"

// utils data
var flexColumData = []external.FluffyFlexColumn{
	{
		MusicResponsiveListItemFlexColumnRenderer: external.FluffyMusicResponsiveListItemFlexColumnRenderer{
			Text: external.TitleClass{
				Runs: []external.FluffyRun{
					{
						Text: "TitleTest",
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
						Text: "artistNameTest",
						NavigationEndpoint: &external.OnTap{
							ClickTrackingParams: "",
							BrowseEndpoint: &external.OnTapBrowseEndpoint{
								BrowseID: "testID",
								BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{
									BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
										PageType: "MUSIC_PAGE_TYPE_ARTIST",
									},
								},
							},
						},
					},
					{
						Text: "albumNameTest",
						NavigationEndpoint: &external.OnTap{
							ClickTrackingParams: "",
							BrowseEndpoint: &external.OnTapBrowseEndpoint{
								BrowseID: "testID",
								BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{
									BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
										PageType: "MUSIC_PAGE_TYPE_ALBUM",
									},
								},
							},
						},
					},

					{
						Text: "22:22",
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

var ITEMS_SONGS_VALID_TWO_SONGS = []external.MischievousContent{
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
			PlaylistItemData: &external.PlaylistItemData{
				VideoID: "testID",
			},
			Badges: &[]external.Badge{},
		},
	},
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
			PlaylistItemData: &external.PlaylistItemData{
				VideoID: "testID",
			},
			Badges: &[]external.Badge{},
		},
	},
}

var ITEMS_SONGS_VALID_ONE_SONG = []external.MischievousContent{
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
			PlaylistItemData: &external.PlaylistItemData{
				VideoID: "testID",
			},
			Badges: &[]external.Badge{},
		},
	},
}

var ITEMS_SONGS_EMPTY = []external.MischievousContent{}
