package pagination

import (
	"fmt"
	"katze/src/mappers/utils/mappers"
	"katze/src/models/external"
	"katze/src/models/shelves"
)

func Map(paginationResult external.MusicPagination) (shelves.Music, error) {

	pagination := paginationResult.ContinuationContents
	if pagination == nil {
		err := fmt.Errorf("error: paginationResult is empty")
		return shelves.Music{}, err
	}

	// Check if the musicResults.ResponseContext is empty
	responseContext := paginationResult.ResponseContext
	if responseContext == nil {
		err := fmt.Errorf("error: paginationResult.ResponseContext is empty")
		return shelves.Music{}, err
	}

	playListShelf := pagination.MusicShelfContinuation.Contents
	if leng := len(playListShelf); leng == 0 {
		return shelves.Music{}, nil
	}

	musicItems, err := mappers.Songs(playListShelf)
	if err != nil {
		return shelves.Music{}, err
	}
	// Get the visitorID is it exists
	visitorID := responseContext.VisitorData
	continuationID := pagination.MusicShelfContinuation.Continuations[0].NextContinuationData.Continuation

	return shelves.Music{
		Tracks:         musicItems,
		ContinuationID: continuationID,
		VisitorID:      visitorID,
	}, nil
}
