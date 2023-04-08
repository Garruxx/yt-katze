package data

import "katze/src/models/external"

var THUMBNAIL_RENDERER = external.ThumbnailRendererClass{
	MusicThumbnailRenderer: external.MusicThumbnailRenderer{
		Thumbnail: external.MusicThumbnailRendererThumbnail{
			Thumbnails: []external.ThumbnailElement{
				{
					URL:    "https://example.com/image1.jpg",
					Width:  100,
					Height: 100,
				},
				{
					URL:    "https://example.com/image2.jpg",
					Width:  200,
					Height: 200,
				},
			},
		},
	},
}

var THUMBNAIL_RENDERER_EMPTY = external.ThumbnailRendererClass{}