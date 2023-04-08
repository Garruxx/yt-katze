package formatters

import (
	"katze/src/mappers/search/general/utils"
	"katze/src/mappers/search/models"
	"katze/src/mappers/utils/columns"
	"katze/src/models/music"
)

func Album(itemRenderer models.MusicResponsiveListItemRenderer) (
	music.Album, error,
) {

	err := utils.ValidateItemRenderer(itemRenderer, "MUSIC_PAGE_TYPE_ALBUM", 2)
	if err != nil {
		return music.Album{}, err
	}

	titleColumn := itemRenderer.FlexColumns[0]
	infoColumn := itemRenderer.FlexColumns[1]
	title := titleColumn.Items[0].Text
	single := infoColumn.Items[0].Text == "Single"
	ep := infoColumn.Items[0].Text == "EP"

	artists, err := columns.GetArtists(infoColumn.Items)
	if err != nil {
		return music.Album{}, err
	}
	year, err := columns.GetYear(infoColumn.Items)
	if err != nil {
		return music.Album{}, err
	}

	return music.Album{
		Title:      title,
		ID:         itemRenderer.ID,
		Artists:    &artists,
		Thumbnails: &itemRenderer.Thumbnails,
		Single:     &single,
		EP:         &ep,
		Year:       &year,
		Explicit:   &itemRenderer.Explicit,
	}, nil
}
