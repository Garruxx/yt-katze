package data

import "katze/src/models/external"

var SONG_MUSIC_CARD_SHELF_VALID = external.MusicCardShelfRenderer{
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

var SONG_MUSIC_CARD_SHELF_NO_THUMBNAIL = external.MusicCardShelfRenderer{
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

var SONG_MUSIC_CARD_SHELF_EMPTY = external.MusicCardShelfRenderer{}
var SONG_MUSIC_CARD_SHELF_NO_TYPE_MUSIC = external.MusicCardShelfRenderer{
	Title: external.TitleClass{
		Runs: []external.FluffyRun{
			{
				Text: "test",
				NavigationEndpoint: &external.OnTap{
					ClickTrackingParams: "testClickTrackingParams",
					BrowseEndpoint: &external.OnTapBrowseEndpoint{
						BrowseID: "browseIDTest",
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

var SONG_MUSIC_CARD_SHELF_NO_ARTIST = external.MusicCardShelfRenderer{
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
	Subtitle: external.Subtitle{
		Runs: []external.TentacledRun{
			{
				Text: "albumTest",
				NavigationEndpoint: &external.MusicResponsiveListItemRendererNavigationEndpoint{
					ClickTrackingParams: "testClickTrackingParams",
					BrowseEndpoint: external.OnTapBrowseEndpoint{
						BrowseID: "browseIDTest",
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

var SONG_MUSIC_CARD_SHELF_NO_ALBUM = external.MusicCardShelfRenderer{
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
	Subtitle: external.Subtitle{
		Runs: []external.TentacledRun{
			{
				Text: "artistTest",
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
				Text: "2:22",
			},
		},
	},
}

var SONG_MUSIC_CARD_SHELF_NO_DURATION = external.MusicCardShelfRenderer{
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
	Subtitle: external.Subtitle{
		Runs: []external.TentacledRun{
			{
				Text: "artistTest",
				NavigationEndpoint: &external.MusicResponsiveListItemRendererNavigationEndpoint{
					ClickTrackingParams: "testClickTrackingParams",
					BrowseEndpoint: external.OnTapBrowseEndpoint{
						BrowseID: "testBrowseID",
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
						BrowseID: "testBrowseID",
						BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{
							BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
								PageType: "MUSIC_PAGE_TYPE_ARTIST",
							},
						},
					},
				},
			},
			{
				Text: "albumTest",
				NavigationEndpoint: &external.MusicResponsiveListItemRendererNavigationEndpoint{
					ClickTrackingParams: "testClickTrackingParams",

					BrowseEndpoint: external.OnTapBrowseEndpoint{
						BrowseID: "testBrowseID",
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

var SONG_MUSIC_CARD_SHELF_INVALID_ARTIST = external.MusicCardShelfRenderer{
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
	Subtitle: external.Subtitle{
		Runs: []external.TentacledRun{
			{
				Text: "albumTest",
				NavigationEndpoint: &external.MusicResponsiveListItemRendererNavigationEndpoint{
					ClickTrackingParams: "testClickTrackingParams",
					BrowseEndpoint: external.OnTapBrowseEndpoint{
						// No BrowseID
						BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{
							BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
								PageType: "MUSIC_PAGE_TYPE_ARTIST",
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

var SONG_MUSIC_CARD_SHELF_INVALID_ALBUM = external.MusicCardShelfRenderer{
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
	Subtitle: external.Subtitle{
		Runs: []external.TentacledRun{
			{
				Text: "artistTest",
				NavigationEndpoint: &external.MusicResponsiveListItemRendererNavigationEndpoint{
					ClickTrackingParams: "testClickTrackingParams",
					BrowseEndpoint: external.OnTapBrowseEndpoint{
						BrowseID: "testBrowseID",
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
						BrowseID: "testBrowseID",
						BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{
							BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
								PageType: "MUSIC_PAGE_TYPE_ARTIST",
							},
						},
					},
				},
			},
			{
				Text: "albumTest",
				NavigationEndpoint: &external.MusicResponsiveListItemRendererNavigationEndpoint{
					ClickTrackingParams: "testClickTrackingParams",
					BrowseEndpoint: external.OnTapBrowseEndpoint{
						// No BrowseID
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
