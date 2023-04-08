package data

import "katze/src/models/external"

var FLEX_COLUMN_DATA = external.FluffyRun{
	Text: "Test",
	NavigationEndpoint: &external.OnTap{
		BrowseEndpoint: &external.OnTapBrowseEndpoint{
			BrowseID: "TestID",
			BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{
				BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
					PageType: "PageTypeTest",
				},
			},
		},
	},
}

var FLEX_COLUMN_DATA_EMPTY_NAVIGATION_ENDPOINT = external.FluffyRun{
	Text: "Test",
}

var FLEX_COLUMN_DATA_EMPTY_TEXT = external.FluffyRun{
	Text: "",
}

var FLEX_COLUMN_DATA_EMPTY_BROWSE_ID = external.FluffyRun{
	Text: "Test",
	NavigationEndpoint: &external.OnTap{
		BrowseEndpoint: &external.OnTapBrowseEndpoint{
			BrowseID: "",
			BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{
				BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
					PageType: "PageTypeTest",
				},
			},
		},
	},
}

var FLEX_COLUMN_DATA_EMPTY_PAGE_TYPE = external.FluffyRun{
	Text: "Test",
	NavigationEndpoint: &external.OnTap{
		BrowseEndpoint: &external.OnTapBrowseEndpoint{
			BrowseID: "TestID",
			BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{
				BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
					PageType: "",
				},
			},
		},
	},
}

var FLEX_COLUMN_DATA_EMPTY_BROWSE_ID_AND_PAGE_TYPE = external.FluffyRun{
	Text: "Test",
	NavigationEndpoint: &external.OnTap{
		BrowseEndpoint: &external.OnTapBrowseEndpoint{
			BrowseID: "",
			BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{
				BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
					PageType: "",
				},
			},
		},
	},
}

var FLEX_COLUMN_DATA_WITH_MUSIC_ID = external.FluffyRun{
	Text: "test",
	NavigationEndpoint: &external.OnTap{
		WatchEndpoint: &external.OnTapWatchEndpoint{
			VideoID: "test",
			WatchEndpointMusicSupportedConfigs: external.WatchEndpointMusicSupportedConfigs{
				WatchEndpointMusicConfig: external.WatchEndpointMusicConfig{
					MusicVideoType: "MusicVideoTypeTest",
				},
			},
		},
	},
}

var FLEX_COLUMN_DATA_EMPTY = external.FluffyRun{}
