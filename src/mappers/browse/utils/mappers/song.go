package mappers

import (
	"fmt"
	"katze/src/mappers/utils"
	"katze/src/mappers/utils/columns"
	"katze/src/models"
	"katze/src/models/external"
	"katze/src/models/music"
	"strconv"
)

func Song(
	songPlaylistShelf external.MusicPlaylistShelfContinuationContent,
) (songItem music.Song, err error) {

	songContent := songPlaylistShelf.MusicResponsiveListItemRenderer
	// Map flex columns
	flexColumns, err := columns.Mapper(songContent.FlexColumns)
	if err != nil {
		return music.Song{}, err
	}
	if len(flexColumns) > 3 || len(flexColumns) < 1 {
		return music.Song{}, fmt.Errorf(
			"expected 1 to 3  flex columns, got %v", len(flexColumns),
		)
	}

	// Map fixed columns (only one)
	// errors validations
	fixedColumns := songContent.FixedColumns
	if len(fixedColumns) == 0 {
		return music.Song{}, fmt.Errorf("fixed columns is empty")
	}

	fixedColumnRuns := fixedColumns[0].MusicResponsiveListItemFixedColumnRenderer.Text.Runs
	if len(fixedColumnRuns) <= 0 {
		return music.Song{}, fmt.Errorf("fixed columns runs is empty")
	}

	// fixed flex column is a fixedcolumn mapped as a flexcolumn
	fixedFlexColumn, err := columns.ColumnMapper(fixedColumnRuns[0])
	if err != nil {
		return music.Song{}, err
	}

	// Get duration
	durationColumn := []models.FlexColumn{fixedFlexColumn}
	duration, err := columns.GetDuration(durationColumn)
	if err != nil {
		return music.Song{}, err
	}
	// validate if the song has an album and artists
	var albumTitle string
	var albumId string
	artists := []music.Artist{}
	if len(flexColumns) == 3 {

		// Get artists
		artists, err = columns.GetArtists(flexColumns[1].Items)
		if err != nil {
			return music.Song{}, err
		}

		// Get albums
		albums, err := columns.GetAlbums(flexColumns[2].Items)
		if err != nil {
			return music.Song{}, err
		}
		albumTitle = albums[0].Title
		albumId = albums[0].ID
	}

	// Get track number (if exists)
	var trackNumner int
	if runs := songContent.Index.Runs; len(runs) > 0 {
		trackNumner, err = strconv.Atoi(runs[0].Text)
		if err != nil {
			return music.Song{}, err
		}
	}

	title := flexColumns[0].Items[0].Text
	Id := flexColumns[0].Items[0].WatchID
	explicit := songContent.Badges != nil

	return music.Song{
		Title:       title,
		ID:          *Id,
		Artists:     artists,
		AlbumTitle:  &albumTitle,
		AlbumID:     &albumId,
		TrackNumber: &trackNumner,
		Thumbnails:  utils.GetThumbnail(songContent.Thumbnail),
		Explicit:    &explicit,
		Duration:    &duration,
	}, nil
}
