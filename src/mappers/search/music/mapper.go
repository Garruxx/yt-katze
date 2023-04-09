package music

import (
	"fmt"
	"katze/src/mappers/utils/mappers"
	"katze/src/models/external"
	"katze/src/models/music"
	"katze/src/models/shelves"
)

func Mapper(musicResults external.MusicList) (shelves.Music, error) {

	// Check if the musicResults is empty
	results := musicResults.Contents
	if results == nil {
		err := fmt.Errorf("error: musicResults.Contents is empty")
		return shelves.Music{}, err
	}

	// Check if the musicResults.ResponseContext is empty
	if musicResults.ResponseContext == nil {
		err := fmt.Errorf("error: musicResults.ResponseContext is empty")
		return shelves.Music{}, err
	}

	playListShelf := results.TabbedSearchResultsRenderer.Tabs[0].TabRenderer.Content.SectionListRenderer.Contents[0].MusicShelfRenderer
	musicShelf := playListShelf.Contents

	if leng := len(musicShelf); leng == 0 {
		return shelves.Music{}, nil
	}

	var musicItems []music.Song
	for _, item := range musicShelf {

		song, err := mappers.Song(item)
		if err != nil {
			return shelves.Music{}, err
		}
		musicItems = append(musicItems, song)
	}

	// Get the visitorID is it exists
	visitorID := musicResults.ResponseContext.VisitorData
	
	continuationID := playListShelf.Continuations[0].NextContinuationData.Continuation

	return shelves.Music{
		Tracks:         musicItems,
		ContinuationID: continuationID,
		VisitorID:      visitorID,
	}, nil
}
