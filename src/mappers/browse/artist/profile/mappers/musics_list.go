package mappers

import (
	"fmt"
	"katze/src/models/artist"
	"katze/src/models/external"
	"katze/src/utils"
)

func MusicsList(musicShelf external.TentacledMusicShelfRenderer) (
	artist.MusicsList, error,
) {

	if len(musicShelf.Contents) == 0 {
		err := fmt.Errorf("no music found")
		return artist.MusicsList{}, err
	}

	songs, err := utils.ArrayMap(musicShelf.Contents, Song)
	if err != nil {
		return artist.MusicsList{}, err
	}

	browseEndpoint := musicShelf.BottomEndpoint.BrowseEndpoint
	return artist.MusicsList{
		Songs:              songs,
		ContinuationParams: *browseEndpoint.Params,
		ContinuationID:     browseEndpoint.BrowseID,
	}, nil
}
