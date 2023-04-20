package card

import (
	"fmt"
	"katze/src/mappers/browse/utils/mappers"
	"katze/src/models/artist"
	"katze/src/models/external"
	"katze/src/utils"
)

func TwoRowItem(albumsData external.ArtistTwoRowItem) (
	albums []artist.CardItem, err error,
) {

	if albumsData.Error != nil {
		return albums, fmt.Errorf("error: %s", albumsData.Error.Message)
	}
	resContecxt := albumsData.ResponseContext
	if resContecxt == nil {
		err := fmt.Errorf("error: response context is empty")
		return albums, err
	}
	albumsTabs := albumsData.Contents.SingleColumnBrowseResultsRenderer.Tabs
	if len(albumsTabs) == 0 {
		err := fmt.Errorf("error: tabs is empty")
		return albums, err
	}

	albumsTabContents := albumsTabs[0].TabRenderer.Content.SectionListRenderer.Contents

	// This problem is caused by a wrong visitorID
	if len(albumsTabContents) == 0 && resContecxt.VisitorData != "" {
		return albums, fmt.Errorf(
			"error possibly caused by the visitor ID in the request",
		)
	}

	if len(albumsTabContents) == 0 {

		err := fmt.Errorf("error: tab contents is empty")
		return albums, err
	}

	albumsItems := albumsTabContents[0].GridRenderer.Items
	if len(albumsItems) == 0 {
		err := fmt.Errorf("error: albums items is empty")
		return albums, err
	}

	return utils.ArrayMap(albumsItems, mappers.TwoRowItemRenderer)
}
