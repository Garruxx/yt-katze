package songs

import (
	"katze/src/mappers/utils/mappers/song"
	"katze/src/models/external"
	"katze/src/models/music"
)

func Mapper(listItemsRenderer []external.MischievousContent) (
	[]music.Song, error,
) {

	songs := []music.Song{}
	for _, item := range listItemsRenderer {
		song, err := song.Mapper(item)
		if err != nil {
			return []music.Song{}, err
		}
		songs = append(songs, song)
	}
	return songs, nil
}
