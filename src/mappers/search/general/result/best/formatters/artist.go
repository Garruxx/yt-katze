package formatters

import (
	"fmt"
	bestUtils "katze/src/mappers/search/general/result/best/utils"
	"katze/src/mappers/utils"
	"katze/src/models/external"
	"katze/src/models/music"
	"strings"
)

// Artist maps an external.MusicCardShelfRenderer to a music.BestMatch
func Artist(bestMatch external.MusicCardShelfRenderer) (
	music.BestMatch, error,
) {

	bestMatchType, err := bestUtils.ShelfTypeValidator(bestMatch)
	if err != nil {
		return music.BestMatch{}, err
	}
	if len(bestMatch.Title.Runs) == 0 {
		err := fmt.Errorf("the best result does not have a title")
		return music.BestMatch{}, err
	}
	titleData := bestMatch.Title.Runs[0]

	if !strings.Contains(bestMatchType, "MUSIC_PAGE_TYPE_ARTIST") {
		err := fmt.Errorf("the best result is not of the artist type")
		return music.BestMatch{}, err
	}

	thumbnails := utils.GetThumbnail(bestMatch.Thumbnail)

	return music.BestMatch{
		Type:       "artist",
		Title:      titleData.Text,
		ID:         titleData.NavigationEndpoint.BrowseEndpoint.BrowseID,
		Thumbnails: thumbnails,
	}, nil
}
