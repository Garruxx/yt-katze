package utils

import (
	"katze/src/models"
	"katze/src/models/external"
)

func GetThumbnail(thumbnailRender external.ThumbnailRendererClass) []models.Thumbnail {

	var thumbnails []models.Thumbnail
	thumbnailsRenderer := thumbnailRender.MusicThumbnailRenderer.Thumbnail.Thumbnails
	for _, thumbnail := range thumbnailsRenderer {
		thumbnails = append(thumbnails, models.Thumbnail{
			URL:    thumbnail.URL,
			Width:  int(thumbnail.Width),
			Height: int(thumbnail.Height),
		})
	}

	if len(thumbnails) == 0 {
		thumbnails = append(thumbnails, models.Thumbnail{
			URL:    "https://picsum.photos/200/200?random=2",
			Width:  200,
			Height: 200,
		})
	}
	return thumbnails
}
