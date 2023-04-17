package test

import (
	"katze/src/mappers/utils"
	"katze/src/mappers/utils/tests/data"
	"katze/src/models/external"
	"testing"
)

func TestGetThumbnail(t *testing.T) {

	// Test case 1: ThumbnailRendererClass is valid
	thumbnailRenderer := data.THUMBNAIL_RENDERER
	expected := []external.ThumbnailElement{
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
	}

	thumbnails := utils.GetThumbnail(thumbnailRenderer)
	if thbLen := len(thumbnails); thbLen != len(expected) {
		t.Errorf(
			"Test case 1 failed: expected %v, got %v",
			len(expected), thbLen,
		)
	}
	// Test case 1.1: thumbnails are equal to expected
	for i := range thumbnails {
		if thumbnails[i] != expected[i] {
			t.Errorf(
				"failed test 1.1 expected thumbnail %v, but got %v",
				expected[i], thumbnails[i],
			)
		}
	}

	// Test case 2: ThumbnailRendererClass is empty
	thumbnailRenderer = data.THUMBNAIL_RENDERER_EMPTY
	thumbnails = utils.GetThumbnail(thumbnailRenderer)
	if len(thumbnails) != 1 {
		t.Errorf(
			"Test case 2 failed: expected 1 thumbnail, got %v",
			len(thumbnails),
		)
	}
	// Test case 2.1: thumbnail URL is equal to default
	if thumbnails[0].URL != "https://picsum.photos/200/200?random=2" {
		t.Errorf(
			`Test case 2.1 failed: expected URL to be 
			 https://picsum.photos/200/200?random=2, got %v`,
			thumbnails[0].URL,
		)
	}
}
