package data

import "katze/src/models/external"

var GET_BEST_RESULT = []external.MagentaContent{

	{
		MusicShelfRenderer: &external.FluffyMusicShelfRenderer{
			Title: external.TitleElement{
				Runs: []external.DescriptionRun{
					{
						Text: "test0",
					},
				},
			},
		},
	},
	{
		MusicShelfRenderer: &external.FluffyMusicShelfRenderer{
			Title: external.TitleElement{
				Runs: []external.DescriptionRun{
					{
						Text: "test1",
					},
				},
			},
		},
	},
	{
		MusicCardShelfRenderer: &external.MusicCardShelfRenderer{
			Title: external.TitleClass{
				Runs: []external.FluffyRun{
					{
						Text: "test",
					},
				},
			},
		},
	},
}

var GET_BEST_RESULT_NO_BEST_OUTCOME = []external.MagentaContent{

	{
		MusicShelfRenderer: &external.FluffyMusicShelfRenderer{
			Title: external.TitleElement{
				Runs: []external.DescriptionRun{
					{
						Text: "test0",
					},
				},
			},
		},
	},
	{
		MusicShelfRenderer: &external.FluffyMusicShelfRenderer{
			Title: external.TitleElement{
				Runs: []external.DescriptionRun{
					{
						Text: "test1",
					},
				},
			},
		},
	},
}
