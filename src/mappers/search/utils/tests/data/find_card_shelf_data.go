package data

import "katze/src/models/external"

var TAB_CONTENT_RENDERER = []external.MagentaContent{

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

var TAB_CONTENT_RENDERER_EMPTY = []external.MagentaContent{}