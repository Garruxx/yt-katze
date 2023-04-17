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
		if column.PageType == nil ||
			*column.PageType != "MUSIC_PAGE_TYPE_ARTIST" {
			continue
		}

		if column.WatchID != nil {
			return []music.Artist{}, fmt.Errorf(
				"error invalid data, the watch id should be empty",
			)
		}

		if column.BrowseID == nil {
			return []music.Artist{}, fmt.Errorf(
				"error invalid data, the browse id should not be empty",
			)
		}

		artists = append(artists, music.Artist{
			Name: column.Text,
			ID:   *column.BrowseID,
		})
	}

	return artists, nil
}
