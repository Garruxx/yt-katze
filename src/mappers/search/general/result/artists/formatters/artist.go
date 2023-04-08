package formatters

import (
	"katze/src/mappers/search/general/utils"
	"katze/src/mappers/search/models"
	"katze/src/models/music"
)

func Artist(itemRenderer models.MusicResponsiveListItemRenderer) (
	music.Artist, error,
) {
	// Validate the item renderer
	err := utils.ValidateItemRenderer(itemRenderer, "MUSIC_PAGE_TYPE_ARTIST", 2)
	if err != nil {
		return music.Artist{}, err
	}

	titleColumn := itemRenderer.FlexColumns[0]
	artistName := titleColumn.Items[0].Text
	thumbnails := itemRenderer.Thumbnails
	return music.Artist{
		Name:       artistName,
		ID:         itemRenderer.ID,
		Thumbnails: thumbnails,
	}, nil
}
