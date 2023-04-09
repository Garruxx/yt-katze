package utils

import (
	"fmt"
	"katze/src/models/external"
)

// FindShelf finds the shelf that contains the query
func FindCardShelf(
	tabContents []external.MagentaContent,
	query string,
) (external.FluffyMusicShelfRenderer, error) {

	// If there are no tab contents, return an error
	if len(tabContents) == 0 {
		err := fmt.Errorf("no tab contents found")
		return external.FluffyMusicShelfRenderer{}, err
	}

	// Loop through the tab contents
	// and find the shelf that contains the query
	for _, tabContent := range tabContents {

		if msr := tabContent.MusicShelfRenderer; msr != nil {

			switch title := msr.Title.Runs[0].Text; title {
			case query:
				return *msr, nil
			}
		}
	}

	err := fmt.Errorf("could not find shelf")
	return external.FluffyMusicShelfRenderer{}, err
}
