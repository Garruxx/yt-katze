package formatters

import (
	"fmt"
	bestUtils "katze/src/mappers/search/general/result/best/utils"
	"katze/src/mappers/search/general/result/best/utils/subtitle"
	"katze/src/mappers/utils"
	"katze/src/mappers/utils/columns"
	"katze/src/models/external"
	"katze/src/models/music"
	"strings"
)

func Album(bestMatch external.MusicCardShelfRenderer) (
	music.BestMatch, error,
) {

	bestMatchType, err := bestUtils.ShelfTypeValidator(bestMatch)
	if err != nil {
		return music.BestMatch{}, err
	}
	titleData := bestMatch.Title.Runs[0]

	if !strings.Contains(bestMatchType, "MUSIC_PAGE_TYPE_ALBUM") {
		err := fmt.Errorf("the best result is not of the album type")
		return music.BestMatch{}, err
	}

	flexcolumns, err := subtitle.ToFlexColumns(bestMatch.Subtitle)
	if err != nil {
		return music.BestMatch{}, err
	}

	artists, err := columns.GetArtists(flexcolumns)
	if err != nil {
		return music.BestMatch{}, err
	}

	thumbnails := utils.GetThumbnail(bestMatch.Thumbnail)

	return music.BestMatch{
		Type:       "album",
		Title:      titleData.Text,
		ID:         titleData.NavigationEndpoint.BrowseEndpoint.BrowseID,
		AlbumType:  flexcolumns[0].Text,
		Artists:    artists,
		Thumbnails: thumbnails,
	}, nil
}
