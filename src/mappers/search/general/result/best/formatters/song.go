package formatters

import (
	"fmt"
	bestUtils "katze/src/mappers/search/general/result/best/utils"
	"katze/src/mappers/utils"
	"katze/src/mappers/utils/columns"
	"katze/src/mappers/utils/subtitle"
	"katze/src/models/external"
	"katze/src/models/music"
	"regexp"
	"strings"
)

/*
	this function does not use shelfTypeValidator, because it requires watchEndpoint and not browseEndpoint
*/

// Song convert the data into a bestMatch if possible
func Song(bestMatch external.MusicCardShelfRenderer) (
	music.BestMatch, error,
) {
	var regexpTimeFormats = regexp.MustCompile(`^\d{1,3}(?::\d{1,2}){1,2}$`)
	bestMatchType, err := bestUtils.ShelfTypeValidator(bestMatch)
	if err != nil {
		return music.BestMatch{}, err
	}

	if !strings.Contains(bestMatchType, "MUSIC_VIDEO") {
		err := fmt.Errorf("the best result is not of the music type")
		return music.BestMatch{}, err
	}

	flexColumns, err := subtitle.ToFlexColumns(bestMatch.Subtitle)
	if err != nil {
		return music.BestMatch{}, err
	}

	albums, err := columns.GetAlbums(flexColumns)
	if err != nil {
		return music.BestMatch{}, err
	}

	var duration string
	lastText := flexColumns[len(flexColumns)-1].Text
	if regexpTimeFormats.MatchString(lastText) {
		duration = lastText
	}

	artists, err := columns.GetArtists(flexColumns)
	if err != nil {
		return music.BestMatch{}, err
	}

	var albumTitle string
	var albumID string
	if len(albums) > 0 {
		albumTitle = albums[0].Title
		albumID = albums[0].ID
	}

	if len(bestMatch.Title.Runs) == 0 {
		err := fmt.Errorf("Could not find runs in title")
		return music.BestMatch{}, err
	}

	titleData := bestMatch.Title.Runs[0]
	explicit := bestMatch.SubtitleBadges != nil
	thumbnails := utils.GetThumbnail(bestMatch.Thumbnail)
	title := titleData.Text
	ID := titleData.NavigationEndpoint.WatchEndpoint.VideoID

	return music.BestMatch{
		Type:       "song",
		Title:      title,
		ID:         ID,
		Artists:    artists,
		AlbumTitle: albumTitle,
		AlbumID:    albumID,
		Duration:   duration,
		Explicit:   explicit,
		Thumbnails: thumbnails,
	}, nil
}
