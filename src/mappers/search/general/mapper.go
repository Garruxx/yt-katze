package general

import (
	"fmt"
	"katze/src/mappers/search/general/result/albums"
	"katze/src/mappers/search/general/result/artists"
	"katze/src/mappers/search/general/result/best"
	"katze/src/mappers/search/general/result/music"
	"katze/src/mappers/search/utils"
	"katze/src/models/external"
	"katze/src/models/lists"
	musicModels "katze/src/models/music"
	"katze/src/models/shelves"
)

func Mapper(searchResult external.General) (
	shelves.General, error,
) {

	if searchResult.Contents == nil {
		err := fmt.Errorf("error: searchResult.Contents is empty")
		return shelves.General{}, err
	}

	var albumList lists.Albums
	var tracksList lists.Music
	var artistsList lists.Artists
	var bestMatch musicModels.BestMatch

	tabContents :=
		searchResult.Contents.TabbedSearchResultsRenderer.
			Tabs[0].TabRenderer.Content.SectionListRenderer.Contents

	/*
	* if the error is not nil and the error is "could not find shelf"
	* or "cloud not find best result" then return the list empty
	 */

	//Get the best result
	bestResultTab, err := utils.GetBestResult(tabContents)
	if err == nil {
		bestMatch, err = best.Mapper(bestResultTab)
		if err != nil {
			return shelves.General{}, err
		}
	} else if err != nil && err.Error() != "cloud not find best result" {
		return shelves.General{}, err
	}

	// Get the albums
	albumsTab, err := utils.FindCardShelf(tabContents, "Albums")
	if err == nil {
		albumList, err = albums.Mapper(albumsTab)
		if err != nil {
			return shelves.General{}, err
		}
	} else if err != nil && err.Error() != "could not find shelf" {
		return shelves.General{}, err
	}

	// Get the tracks
	tracksTab, err := utils.FindCardShelf(tabContents, "Songs")
	if err == nil {
		tracksList, err = music.Mapper(tracksTab)
		if err != nil {
			return shelves.General{}, err
		}
	} else if err != nil && err.Error() != "could not find shelf" {
		return shelves.General{}, err
	}

	// Get the artists
	artistsTab, err := utils.FindCardShelf(tabContents, "Artists")
	if err == nil {
		artistsList, err = artists.Mapper(artistsTab)
		if err != nil {
			return shelves.General{}, err
		}
	} else if err != nil && err.Error() != "could not find shelf" {
		return shelves.General{}, err
	}

	// Get the visitor ID, if it exists
	visitorID := searchResult.ResponseContext.VisitorData

	return shelves.General{
		BestMatch: bestMatch,
		Albums:    albumList,
		Tracks:    tracksList,
		Artists:   artistsList,
		VisitorID: visitorID,
	}, nil
}
