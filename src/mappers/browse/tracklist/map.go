package tracklist

import (
	"fmt"
	"katze/src/mappers/browse/utils/mappers"
	"katze/src/mappers/browse/utils/simplifier"
	"katze/src/models/external"
	"katze/src/models/shelves"
	"katze/src/utils"
)

func Map(albumData external.Tracklist) (
	shelves.Album, error,
) {
	if albumData.Error != nil {
		return shelves.Album{}, fmt.Errorf("error: %s", albumData.Error.Message)
	}
	if albumData.ResponseContext == nil {
		return shelves.Album{}, fmt.Errorf("response context is nil")
	}

	header, err := simplifier.MusicDetailHeaderRenderer(
		albumData.Header.MusicDetailHeaderRenderer,
	)
	if err != nil {
		return shelves.Album{}, err
	}

	tabs := albumData.Contents.SingleColumnBrowseResultsRenderer.Tabs
	if len(tabs) == 0 {
		return shelves.Album{}, fmt.Errorf("error: album tabs is empty")
	}
	tabsContents := tabs[0].TabRenderer.Content.SectionListRenderer.Contents
	if len(tabsContents) == 0 {
		return shelves.Album{}, fmt.Errorf("error: album tab contents is empty")
	}

	playlistShelf := tabsContents[0].MusicShelfRenderer.Contents
	songs, err := utils.ArrayMap(playlistShelf, mappers.Song)
	if err != nil {
		return shelves.Album{}, err
	}

	return shelves.Album{
		Title:      header.Title,
		Artists:    header.Artists,
		Thumbnails: header.Thumbnails,
		Tracks:     songs,
		Duration:   header.Duration,
		Year:       header.Year,
		VisitorID:  albumData.ResponseContext.VisitorData,
	}, nil
}
