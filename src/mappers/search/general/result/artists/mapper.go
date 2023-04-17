package artists

import (
	"fmt"
	"katze/src/mappers/utils/mappers"
	"katze/src/models"
	"katze/src/models/external"
	"katze/src/models/lists"
	"katze/src/utils"
)

// Mapper maps an external.MusicShelfRenderer to a lists.Albums
func Mapper(musicShelfRenderer external.FluffyMusicShelfRenderer) (
	lists.Artists, error,
) {
	// Check if the musicShelfRenderer is empty
	if len(musicShelfRenderer.Contents) == 0 {
		err := fmt.Errorf("error: musicShelfRenderer is empty")
		return lists.Artists{}, err
	}

	// Check if the musicShelfRenderer title runs is not empty
	if len(musicShelfRenderer.Title.Runs) == 0 {
		err := fmt.Errorf("error: musicShelfRenderer title runs is empty")
		return lists.Artists{}, err
	}
	// Check if the shelf title is Artists
	if musicShelfRenderer.Title.Runs[0].Text != "Artists" {
		err := fmt.Errorf("error: shelf title is not Artists")
		return lists.Artists{}, err
	}

	// Loop through the musicShelfRenderer contents and get the artists
	artists, err := utils.ArrayMap(musicShelfRenderer.Contents, mappers.Artist)
	if err != nil {
		return lists.Artists{}, err
	}

	searchEndpoint := musicShelfRenderer.BottomEndpoint.SearchEndpoint
	continuation := models.Continuation(searchEndpoint)
	return lists.Artists{
		Artists:      artists,
		Continuation: continuation,
	}, nil
}
