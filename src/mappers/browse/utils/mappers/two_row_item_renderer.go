package mappers

import (
	"fmt"
	"katze/src/mappers/utils"
	"katze/src/mappers/utils/columns"
	"katze/src/mappers/utils/subtitle"
	"katze/src/models/artist"
	"katze/src/models/external"
)

// TwoRowItemRenderer maps a two row item renderer to a card item
func TwoRowItemRenderer(
	twoRowItem external.MusicCarouselShelfRendererContent,
) (
	artist.CardItem, error,
) {

	// Validate the two row item renderer is not empty
	titleRuns := twoRowItem.MusicTwoRowItemRenderer.Title.Runs
	title := ""
	if len(titleRuns) <= 0 {
		return artist.CardItem{}, fmt.Errorf(
			"error mapping two row item renderer: title runs is empty %v",
			titleRuns,
		)
	}
	browseEndpoint := twoRowItem.MusicTwoRowItemRenderer.NavigationEndpoint.BrowseEndpoint
	if browseEndpoint == nil {
		return artist.CardItem{}, fmt.Errorf(
			"error mapping two row item renderer: browseEndpoint is empty %v",
			browseEndpoint,
		)
	}

	// To flex columns
	flexColumn, err := subtitle.ToFlexColumns(
		twoRowItem.MusicTwoRowItemRenderer.Subtitle,
	)
	if err != nil {
		return artist.CardItem{}, err
	}

	// Get the data
	Year, err := columns.GetYear(flexColumn)
	if err != nil {
		return artist.CardItem{}, err
	}
	title = titleRuns[0].Text
	explicit := twoRowItem.MusicTwoRowItemRenderer.SubtitleBadges != nil

	return artist.CardItem{
		Title:    title,
		Year:     Year,
		Explicit: explicit,
		BrowseID: browseEndpoint.BrowseID,
		Params:   *browseEndpoint.Params,
		Thumbnails: utils.GetThumbnail(
			twoRowItem.MusicTwoRowItemRenderer.ThumbnailRenderer,
		),
	}, nil
}
