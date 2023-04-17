package music

import (
	"fmt"
	"katze/src/mappers/utils/mappers"
	"katze/src/models"
	"katze/src/models/external"
	"katze/src/models/lists"
	"katze/src/utils"
)

func Mapper(musicShelfRenderer external.FluffyMusicShelfRenderer) (
	lists.Music, error,
) {

	// Check if the musicShelfRenderer is empty
	if len(musicShelfRenderer.Contents) == 0 {
		err := fmt.Errorf("error: musicShelfRenderer is empty")
		return lists.Music{}, err
	}
	// Check if the shelf title is Songs
	if musicShelfRenderer.Title.Runs[0].Text != "Songs" {
		err := fmt.Errorf("error: shelf title is not Songs")
		return lists.Music{}, err
	}

	// Loop through the musicShelfRenderer contents and get the songs
	songs, err := utils.ArrayMap(musicShelfRenderer.Contents, mappers.Song)
	if err != nil {
		return lists.Music{}, err
	}
	searchEndpoint := musicShelfRenderer.BottomEndpoint.SearchEndpoint
	continuation := models.Continuation(searchEndpoint)

	return lists.Music{
		Songs:        songs,
		Continuation: continuation,
	}, nil
}
