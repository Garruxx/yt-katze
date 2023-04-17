package data

import "katze/src/models/external"

var TO_COLUMN_RUN_VALID = external.TentacledRun{
	Text: "testText",
	NavigationEndpoint: &external.MusicResponsiveListItemRendererNavigationEndpoint{

		ClickTrackingParams: "testClickTrackingParams",
		BrowseEndpoint: external.OnTapBrowseEndpoint{
			BrowseID: "testID",
			BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{
				BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
					PageType: "testPageType",
				},
			},
		},
	},
}

var TO_COLUMN_RUN_INVALID_EMPTY = external.TentacledRun{}

var TO_COLUMN_RUN_INVALID_EMPTY_TEXT = external.TentacledRun{
	Text: "",
	NavigationEndpoint: &external.MusicResponsiveListItemRendererNavigationEndpoint{

		ClickTrackingParams: "testClickTrackingParams",
		BrowseEndpoint: external.OnTapBrowseEndpoint{
			BrowseID: "testID",
			BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{
				BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
					PageType: "testPageType",
				},
			},
		},
	},
}

var TO_COLUMN_RUN_INVALID_EMPTY_BROWSE_ID = external.TentacledRun{
	Text: "testText",
	NavigationEndpoint: &external.MusicResponsiveListItemRendererNavigationEndpoint{

		ClickTrackingParams: "testClickTrackingParams",
		BrowseEndpoint: external.OnTapBrowseEndpoint{
			BrowseID: "",
			BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{
				BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
					PageType: "testPageType",
				},
			},
		},
	},
}

var TO_COLUMN_RUN_INVALID_EMPTY_PAGE_TYPE = external.TentacledRun{
	Text: "testText",
	NavigationEndpoint: &external.MusicResponsiveListItemRendererNavigationEndpoint{

		ClickTrackingParams: "testClickTrackingParams",
		BrowseEndpoint: external.OnTapBrowseEndpoint{
			BrowseID: "testID",
			BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{
				BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
					PageType: "",
				},
			},
		},
	},
}

var TO_COLUMN_RUN_VALID_JUST_TEXT = external.TentacledRun{
	Text: "testText",
}
