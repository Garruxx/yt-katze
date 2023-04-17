package utils

import (
	"katze/src/models/external"
)

func GetThumbnail(thumbnailRender external.ThumbnailRendererClass) []external.ThumbnailElement {

	thumbnails := thumbnailRender.MusicThumbnailRenderer.Thumbnail.Thumbnails
	if len(thumbnails) == 0 {
		thumbnails = append(thumbnails, external.ThumbnailElement{
			URL:    "https://picsum.photos/200/200?random=2",
			Width:  200,
			Height: 200,
		})
	}
	return thumbnails
}
