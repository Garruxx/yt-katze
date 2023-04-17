package best

import (
	"katze/src/mappers/search/general/result/best/formatters"
	"katze/src/mappers/search/general/result/best/utils"
	"katze/src/models/external"
	"katze/src/models/music"
	"strings"
)

// Mapper maps an external.MusicCardShelfRenderer to a music.BestMatch
func Mapper(bestMatch external.MusicCardShelfRenderer) (
	music.BestMatch, error,
) {

	bestMatchType, err := utils.ShelfTypeValidator(bestMatch)

	if strings.Contains(bestMatchType, "MUSIC_VIDEO") {
		return formatters.Song(bestMatch)
	}

	if err != nil {
		if err.Error() == "the shelf does not have a title" {
			return music.BestMatch{}, nil
		}
		return music.BestMatch{}, err
	}

	if bestMatchType == "MUSIC_PAGE_TYPE_ALBUM" {
		return formatters.Album(bestMatch)
	}
	if bestMatchType == "MUSIC_PAGE_TYPE_ARTIST" {
		return formatters.Artist(bestMatch)
	}

	return music.BestMatch{}, nil
}
