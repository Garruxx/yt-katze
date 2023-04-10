package list

import (
	"fmt"
	"katze/src/mappers/utils/mappers"
	"katze/src/models/external"
	"katze/src/models/lists"
)

func Map(artistResult external.ArtistList) (lists.Artists, error) {

	// Check if the artistResult is empty
	results := artistResult.Contents
	if results == nil {
		err := fmt.Errorf("error: artistResult.Contents is empty")
		return lists.Artists{}, err
	}

	// Check if the artistResult.ResponseContext is empty
	if artistResult.ResponseContext == nil {
		err := fmt.Errorf("error: artistResult.ResponseContext is empty")
		return lists.Artists{}, err
	}

	// Get the artistShelf
	artistShelfRenderer := results.TabbedSearchResultsRenderer.Tabs[0].
		TabRenderer.Content.SectionListRenderer.Contents[0].MusicShelfRenderer
	artistsShelf := artistShelfRenderer.Contents

	artistList, err := mappers.Artists(artistsShelf)
	if err != nil {
		return lists.Artists{}, err
	}

	// Get the visitorID is it exists
	visitorID := artistResult.ResponseContext.VisitorData
	continuation := artistShelfRenderer.Continuations[0].NextContinuationData.Continuation

	return lists.Artists{
		Artists:        artistList,
		ContinuationID: continuation,
		VisitorID:      visitorID,
	}, nil
}
