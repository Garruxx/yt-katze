package simplifier

import (
	"fmt"
	"katze/src/mappers/browse/models"
	"katze/src/models/external"
)

// CarouselShelfRenderer maps
// an external.MusicCarouselShelfRenderer to a models.CarouselShelfRenderer
func CarouselShelfRenderer(
	shelfCarousel external.MusicCarouselShelfRenderer,
) (
	carousel models.CarouselShelfRenderer, err error,
) {

	// Verfiy that the header is not nil and that it has a title
	header := shelfCarousel.Header.MusicCarouselShelfBasicHeaderRenderer
	titleRuns := header.Title.Runs
	if len(titleRuns) == 0 {
		err = fmt.Errorf(
			"error mapping carousel shelf renderer: title runs is empty %v", titleRuns,
		)
		return carousel, err
	}

	// Get the next page token if it exists
	browseID := ""
	params := ""
	if moreButton := header.MoreContentButton; moreButton != nil {
		endpoints := moreButton.ButtonRenderer.NavigationEndpoint.BrowseEndpoint
		browseID = endpoints.BrowseID
		params = *endpoints.Params
	}

	title := titleRuns[0].Text

	carousel = models.CarouselShelfRenderer{
		Title:    title,
		BrowseID: browseID,
		Params:   params,
		Contents: shelfCarousel.Contents,
	}

	return carousel, err
}
