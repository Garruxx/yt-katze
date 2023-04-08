package data

import "katze/src/models/external"

var MAPPER_FLEX_COLUMNS_VALID = []external.FluffyFlexColumn{
	{
		MusicResponsiveListItemFlexColumnRenderer: external.FluffyMusicResponsiveListItemFlexColumnRenderer{
			Text: external.TitleClass{
				Runs: []external.FluffyRun{
					{
						Text: "test",
						NavigationEndpoint: &external.OnTap{
							BrowseEndpoint: &external.OnTapBrowseEndpoint{
								BrowseID: "testBrowseID",
								BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{
									BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
										PageType: "testPageType",
									},
								},
							},
						},
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
						Text: "test1",
						NavigationEndpoint: &external.OnTap{
							WatchEndpoint: &external.OnTapWatchEndpoint{
								VideoID: "testWatchID",
								WatchEndpointMusicSupportedConfigs: external.WatchEndpointMusicSupportedConfigs{
									WatchEndpointMusicConfig: external.WatchEndpointMusicConfig{
										MusicVideoType: "testType",
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

var MAPPER_FLEX_COLUMNS_WITH_AN_EMPTY_TEXT = []external.FluffyFlexColumn{
	{
		MusicResponsiveListItemFlexColumnRenderer: external.FluffyMusicResponsiveListItemFlexColumnRenderer{
			Text: external.TitleClass{
				Runs: []external.FluffyRun{
					{
						Text: "",
						NavigationEndpoint: &external.OnTap{
							BrowseEndpoint: &external.OnTapBrowseEndpoint{
								BrowseID: "test",
								BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{
									BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
										PageType: "test",
									},
								},
							},
						},
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
						Text: "test",
						NavigationEndpoint: &external.OnTap{
							WatchEndpoint: &external.OnTapWatchEndpoint{
								VideoID: "test",
								WatchEndpointMusicSupportedConfigs: external.WatchEndpointMusicSupportedConfigs{
									WatchEndpointMusicConfig: external.WatchEndpointMusicConfig{
										MusicVideoType: "testType",
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

var MAPPER_FLEX_COLUMNS_WITH_EMPTY_TEXT = []external.FluffyFlexColumn{
	{
		MusicResponsiveListItemFlexColumnRenderer: external.FluffyMusicResponsiveListItemFlexColumnRenderer{
			Text: external.TitleClass{
				Runs: []external.FluffyRun{
					{
						Text: "",
						NavigationEndpoint: &external.OnTap{
							BrowseEndpoint: &external.OnTapBrowseEndpoint{
								BrowseID: "test",
								BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{
									BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
										PageType: "test",
									},
								},
							},
						},
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
						Text: "",
						NavigationEndpoint: &external.OnTap{
							WatchEndpoint: &external.OnTapWatchEndpoint{
								VideoID: "test",
								WatchEndpointMusicSupportedConfigs: external.WatchEndpointMusicSupportedConfigs{
									WatchEndpointMusicConfig: external.WatchEndpointMusicConfig{
										MusicVideoType: "testType",
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

var MAPPER_FLEX_COLUMNS_WITHOUT_NAVIGATION_ENDPOINT = []external.FluffyFlexColumn{
	{
		MusicResponsiveListItemFlexColumnRenderer: external.FluffyMusicResponsiveListItemFlexColumnRenderer{
			Text: external.TitleClass{
				Runs: []external.FluffyRun{
					{
						Text: "test",
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
						Text: "test",
					},
				},
			},
		},
	},
}

var MAPPER_FLEX_COLUMNS_EMPTY = []external.FluffyFlexColumn{}

var MAPPER_FLEX_COLUMNS_ONLY_ONE_COLUMN = []external.FluffyFlexColumn{
	{
		MusicResponsiveListItemFlexColumnRenderer: external.FluffyMusicResponsiveListItemFlexColumnRenderer{
			Text: external.TitleClass{
				Runs: []external.FluffyRun{
					{
						Text: "test",
					},
				},
			},
		},
	},
}

var MAPPER_FLEX_COLUMNS_THREE_COLUMNS = []external.FluffyFlexColumn{
	{
		MusicResponsiveListItemFlexColumnRenderer: external.FluffyMusicResponsiveListItemFlexColumnRenderer{
			Text: external.TitleClass{
				Runs: []external.FluffyRun{
					{
						Text: "test",
						NavigationEndpoint: &external.OnTap{
							BrowseEndpoint: &external.OnTapBrowseEndpoint{
								BrowseID: "test",
								BrowseEndpointContextSupportedConfigs: external.BrowseEndpointContextSupportedConfigs{
									BrowseEndpointContextMusicConfig: external.BrowseEndpointContextMusicConfig{
										PageType: "test",
									},
								},
							},
						},
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
						Text: "test",
						NavigationEndpoint: &external.OnTap{
							WatchEndpoint: &external.OnTapWatchEndpoint{
								VideoID: "test",
								WatchEndpointMusicSupportedConfigs: external.WatchEndpointMusicSupportedConfigs{
									WatchEndpointMusicConfig: external.WatchEndpointMusicConfig{
										MusicVideoType: "testType",
									},
								},
							},
						},
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
						Text: "test",
						NavigationEndpoint: &external.OnTap{
							WatchEndpoint: &external.OnTapWatchEndpoint{
								VideoID: "test",
								WatchEndpointMusicSupportedConfigs: external.WatchEndpointMusicSupportedConfigs{
									WatchEndpointMusicConfig: external.WatchEndpointMusicConfig{
										MusicVideoType: "testType",
									},
								},
							},
						},
					},
				},
			},
		},
	},
}
