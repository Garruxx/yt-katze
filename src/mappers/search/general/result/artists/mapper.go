package artists

import (
	"fmt"
	"katze/src/mappers/search/general/result/artists/formatters"
	"katze/src/mappers/search/utils/simplify"
	"katze/src/models"
	"katze/src/models/external"
	"katze/src/models/lists"
	"katze/src/models/music"
)

func Mapper(musicShelfRenderer external.FluffyMusicShelfRenderer) (
	lists.Artists, error,
) {
	// Check if the musicShelfRenderer is empty
	if len(musicShelfRenderer.Contents) == 0 {
		err := fmt.Errorf("error: musicShelfRenderer is empty")
		return lists.Artists{}, err
	}
	// Check if the shelf title is Artists
	if musicShelfRenderer.Title.Runs[0].Text != "Artists" {
		err := fmt.Errorf("error: shelf title is not Artists")
		return lists.Artists{}, err
	}

	// Loop through the musicShelfRenderer contents and get the artists
	artists := []music.Artist{}
	for _, item := range musicShelfRenderer.Contents {

		simpifyData, err := simplify.MusicResponsiveListItemRenderer(item)
		if err != nil {
			return lists.Artists{}, err
		}

		artist, err := formatters.Artist(simpifyData)
		if err != nil {
			return lists.Artists{}, err
		}
		artists = append(artists, artist)
	}

	searchEndpoint := musicShelfRenderer.BottomEndpoint.SearchEndpoint
	continuation := models.Continuation(searchEndpoint)
	return lists.Artists{
		Artists:      artists,
		Continuation: continuation,
	}, nil
}
