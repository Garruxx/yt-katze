package subtitle

import (
	"fmt"
	"katze/src/models"
	"katze/src/models/external"
)

func RunToFlexColumn(run external.TentacledRun) (models.FlexColumn, error) {

	browseID := ""
	pageType := ""
	text := run.Text

	if run.NavigationEndpoint != nil {
		browseEndpoint := run.NavigationEndpoint.BrowseEndpoint
		supportedConfigs := browseEndpoint.BrowseEndpointContextSupportedConfigs
		browseID = browseEndpoint.BrowseID
		pageType = supportedConfigs.BrowseEndpointContextMusicConfig.PageType

		if browseID == "" {
			err := fmt.Errorf("Could not find browseID in flexcolumn")
			return models.FlexColumn{}, err
		}
		if pageType == "" {
			err := fmt.Errorf("Could not find pageType in flexcolumn")
			return models.FlexColumn{}, err
		}
	}

	if text == "" {
		err := fmt.Errorf("Could not find text in flexcolumn")
		return models.FlexColumn{}, err
	}

	return models.FlexColumn{
		Text:     text,
		BrowseID: browseID,
		PageType: pageType,
	}, nil
}
