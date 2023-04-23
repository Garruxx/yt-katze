package recomendations

import (
	"fmt"
	"katze/src/mappers/browse/utils/mappers"
	"katze/src/models/external"
	"katze/src/models/music"
	"katze/src/utils"
)

func SongUpNext(recomendations external.MusicRecomendations) (
	[]music.Song, error,
) {
	if recomendations.Error != nil {
		return []music.Song{}, fmt.Errorf(
			"error getting recomendations: %v", recomendations.Error.Message,
		)
	}
	if recomendations.Contents == nil {
		return []music.Song{}, fmt.Errorf("no contents found")
	}
	if recomendations.ResponseContext == nil {
		return []music.Song{}, fmt.Errorf("no response context found")
	}
	Carouselcontents := recomendations.Contents.SectionListRenderer.Contents
	if len(Carouselcontents) == 0 {
		return []music.Song{}, fmt.Errorf("no carousel contents found")
	}
	contents := Carouselcontents[0].MusicCarouselShelfRenderer.Contents

	songs, err := utils.ArrayMap(contents, mappers.MischievousSong)
	if err != nil {
		return []music.Song{}, fmt.Errorf("error mapping songs: %v", err)
	}
	return songs, nil
}
