package artist

import (
	"katze/src/mappers/search/general/utils"
	"katze/src/mappers/utils/simplifier"
	"katze/src/models/external"
	"katze/src/models/music"
)

// Mapper maps an external.MischievousContent to a music.Artist
func Mapper(itemRenderer external.MischievousContent) (
	music.Artist, error,
) {

	artistItemRenderer, err := simplifier.MusicResponsiveListItemRenderer(
		itemRenderer,
	)

	if err != nil {
		return music.Artist{}, err
	}

	// Validate the item renderer
	err = utils.ValidateItemRenderer(
		artistItemRenderer, "MUSIC_PAGE_TYPE_ARTIST", 2,
	)
	errorAcepted := "error: pageType is not MUSIC_PAGE_TYPE_ARTIST is MUSIC_PAGE_TYPE_USER_CHANNEL"
	if err != nil && err.Error() != errorAcepted {
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
