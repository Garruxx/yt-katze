package data

import "katze/src/models/external"

var TO_FLEX_COLUMN_SUBTITLE_VALID = external.Subtitle{
	Runs: []external.TentacledRun{
		{
			Text: "testText",
		},
		{
			Text: "testText1",
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
		},
		{
			Text: "testText2",
		},
	},
}

var TO_FLEX_COLUMN_SUBTITLE_VALID_JUST_TEXT = external.Subtitle{
	Runs: []external.TentacledRun{
		{
			Text: "testText",
		},
	},
}

var TO_FLEX_COLUMN_SUBTITLE_EMPTY = external.Subtitle{}

var TO_FLEX_COLUMN_SUBTITLE_INVALID_EMPTY_TEXT = external.Subtitle{
	Runs: []external.TentacledRun{
		{
			Text: "",
		},
	},
}

var TO_FLEX_COLUMN_SUBTITLE_INVALID_AN_EMPTY_TEXT = external.Subtitle{
	Runs: []external.TentacledRun{
		{
			Text: "",
		},
		{
			Text: "testText1",
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
		},
	},
}

var TO_FLEX_COLUMN_SUBTITLE_INVALID_EMPTY_BROWSE_ID = external.Subtitle{
	Runs: []external.TentacledRun{
		{
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
		},
	},
}

var TO_FLEX_COLUMN_SUBTITLE_INVALID_EMPTY_PAGE_TYPE = external.Subtitle{
	Runs: []external.TentacledRun{
		{Text: "testText",
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
		},
		{
			Text: "testText1",
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
		},
	},
}
