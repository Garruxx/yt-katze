package utils

import (
	"fmt"
	"katze/src/models/external"
)

// GetBestResult returns the best result in the type extenral.MusicCardShelfRenderer.
func GetBestResult(
	tabContents []external.MagentaContent,
) (external.MusicCardShelfRenderer, error) {

	// If there are no tab contents, return an empty shelf
	if len(tabContents) == 0 {
		err := fmt.Errorf("No tab contents found")
		return external.MusicCardShelfRenderer{}, err
	}

	// Loop through the tab contents
	for _, tabContent := range tabContents {

		if mcsr := tabContent.MusicCardShelfRenderer; mcsr != nil {

			return *mcsr, nil
		}
	}

	err := fmt.Errorf("cloud not find best result")
	return external.MusicCardShelfRenderer{}, err
}
