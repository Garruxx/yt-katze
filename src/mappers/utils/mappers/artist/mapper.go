package artist

import (
	"katze/src/mappers/search/general/utils"
	"katze/src/mappers/search/utils/simplify"
	"katze/src/models/external"
	"katze/src/models/music"
)

func Mapper(itemRenderer external.MischievousContent) (
	music.Artist, error,
) {

	artistItemRenderer, err := simplify.MusicResponsiveListItemRenderer(
		itemRenderer,
	)

	if err != nil {
		return music.Artist{}, err
	}

	// Validate the item renderer
	err = utils.ValidateItemRenderer(
		artistItemRenderer, "MUSIC_PAGE_TYPE_ARTIST", 2,
	)
	if err != nil {
		return music.Artist{}, err
	}

	titleColumn := artistItemRenderer.FlexColumns[0]
	artistName := titleColumn.Items[0].Text
	thumbnails := artistItemRenderer.Thumbnails
	return music.Artist{
		Name:       artistName,
		ID:         artistItemRenderer.ID,
		Thumbnails: thumbnails,
	}, nil
}
