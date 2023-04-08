package utils

import (
	"fmt"
	"katze/src/models/external"
)

func ShelfTypeValidator(bestResult external.MusicCardShelfRenderer) (
	string, error,
) {

	if len(bestResult.Title.Runs) == 0 {
		err := fmt.Errorf("the shelf does not have a title")
		return "EMPTY", err
	}

	if bestResult.Title.Runs[0].NavigationEndpoint == nil {
		err := fmt.Errorf(
			"the shelf does not have a navigation endpoint",
		)
		return "", err
	}

	navEndpoint := bestResult.Title.Runs[0].NavigationEndpoint
	if navEndpoint.BrowseEndpoint == nil && navEndpoint.WatchEndpoint == nil {

		err := fmt.Errorf(
			"the shelf does not have a browse or watch endpoint",
		)
		return "", err
	}

	// If the best match is a music video, return it
	if navEndpoint.BrowseEndpoint == nil {

		if navEndpoint.WatchEndpoint.VideoID == "" {
			err := fmt.Errorf("the shelf does not have a videoID")
			return "", err
		}
		shelfType := navEndpoint.WatchEndpoint.WatchEndpointMusicSupportedConfigs.WatchEndpointMusicConfig.MusicVideoType
		if navEndpoint.WatchEndpoint != nil {
			return shelfType, nil
		}
	}

	// If the best match is an album or an artist, return it
	if navEndpoint.WatchEndpoint == nil {

		shelfType := bestResult.Title.Runs[0].NavigationEndpoint.
			BrowseEndpoint.BrowseEndpointContextSupportedConfigs.
			BrowseEndpointContextMusicConfig.PageType

		if shelfType == "" {
			err := fmt.Errorf(
				"the shelf does not have a page type",
			)
			return "", err
		}
		if navEndpoint.BrowseEndpoint.BrowseID == "" {
			err := fmt.Errorf("the shelf does not have a browseID")
			return "", err
		}
		if navEndpoint.BrowseEndpoint != nil {
			return shelfType, nil
		}
	}

	return "UNKNOWN", nil
}
