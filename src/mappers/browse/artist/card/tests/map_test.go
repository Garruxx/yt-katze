package test

import (
	"katze/src/mappers/browse/artist/card"
	"katze/src/models/external"
	"katze/src/utils"
	"testing"
)

func TestMap(t *testing.T) {

	// Test case 1 - valid albums artist data
	albumsData := external.ArtistAlbums{}
	err := utils.GetStructFromJson(
		"./data/json/valid_albums_artist.json",
		&albumsData,
	)
	if err != nil {
		t.Errorf("test 1 failed error while parsing json: %v", err)
	}
	// Test case 1.1 - check if the data has an error
	result, err := card.TwoRowItem(albumsData)
	if err != nil {
		t.Errorf("test 1.1 failed, error: %v", err)
	}
	// Test case 1.2 - check if the result has 34 items
	if len(result) != 34 {
		t.Errorf("test 1.2 failed, expected 34 items, got %v", len(result))
	}
	// Test case 1.3 - check if the last result has a title
	//"Farewell My Summer Love"
	if title := result[33].Title; title != "Farewell My Summer Love" {
		t.Errorf(
			"test 1.3 failed, expected title Farewell My Summer Love, got %v",
			title,
		)
	}
	// Test case 1.4 - check if the last result has year 1984
	if year := result[33].Year; year != 1984 {
		t.Errorf(
			"test 1.4 failed, expected year 1984, got %v",
			year,
		)
	}
	// Test case 1.5 - check if the first result has a title Thriller 40
	if title := result[0].Title; title != "Thriller 40" {
		t.Errorf(
			"test 1.5 failed, expected title Thriller 40, got %v",
			title,
		)
	}
	// Test case 1.6 - check if the first result has year 2022
	if year := result[0].Year; year != 2022 {
		t.Errorf(
			"test 1.6 failed, expected year 2022, got %v",
			year,
		)
	}

	// Test case 2 - invalid albums artist data bad visitor ID
	albumsData = external.ArtistAlbums{}
	err = utils.GetStructFromJson(
		"./data/json/invalid_visitor_id.json",
		&albumsData,
	)
	if err != nil {
		t.Errorf("test 2 failed error while parsing json: %v", err)
	}
	// Test case 2.1 - check if the data has an error
	_, err = card.TwoRowItem(albumsData)
	if err == nil {
		t.Errorf("test 2.1 failed, expected error, got nil")
	}
	// Test case 2.2 - check if the error is
	//"error possibly caused by the visitor ID in the request"
	if err.Error() != "error possibly caused by the visitor ID in the request" {
		t.Errorf(
			"test 2.2 failed, expected error error possibly caused by the visitor ID in the request, got %v",
			err.Error(),
		)
	}

	// Test case 3 - invalid albums artist data error
	albumsData = external.ArtistAlbums{}
	err = utils.GetStructFromJson(
		"./data/json/invalid_error.json",
		&albumsData,
	)
	if err != nil {
		t.Errorf("test 3 failed error while parsing json: %v", err)
	}
	// Test case 3.1 - check if the data has an error
	_, err = card.TwoRowItem(albumsData)
	if err == nil {
		t.Errorf("test 3.1 failed, expected error, got nil")
	}
}
