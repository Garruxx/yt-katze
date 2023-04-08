package data

import "katze/src/models/external"

//utils data
var searchParams = "test1"

var flexColumData = []external.FluffyFlexColumn{
	{
		MusicResponsiveListItemFlexColumnRenderer: external.FluffyMusicResponsiveListItemFlexColumnRenderer{
			Text: external.TitleClass{
				Runs: []external.FluffyRun{
					{
						Text: "song name",
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

// MusicResponsiveListItemRenderer data
var MUSIC_SHELF_RENDERER_VALID_MUSIC = external.FluffyMusicShelfRenderer{

	Title: external.TitleElement{
		Runs: []external.DescriptionRun{
			{
				Text: "Songs",
			},
		},
	},
	Contents: []external.MischievousContent{
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
			},
		},
	},
	BottomEndpoint: external.ChipCloudChipRendererNavigationEndpoint{
		ClickTrackingParams: "test1",
		SearchEndpoint: external.SearchEndpoint{
			Params: &searchParams,
			Query:  "test1",
		},
	},
}

var MUSIC_SHELF_RENDERER_INVALID_EMPTY = external.FluffyMusicShelfRenderer{}

var MUSIC_SHELF_RENDERER_INVALID_NO_SONGS = external.FluffyMusicShelfRenderer{

	Title: external.TitleElement{
		Runs: []external.DescriptionRun{
			{
				Text: "Songs",
			},
		},
	},
	Contents: []external.MischievousContent{
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
	},
	BottomEndpoint: external.ChipCloudChipRendererNavigationEndpoint{
		ClickTrackingParams: "test1",
		SearchEndpoint: external.SearchEndpoint{
			Params: &searchParams,
			Query:  "test1",
		},
	},
}

var MUSIC_SHELF_RENDERER_INVALID_NO_CONTENTS = external.FluffyMusicShelfRenderer{
	Title: external.TitleElement{
		Runs: []external.DescriptionRun{
			{
				Text: "Songs",
			},
		},
	},
	BottomEndpoint: external.ChipCloudChipRendererNavigationEndpoint{
		ClickTrackingParams: "test1",
		SearchEndpoint: external.SearchEndpoint{
			Params: &searchParams,
			Query:  "test1",
		},
	},
}

var MUSIC_SHELF_RENDERER_INVALID_NO_BOTTOM_ENDPOINT = external.FluffyMusicShelfRenderer{
	Title: external.TitleElement{
		Runs: []external.DescriptionRun{
			{
				Text: "Albums",
			},
		},
	},
	Contents: []external.MischievousContent{
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
								PageType: "MUSIC_PAGE_TYPE_ALBUM",
							},
						},
					},
				},
			},
		},
	},
}

var MUSIC_SHELF_RENDERER_INVALID_TITLE = external.FluffyMusicShelfRenderer{

	Title: external.TitleElement{
		Runs: []external.DescriptionRun{
			{
				Text: "albums",
			},
		},
	},
	Contents: []external.MischievousContent{
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
								PageType: "MUSIC_PAGE_TYPE_ALBUM",
							},
						},
					},
				},
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
				NavigationEndpoint: &external.MusicResponsiveListItemRendererNavigationEndpoint{

					ClickTrackingParams: "CAQQybcCIhMI8f3X2J3ZAhXZ6QIVHJ0aHw5a",
					BrowseEndpoint: external.OnTapBrowseEndpoint{
						BrowseID: "testID",

						BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{

							BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
								PageType: "MUSIC_PAGE_TYPE_ALBUM",
							},
						},
					},
				},
			},
		},
	},
	BottomEndpoint: external.ChipCloudChipRendererNavigationEndpoint{
		ClickTrackingParams: "test1",
		SearchEndpoint: external.SearchEndpoint{
			Params: &searchParams,
			Query:  "test1",
		},
	},
}
