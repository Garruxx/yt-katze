package mappers

import (
	"fmt"
	"katze/src/models/artist"
	"katze/src/models/external"
	"katze/src/utils"
)

func MusicList(musicShelf external.TentacledMusicShelfRenderer) (
	artist.MusicList, error,
) {

	if len(musicShelf.Contents) == 0 {
		err := fmt.Errorf("no music found")
		return artist.MusicList{}, err
	}

	songs, err := utils.ArrayMap(musicShelf.Contents, Song)
	if err != nil {
		return artist.MusicList{}, err
	}

	browseEndpoint := musicShelf.BottomEndpoint.BrowseEndpoint
	return artist.MusicList{
		Songs:              songs,
		ContinuationParams: *browseEndpoint.Params,
		ContinuationID:     browseEndpoint.BrowseID,
	}, nil
}
