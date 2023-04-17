package song

import (
	"fmt"
	"katze/src/mappers/search/general/utils"
	"katze/src/mappers/utils/columns"
	"katze/src/mappers/utils/simplifier"
	"katze/src/models/external"
	"katze/src/models/music"
)

// Mapper maps an external.MischievousContent to a music.Song
func Mapper(listItemRenderer external.MischievousContent) (
	music.Song, error,
) {

	//simplifyData
	itemRenderer, err := simplifier.MusicResponsiveListItemRenderer(
		listItemRenderer,
	)
	if err != nil {
		return music.Song{}, err
	}
	// Validate the item renderer
	err = utils.ValidateItemRenderer(itemRenderer, "MUSIC_TYPE", 2)
	if err != nil {
		return music.Song{}, err
	}

	if len(itemRenderer.FlexColumns) != 2 {
		err := fmt.Errorf(
			"error invalid data, the flex columns should have 2 items",
		)
		return music.Song{}, err
	}
	titleColumn := itemRenderer.FlexColumns[0]
	infoColumn := itemRenderer.FlexColumns[1]

	title := titleColumn.Items[0].Text
	artists, err := columns.GetArtists(infoColumn.Items)
	if err != nil {
		return music.Song{}, err
	}
	albums, err := columns.GetAlbums(infoColumn.Items)
	if err != nil {
		return music.Song{}, err
	}

	duration, err := columns.GetDuration(infoColumn.Items)
	if err != nil {
		return music.Song{}, err
	}

	return music.Song{
		Title:      title,
		ID:         itemRenderer.ID,
		Artists:    artists,
		AlbumTitle: &albums[0].Title,
		AlbumID:    &albums[0].ID,
		Duration:   &duration,
		Explicit:   &itemRenderer.Explicit,
		Thumbnails: itemRenderer.Thumbnails,
	}, nil
}
