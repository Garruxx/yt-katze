package formatters

import (
	"katze/src/mappers/search/general/utils"
	"katze/src/mappers/search/models"
	"katze/src/mappers/utils/columns"
	"katze/src/models/music"
)

func Song(itemRenderer models.MusicResponsiveListItemRenderer) (
	music.Song, error,
) {
	// Validate the item renderer
	err := utils.ValidateItemRenderer(itemRenderer, "MUSIC_TYPE", 2)
	if err != nil {
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
