package pagination

import (
	"fmt"
	"katze/src/mappers/browse/utils/mappers"
	"katze/src/models/external"
	"katze/src/models/shelves"
	"katze/src/utils"
)

func Map(
	paginationResult external.ArtistTracklistPagination,
) (
	music shelves.Music, err error,
) {
	if paginationResult.Error != nil {
		return music, fmt.Errorf("error: %s", paginationResult.Error.Message)
	}
	if paginationResult.ResponseContext == nil {
		return music, fmt.Errorf("response context is nil")
	}

	contents := paginationResult.ContinuationContents.MusicPlaylistShelfContinuation.Contents
	if len(contents) == 0 {
		return music, fmt.Errorf("artist music shelf is empty")
	}

	continuations := paginationResult.ContinuationContents.MusicPlaylistShelfContinuation.Continuations
	continuationID := ""
	if len(continuations) != 0 {
		continuationID = continuations[0].NextContinuationData.Continuation
	}

	// Map music shelf to songs
	songs, err := utils.ArrayMap(contents, mappers.Song)
	if err != nil {
		return music, err
	}

	return shelves.Music{
		Songs:          songs,
		ContinuationID: continuationID,
		VisitorID:      paginationResult.ResponseContext.VisitorData,
	}, nil
}
