package data

import "katze/src/models/external"

var BEST_VALID_ALBUM = external.MusicCardShelfRenderer{
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

var BEST_VALID_SONG = external.MusicCardShelfRenderer{
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
	Title: external.TitleClass{
		Runs: []external.FluffyRun{
			{
				Text: "test",
				NavigationEndpoint: &external.OnTap{
					ClickTrackingParams: "testClickTrackingParams",
					WatchEndpoint: &external.OnTapWatchEndpoint{
						VideoID: "testVideoID",
						WatchEndpointMusicSupportedConfigs: external.WatchEndpointMusicSupportedConfigs{
							WatchEndpointMusicConfig: external.WatchEndpointMusicConfig{
								MusicVideoType: "MUSIC_VIDEO_TYPE_ATV",
							},
						},
					},
				},
			},
		},
	},
	SubtitleBadges: []external.SubtitleBadge{},
	Subtitle: external.Subtitle{
		Runs: []external.TentacledRun{
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
			{
				Text: "artistTest1",
				NavigationEndpoint: &external.MusicResponsiveListItemRendererNavigationEndpoint{
					ClickTrackingParams: "testClickTrackingParams",
					BrowseEndpoint: external.OnTapBrowseEndpoint{
						BrowseID: "browseIDTest",
						BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{
							BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
								PageType: "MUSIC_PAGE_TYPE_ARTIST",
							},
						},
					},
				},
			},
			{
				Text: "testAlbumTitle",
				NavigationEndpoint: &external.MusicResponsiveListItemRendererNavigationEndpoint{
					ClickTrackingParams: "testClickTrackingParams",
					BrowseEndpoint: external.OnTapBrowseEndpoint{
						BrowseID: "testAlbumID",
						BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{
							BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
								PageType: "MUSIC_PAGE_TYPE_ALBUM",
							},
						},
					},
				},
			},
			{
				Text: "2:22",
			},
		},
	},
}

var BEST_VALID_ARTIST = external.MusicCardShelfRenderer{
	Title: external.TitleClass{
		Runs: []external.FluffyRun{
			{
				Text: "ArtistNameTest",
				NavigationEndpoint: &external.OnTap{
					ClickTrackingParams: "ClickTrackingParamsTest",
					BrowseEndpoint: &external.OnTapBrowseEndpoint{
						BrowseID: "BrowseIDTest",
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

var BEST_INVALID_EMPTY = external.MusicCardShelfRenderer{}

var BEST_VALID_DONT_KNOW_TYPE = external.MusicCardShelfRenderer{
	Title: external.TitleClass{
		Runs: []external.FluffyRun{
			{
				Text: "ArtistNameTest",
				NavigationEndpoint: &external.OnTap{
					ClickTrackingParams: "ClickTrackingParamsTest",
					BrowseEndpoint: &external.OnTapBrowseEndpoint{
						BrowseID: "BrowseIDTest",
						BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{
							BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
								PageType: "MUSIC_PAGE_TYPE_NKOW",
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