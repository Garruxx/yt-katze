package data

import "katze/src/models/external"

var SHELF_VALIDATOR_VALID = external.MusicCardShelfRenderer{
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
								PageType: "TEST_PAGE_TYPE",
							},
						},
					},
				},
			},
		},
	},
}

var SHELF_VALIDATOR_VALID_MUSIC = external.MusicCardShelfRenderer{
	Title: external.TitleClass{
		Runs: []external.FluffyRun{
			{
				Text: "SongTitleTest",
				NavigationEndpoint: &external.OnTap{
					ClickTrackingParams: "ClickTrackingParamsTest",
					WatchEndpoint: &external.OnTapWatchEndpoint{
						VideoID: "VideoIDTest",
						WatchEndpointMusicSupportedConfigs: external.WatchEndpointMusicSupportedConfigs{
							WatchEndpointMusicConfig: external.WatchEndpointMusicConfig{
								MusicVideoType: "TEST_PAGE_TYPE_MUSIC",
							},
						},
					},
				},
			},
		},
	},
}

var SHELF_VALIDATOR_INVALID_EMPTY = external.MusicCardShelfRenderer{}

var SHELF_VALIDATOR_INVALID_NO_BROWSE_ENDPOINT = external.MusicCardShelfRenderer{
	Title: external.TitleClass{
		Runs: []external.FluffyRun{
			{
				Text: "AlbumTitleTest",
				NavigationEndpoint: &external.OnTap{
					ClickTrackingParams: "ClickTrackingParamsTest",
					BrowseEndpoint:      &external.OnTapBrowseEndpoint{},
				},
			},
		},
	},
}

var SHELF_VALIDATOR_PAGE_TYPE_EMPTY = external.MusicCardShelfRenderer{
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
								PageType: "",
							},
						},
					},
				},
			},
		},
	},
}

var SHELF_VALIDATOR_INVALID_NO_BROWSE_ID = external.MusicCardShelfRenderer{
	Title: external.TitleClass{
		Runs: []external.FluffyRun{
			{
				Text: "AlbumTitleTest",
				NavigationEndpoint: &external.OnTap{
					ClickTrackingParams: "ClickTrackingParamsTest",
					BrowseEndpoint: &external.OnTapBrowseEndpoint{
						BrowseID: "",
						BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{
							BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
								PageType: "TEST_PAGE_TYPE",
							},
						},
					},
				},
			},
		},
	},
}

var SHELF_VALIDATOR_INVALID = external.MusicCardShelfRenderer{
	Title: external.TitleClass{
		Runs: []external.FluffyRun{
			{
				Text:               "AlbumTitleTest",
				NavigationEndpoint: &external.OnTap{},
			},
		},
	},
}


