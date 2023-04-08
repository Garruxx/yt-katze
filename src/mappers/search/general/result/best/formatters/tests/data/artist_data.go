package data

import "katze/src/models/external"

var ARTIST_MUSIC_CARD_SHELF_VALID = external.MusicCardShelfRenderer{
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
	Subtitle: external.Subtitle{
		Runs: []external.TentacledRun{
			{
				Text: "example",
			},
			{
				Text: "example",
			},
			{
				Text: "example",
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

var ARTIST_MUSIC_CARD_SHELF_VALID_NO_SUBTITLE = external.MusicCardShelfRenderer{
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

var ARTIST_MUSIC_CARD_SHELF_INVALID = external.MusicCardShelfRenderer{
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
								PageType: "MUSIC_PAGE_TYPE_INVALID",
							},
						},
					},
				},
			},
		},
	},
}

var ARTIST_MUSIC_CARD_SHELF_INVALID_NO_BROWSE_ENDPOINT = external.MusicCardShelfRenderer{
	Title: external.TitleClass{
		Runs: []external.FluffyRun{
			{
				Text: "ArtistNameTest",
				NavigationEndpoint: &external.OnTap{
					ClickTrackingParams: "ClickTrackingParamsTest",
					BrowseEndpoint:      &external.OnTapBrowseEndpoint{},
				},
			},
		},
	},
}

var ARTIST_MUSIC_CARD_SHELF_INVALID_NO_TYPE_ARTIST = external.MusicCardShelfRenderer{
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
								PageType: "MUSIC_PAGE_TYPE_INVALID",
							},
						},
					},
				},
			},
		},
	},
}

var ARTIST_MUSIC_CARD_SHELF_INVALID_NO_BROWSE_ID = external.MusicCardShelfRenderer{
	Title: external.TitleClass{
		Runs: []external.FluffyRun{
			{
				Text: "ArtistNameTest",
				NavigationEndpoint: &external.OnTap{
					ClickTrackingParams: "ClickTrackingParamsTest",
					BrowseEndpoint: &external.OnTapBrowseEndpoint{
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
}


var ARTIST_MUSIC_CARD_SHELF_NO_BROWSE_ENDPOINT_CONTEXT_SUPPORTED_CONFIGS = external.MusicCardShelfRenderer{
	Title: external.TitleClass{
		Runs: []external.FluffyRun{
			{
				Text: "ArtistNameTest",
				NavigationEndpoint: &external.OnTap{
					ClickTrackingParams: "ClickTrackingParamsTest",
					BrowseEndpoint: &external.OnTapBrowseEndpoint{
						BrowseID: "BrowseIDTest",
					},
				},
			},
		},
	},
}
