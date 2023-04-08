package albums

import (
	"fmt"
	"katze/src/mappers/search/general/result/albums/formatters"
	"katze/src/mappers/search/utils/simplify"
	"katze/src/models"
	"katze/src/models/external"
	"katze/src/models/lists"
	"katze/src/models/music"
)

func Mapper(musicShelfRenderer external.FluffyMusicShelfRenderer) (
	lists.Albums, error,
) {

	// Check if the musicShelfRenderer is empty
	if len(musicShelfRenderer.Contents) == 0 {
		err := fmt.Errorf("error: musicShelfRenderer is empty")
		return lists.Albums{}, err
	}

	// Check if the shelf title is Albums
	if musicShelfRenderer.Title.Runs[0].Text != "Albums" {
		err := fmt.Errorf("error: shelf title is not Albums")
		return lists.Albums{}, err
	}

	// Loop through the musicShelfRenderer contents and get the albums
	albums := []music.Album{}
	for _, item := range musicShelfRenderer.Contents {

		simplifyData, err := simplify.MusicResponsiveListItemRenderer(item)
		if err != nil {
			return lists.Albums{}, err
		}

		album, err := formatters.Album(simplifyData)
		if err != nil {
			return lists.Albums{}, err
		}
		albums = append(albums, album)
	}

	searchEndpoint := musicShelfRenderer.BottomEndpoint.SearchEndpoint
	continuation := models.Continuation(searchEndpoint)
	return lists.Albums{
		Albums:       albums,
		Continuation: continuation,
	}, nil
}
