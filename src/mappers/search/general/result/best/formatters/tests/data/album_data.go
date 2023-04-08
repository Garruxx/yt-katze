package data

import "katze/src/models/external"

var ALBUM_MUSIC_CARD_SHELF_VALID = external.MusicCardShelfRenderer{
	Title: external.TitleClass{
		Runs: []external.FluffyRun{
			{
				Text: "AlbumTitleTest",
				NavigationEndpoint: &external.OnTap{
					ClickTrackingParams: "ClickTrackingParamsTest",
					BrowseEndpoint: &external.OnTapBrowseEndpoint{
						BrowseID: "BrowseIDTest",
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
	Subtitle: external.Subtitle{
		Runs: []external.TentacledRun{
			{
				Text: "album",
			},
			{
				Text: "testArtistName",
				NavigationEndpoint: &external.MusicResponsiveListItemRendererNavigationEndpoint{
					ClickTrackingParams: "testClickTrackingParams",
					BrowseEndpoint: external.OnTapBrowseEndpoint{
						BrowseID: "testArtistID",
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
	Thumbnail: external.ThumbnailRendererClass{
		MusicThumbnailRenderer: external.MusicThumbnailRenderer{
			Thumbnail: external.MusicThumbnailRendererThumbnail{
				Thumbnails: []external.ThumbnailElement{
					{
						URL:    "testThumbnailURL",
						Width:  100,
						Height: 100,
					},
				},
			},
		},
	},
}

var ALBUM_MUSIC_CARD_SHELF_EMPTY = external.MusicCardShelfRenderer{}

var ALBUM_MUSIC_CARD_SHELF_NO_TITLE = external.MusicCardShelfRenderer{
	Title: external.TitleClass{
		Runs: []external.FluffyRun{
			{
				Text: "AlbumTitleTest",
				NavigationEndpoint: &external.OnTap{
					ClickTrackingParams: "ClickTrackingParamsTest",
					BrowseEndpoint: &external.OnTapBrowseEndpoint{
						BrowseID: "BrowseIDTest",
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
	Subtitle: external.Subtitle{
		Runs: []external.TentacledRun{
			{
				Text: "album",
			},
			{
				Text: "",
				NavigationEndpoint: &external.MusicResponsiveListItemRendererNavigationEndpoint{
					ClickTrackingParams: "testClickTrackingParams",
					BrowseEndpoint: external.OnTapBrowseEndpoint{
						BrowseID: "testAlbumID",
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
	Thumbnail: external.ThumbnailRendererClass{
		MusicThumbnailRenderer: external.MusicThumbnailRenderer{
			Thumbnail: external.MusicThumbnailRendererThumbnail{
				Thumbnails: []external.ThumbnailElement{
					{
						URL:    "testThumbnailURL",
						Width:  100,
						Height: 100,
					},
				},
			},
		},
	},
}

var ALBUM_MUSIC_CARD_SHELF_NO_ALBUM_ID = external.MusicCardShelfRenderer{
	Title: external.TitleClass{
		Runs: []external.FluffyRun{
			{
				Text: "AlbumTitleTest",
				NavigationEndpoint: &external.OnTap{
					ClickTrackingParams: "ClickTrackingParamsTest",
					BrowseEndpoint: &external.OnTapBrowseEndpoint{
						BrowseID: "BrowseIDTest",
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
	Subtitle: external.Subtitle{
		Runs: []external.TentacledRun{
			{
				Text: "album",
			},
			{
				Text: "testAlbumTitle",
				NavigationEndpoint: &external.MusicResponsiveListItemRendererNavigationEndpoint{
					ClickTrackingParams: "testClickTrackingParams",
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
		},
	},
	Thumbnail: external.ThumbnailRendererClass{
		MusicThumbnailRenderer: external.MusicThumbnailRenderer{
			Thumbnail: external.MusicThumbnailRendererThumbnail{
				Thumbnails: []external.ThumbnailElement{
					{
						URL:    "testThumbnailURL",
						Width:  100,
						Height: 100,
					},
				},
			},
		},
	},
}

var ALBUM_MUSIC_CARD_SHELF_NO_THUMBNAILS = external.MusicCardShelfRenderer{
	Title: external.TitleClass{
		Runs: []external.FluffyRun{
			{
				Text: "AlbumTitleTest",
				NavigationEndpoint: &external.OnTap{
					ClickTrackingParams: "ClickTrackingParamsTest",
					BrowseEndpoint: &external.OnTapBrowseEndpoint{
						BrowseID: "BrowseIDTest",
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
	Subtitle: external.Subtitle{
		Runs: []external.TentacledRun{
			{
				Text: "album",
			},
			{
				Text: "testAlbumTitle",
				NavigationEndpoint: &external.MusicResponsiveListItemRendererNavigationEndpoint{
					ClickTrackingParams: "testClickTrackingParams",
					BrowseEndpoint: external.OnTapBrowseEndpoint{
						BrowseID: "testAlbumID",
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
}

var ALBUM_MUSIC_CARD_SHELF_NO_ARTIST = external.MusicCardShelfRenderer{
	Title: external.TitleClass{
		Runs: []external.FluffyRun{
			{
				Text: "AlbumTitleTest",
				NavigationEndpoint: &external.OnTap{
					ClickTrackingParams: "ClickTrackingParamsTest",
					BrowseEndpoint: &external.OnTapBrowseEndpoint{
						BrowseID: "BrowseIDTest",
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
	Subtitle: external.Subtitle{
		Runs: []external.TentacledRun{
			{
				Text: "album",
			},
		},
	},
	Thumbnail: external.ThumbnailRendererClass{
		MusicThumbnailRenderer: external.MusicThumbnailRenderer{
			Thumbnail: external.MusicThumbnailRendererThumbnail{
				Thumbnails: []external.ThumbnailElement{
					{
						URL:    "testThumbnailURL",
						Width:  100,
						Height: 100,
					},
				},
			},
		},
	},
}
