package albums

import (
	"fmt"
	"katze/src/mappers/utils/mappers"
	"katze/src/models"
	"katze/src/models/external"
	"katze/src/models/lists"
	"katze/src/utils"
)

// Mapper maps an external.FluffyMusicShelfRenderer to a lists.Albums
func Mapper(musicShelfRenderer external.FluffyMusicShelfRenderer) (
	lists.Albums, error,
) {

	// Check if the musicShelfRenderer is empty
	if len(musicShelfRenderer.Contents) == 0 {
		err := fmt.Errorf("error: musicShelfRenderer is empty")
		return lists.Albums{}, err
	}

	// Check if the musicShelfRenderer title runs is empty
	titleRuns := musicShelfRenderer.Title.Runs
	if len(titleRuns) == 0 {
		err := fmt.Errorf("error: musicShelfRenderer title runs is empty")
		return lists.Albums{}, err
	}

	// Check if the shelf title is Albums
	if titleRuns[0].Text != "Albums" {
		err := fmt.Errorf("error: shelf title is not Albums")
		return lists.Albums{}, err
	}

	// Loop through the musicShelfRenderer contents and get the albums
	albums, err := utils.ArrayMap(musicShelfRenderer.Contents, mappers.Album)
	if err != nil {
		return lists.Albums{}, err
	}

	searchEndpoint := musicShelfRenderer.BottomEndpoint.SearchEndpoint
	continuation := models.Continuation(searchEndpoint)
	return lists.Albums{
		Albums:       albums,
		Continuation: continuation,
	}, nil
}
