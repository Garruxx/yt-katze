package mappers

import (
	"fmt"
	"katze/src/mappers/utils/columns"
	"katze/src/mappers/utils/simplifier"
	"katze/src/models/external"
	"katze/src/models/music"
)

func MischievousSong(
	song external.MischievousContent,
) (
	music.Song, error,
) {

	itemRenderer, err := simplifier.MusicResponsiveListItemRenderer(song)
	if err != nil {
		return music.Song{}, err
	}

	if leng := len(itemRenderer.FlexColumns); leng != 3 {
		return music.Song{}, fmt.Errorf("expected 3 flex columns, got %v", leng)
	}

	artists, err := columns.GetArtists(itemRenderer.FlexColumns[1].Items)
	if err != nil {
		return music.Song{}, err
	}

	albums, err := columns.GetAlbums(itemRenderer.FlexColumns[2].Items)
	if err != nil {
		return music.Song{}, err
	}
	title := itemRenderer.FlexColumns[0].Items[0].Text

	return music.Song{
		Title:      title,
		ID:         itemRenderer.ID,
		Artists:    artists,
		AlbumTitle: &albums[0].Title,
		AlbumID:    &albums[0].ID,
		Thumbnails: itemRenderer.Thumbnails,
		Explicit:   &itemRenderer.Explicit,
	}, nil
}
