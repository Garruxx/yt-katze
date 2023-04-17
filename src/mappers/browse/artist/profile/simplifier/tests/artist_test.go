package tests

import (
	"katze/src/mappers/browse/artist/profile/simplifier"
	"katze/src/models/external"
	"katze/src/utils"
	"testing"
)

func TestArtist(t *testing.T) {

	// Test case 1: Test with valid data
	artistData := external.Artist{}
	err := utils.GetStructFromJson(
		"./data/json/artist_valid.json", &artistData,
	)
	if err != nil {
		t.Errorf("test 1 failed, error: %v", err)
	}
	// Test case 1.1: Check if the data has an error
	result, err := simplifier.Artist(artistData)
	if err != nil {
		t.Errorf("test 1.1 failed, error: %v", err)
	}
	// Test case 1.2: Check if the result has a name
	if result.Name != "Sous Sol" {
		t.Errorf(
			"test 1.2 failed, expected name Artist Name, got %v",
			result.Name,
		)
	}
	// Test case 1.3: Check if the result has a carousels with 4 items
	if leng := len(result.Carousels); leng != 4 {
		t.Errorf("test 1.3 failed, expected 4 items, got %v", leng)
	}
	// Test case 1.4: Check if the result has a valid musicShelf
	if len(result.MusicShelfRenderer.Title.Runs) == 0 {
		t.Errorf("test 1.4 failed, expected musicShelf, got nil")
	}
	// Test case 1.5: Check if the result has a valid visitorID
	if result.VisitorID != "visitorIDTest" {
		t.Errorf(
			"test 1.5 failed, expected visitorID visitorIDTest, got %v",
			result.VisitorID,
		)
	}

	// Test case 2: Test with no contents
	artistData = external.Artist{}
	err = utils.GetStructFromJson(
		"./data/json/artist_invalid.json", &artistData,
	)
	if err != nil {
		t.Errorf("test 2 failed, error: %v", err)
	}
	// Test case 2.1: Check if the data has an error
	_, err = simplifier.Artist(artistData)
	if err == nil {
		t.Errorf("test 2.1 failed, expected an error, got: %v", err)
	}
}
