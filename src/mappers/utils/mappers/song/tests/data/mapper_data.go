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

var ITEM_SONG_VALID = external.MischievousContent{
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
}

var ITEM_SONG_VALID_NO_EXPLICIT = external.MischievousContent{
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
	},
}

var ITEM_SONG_INVALID_EMPTY = external.MischievousContent{
	MusicResponsiveListItemRenderer: external.StickyMusicResponsiveListItemRenderer{},
}

var ITEM_SONG_INVALID_NO_VIDEO_ID = external.MischievousContent{
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
			VideoID: "",
		},
		Badges: &[]external.Badge{},
	},
}

var ITEM_SONG_INVALID_NO_FLEX_COLUMNS = external.MischievousContent{
	MusicResponsiveListItemRenderer: external.StickyMusicResponsiveListItemRenderer{

		Thumbnail: external.ThumbnailRendererClass{
			MusicThumbnailRenderer: external.MusicThumbnailRenderer{
				Thumbnail: external.MusicThumbnailRendererThumbnail{
					Thumbnails: thumbnailsData,
				},
			},
		},

		PlaylistItemData: &external.PlaylistItemData{
			VideoID: "testID",
		},
		Badges: &[]external.Badge{},
	},
}
