package album

import (
	"katze/src/mappers/search/general/utils"
	"katze/src/mappers/utils/columns"
	"katze/src/mappers/utils/simplifier"
	"katze/src/models/external"
	"katze/src/models/music"
)

// Mapper maps an external.MischievousContent to a music.Album
func Mapper(itemRenderer external.MischievousContent) (
	music.Album, error,
) {

	albumItemRenderer, err := simplifier.MusicResponsiveListItemRenderer(
		itemRenderer,
	)
	if err != nil {
		return music.Album{}, err
	}

	err = utils.ValidateItemRenderer(
		albumItemRenderer, "MUSIC_PAGE_TYPE_ALBUM", 2,
	)
	if err != nil {
		return music.Album{}, err
	}

	titleColumn := albumItemRenderer.FlexColumns[0]
	infoColumn := albumItemRenderer.FlexColumns[1]
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
		ID:         albumItemRenderer.ID,
		Artists:    &artists,
		Thumbnails: &albumItemRenderer.Thumbnails,
		Single:     &single,
		EP:         &ep,
		Year:       &year,
		Explicit:   &albumItemRenderer.Explicit,
	}, nil
}
