package artists

import (
	"katze/src/mappers/utils/mappers/artist"
	"katze/src/models/external"
	"katze/src/models/music"
)

func Mapper(artistsList []external.MischievousContent) (
	[]music.Artist, error,
) {

	artists := []music.Artist{}
	for _, item := range artistsList {
		artist, err := artist.Mapper(item)
		if err != nil {
			return []music.Artist{}, err
		}
		artists = append(artists, artist)
	}
	return artists, nil
}
