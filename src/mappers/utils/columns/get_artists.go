package columns

import (
	"fmt"
	"katze/src/models"
	"katze/src/models/music"
)

// GetArtist extract the artists from the columns
func GetArtists(flexcolumns []models.FlexColumn) ([]music.Artist, error) {

	var artists []music.Artist
	for _, column := range flexcolumns {
		if column.PageType == "MUSIC_PAGE_TYPE_ARTIST" {

			if column.WatchID != "" {
				err := fmt.Errorf(
					"error invalid data, the watch id should be empty",
				)
				return []music.Artist{}, err
			}

			if column.BrowseID == "" {
				err := fmt.Errorf(
					"error invalid data, the browse id should not be empty",
				)
				return []music.Artist{}, err
			}

			artist := music.Artist{
				Name: column.Text,
				ID:   column.BrowseID,
			}
			artists = append(artists, artist)
		}
	}

	return artists, nil
}
