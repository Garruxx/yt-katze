package list

import (
	"fmt"
	"katze/src/mappers/browse/utils/mappers"
	"katze/src/models/external"
	"katze/src/models/shelves"
	"katze/src/utils"
)

func Map(artistMusicList external.ArtistTracklist) (
	music shelves.Music, err error,
) {
	if artistMusicList.Error != nil {
		return music, fmt.Errorf("error: %v", artistMusicList.Error)
	}
	if artistMusicList.ResponseContext == nil {
		return music, fmt.Errorf("error: response context is nil")
	}

	// Validate tabb, tab contents and music shelf are not empty
	tabs := artistMusicList.Contents.SingleColumnBrowseResultsRenderer.Tabs
	if len(tabs) == 0 {
		return music, fmt.Errorf("error: artist tabs is empty")
	}

	tabContents := tabs[0].TabRenderer.Content.SectionListRenderer.Contents
	if len(tabContents) == 0 {
		return music, fmt.Errorf("error: artist tab contents is empty")
	}

	playlistShelf := tabContents[0].MusicPlaylistShelfRenderer
	musicShelf := playlistShelf.Contents
	if len(musicShelf) == 0 {
		return music, fmt.Errorf("error: artist music shelf is empty")
	}

	continuations := playlistShelf.Continuations
	continuationID := ""
	if len(continuations) != 0 {
		continuationID = continuations[0].NextContinuationData.Continuation
	}

	// Map music shelf to songs
	songs, err := utils.ArrayMap(musicShelf, mappers.Song)
	if err != nil {
		return music, err
	}

	return shelves.Music{
		Songs:          songs,
		ContinuationID: continuationID,
		VisitorID:      artistMusicList.ResponseContext.VisitorData,
	}, nil
}
