package player

import (
	"fmt"
	"katze/src/mappers/models"
	"katze/src/models/external"
)

func Information(playerInfo external.PlayerInformation) (
	models.PlayerInformation, error,
) {

	if playerInfo.Error != nil {
		return models.PlayerInformation{}, fmt.Errorf(
			"error getting player information: %v", playerInfo.Error.Message,
		)
	}
	if playerInfo.Contents == nil {
		return models.PlayerInformation{}, fmt.Errorf("no contents found")
	}
	if playerInfo.ResponseContext == nil {
		return models.PlayerInformation{}, fmt.Errorf("no response context found")
	}

	var info models.PlayerInformation

	tabs := playerInfo.Contents.SingleColumnMusicWatchNextResultsRenderer.
		TabbedRenderer.WatchNextTabbedResultsRenderer.Tabs

	if len(tabs) == 0 {
		return info, fmt.Errorf("no tabs found")
	}

	for _, tab := range tabs {
		browseEndpoint := tab.TabRenderer.Endpoint.BrowseEndpoint

		if tab.TabRenderer.Title == "Lyrics" && browseEndpoint != nil {
			info.Lyrics.BrowseID = browseEndpoint.BrowseID
		}
		if tab.TabRenderer.Title == "Related" {
			if browseEndpoint != nil {
				info.Related.BrowseID = browseEndpoint.BrowseID
			} else {
				return info, fmt.Errorf("no browse endpoint found")
			}
		}
	}
	return info, nil
}
