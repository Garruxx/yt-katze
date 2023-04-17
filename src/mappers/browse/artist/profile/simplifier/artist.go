package simplifier

import (
	"fmt"
	"katze/src/mappers/browse/models"
	"katze/src/mappers/utils"
	"katze/src/models/external"
)

func Artist(artist external.Artist) (models.ArtistProfile, error) {

	if artist.Error != nil {
		return models.ArtistProfile{}, fmt.Errorf(
			"error mapping artist: %v",
			artist.Error.Message,
		)
	}

	responseContext := artist.ResponseContext
	if responseContext == nil {
		return models.ArtistProfile{}, fmt.Errorf(
			"error mapping artist: response context is nil",
		)
	}

	tabs := artist.Contents.SingleColumnBrowseResultsRenderer.Tabs
	if len(tabs) == 0 {
		return models.ArtistProfile{}, fmt.Errorf(
			"error mapping artist: tabs is empty",
		)
	}

	tabContents := tabs[0].TabRenderer.Content.SectionListRenderer.Contents
	if len(tabContents) == 0 {
		return models.ArtistProfile{}, fmt.Errorf(
			"error mapping artist: tab contents is empty",
		)
	}

	headers := artist.Header
	if headers == nil {
		return models.ArtistProfile{}, fmt.Errorf(
			"error mapping artist: headers is nil",
		)
	}

	headersTitleRuns := headers.MusicImmersiveHeaderRenderer.Title.Runs
	if len(headersTitleRuns) == 0 {
		return models.ArtistProfile{}, fmt.Errorf(
			"error mapping artist: headers title runs is empty",
		)
	}

	// Get music shelf and carousels
	musicShelf := external.TentacledMusicShelfRenderer{}
	carousels := []external.MusicCarouselShelfRenderer{}
	for _, content := range tabContents {

		if msr := content.MusicShelfRenderer; msr != nil {
			musicShelf = *msr
		}
		if mcsr := content.MusicCarouselShelfRenderer; mcsr != nil {
			carousels = append(carousels, *mcsr)
		}
	}

	// Get the artist name and thumbnails
	artistName := headersTitleRuns[0].Text
	thumbnails := utils.GetThumbnail(
		headers.MusicImmersiveHeaderRenderer.Thumbnail,
	)

	return models.ArtistProfile{
		Name:               artistName,
		Thumbnails:         thumbnails,
		MusicShelfRenderer: musicShelf,
		Carousels:          carousels,
		VisitorID:          responseContext.VisitorData,
	}, nil

}
