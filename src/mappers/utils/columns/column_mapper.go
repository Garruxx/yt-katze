package columns

import (
	"fmt"
	"katze/src/models"
	"katze/src/models/external"
)

// ColumnMapper maps an external.FluffyRun to a models.FlexColumn
func ColumnMapper(flexcolumn external.FluffyRun) (models.FlexColumn, error) {

	var browseID *string
	var watchID *string
	var pageType *string 
	text := flexcolumn.Text
	navEndpoint := flexcolumn.NavigationEndpoint

	if text == "" {
		err := fmt.Errorf("could not find text in flexcolumn")
		return models.FlexColumn{}, err
	}

	if navEndpoint == nil {
		return models.FlexColumn{
			Text: text,
		}, nil
	}

	browseEndpoint := navEndpoint.BrowseEndpoint
	watchEndpoint := navEndpoint.WatchEndpoint

	if browseEndpoint != nil {
		supportedConfigs := browseEndpoint.BrowseEndpointContextSupportedConfigs
		browseID = &browseEndpoint.BrowseID
		pageType = &supportedConfigs.BrowseEndpointContextMusicConfig.PageType

		if *browseID == "" {
			err := fmt.Errorf("could not find browseID in flexcolumn")
			return models.FlexColumn{}, err
		}

		if *pageType == "" {
			err := fmt.Errorf("could not find pageType in flexcolumn")
			return models.FlexColumn{}, err
		}
	}

	if watchEndpoint != nil {
		music := "Music"
		watchID = &watchEndpoint.VideoID
		pageType = &music
		if *watchID == "" {
			err := fmt.Errorf("could not find watchID in flexcolumn")
			return models.FlexColumn{}, err
		}
	}

	return models.FlexColumn{
		Text:     text,
		BrowseID: browseID,
		PageType: pageType,
		WatchID:  watchID,
	}, nil
}
