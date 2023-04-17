package columns

import (
	"fmt"
	"katze/src/models"
	"katze/src/models/music"
)

// GetAlbums gets the albums from the flexcolumns
func GetAlbums(flexcolumns []models.FlexColumn) (
	[]music.Album, error,
) {
	var albums []music.Album
	for _, column := range flexcolumns {
		if column.PageType == nil ||
			*column.PageType != "MUSIC_PAGE_TYPE_ALBUM" {
			continue
		}

		if column.WatchID != nil {
			err := fmt.Errorf(
				"error invalid data, the watch id should be empty",
			)
			return []music.Album{}, err
		}

		if column.BrowseID == nil {
			err := fmt.Errorf(
				"error invalid data, the browse id should not be empty",
			)
			return []music.Album{}, err
		}

		if column.Text == "" {
			err := fmt.Errorf(
				"error invalid data, the text should not be empty",
			)
			return []music.Album{}, err
		}

		album := music.Album{
			Title: column.Text,
			ID:    *column.BrowseID,
		}
		albums = append(albums, album)

	}

	if len(albums) == 0 {
		return []music.Album{
			{
				Title: "Unknown",
				ID:    "Unknown",
			},
		}, nil
	}

	return albums, nil
}
