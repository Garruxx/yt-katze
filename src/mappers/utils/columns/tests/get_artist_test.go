package tests

import (
	"katze/src/mappers/utils/columns"
	"katze/src/mappers/utils/columns/tests/data"
	"katze/src/models"
	"testing"
)

func TestGetArtits(t *testing.T) {

	// Test case 1: flex column is valid
	columnsData := data.GET_ARTIST_FLEX_COLUMS_WITH_TWO_ARTISTS
	artists, err := columns.GetArtists(columnsData)
	if err != nil {
		t.Errorf("Test 1 failed, error: %s", err.Error())
	}
	// Test case 1.1: length of artists is equal to 2
	if leng := len(artists); leng != 2 {
		t.Errorf("Test 1.1 failed, expected 2 artist, got %d", leng)
	}
	// Test case 1.2: first artist name is equal to expected "artistName"
	if name := artists[0].Name; name != "artistName" {
		t.Errorf(
			"Test 1 failed, expected: %s, got: %s",
			"artistName", name,
		)
	}
	// Test case 1.3: first artist ID is equal to expected "testID"
	if id := artists[0].ID; id != "testID" {
		t.Errorf(
			"Test 1.3 failed, expected: %s, got: %s",
			"testID", id,
		)
	}

	// Test case 2: flex column is empty
	columnsData = []models.FlexColumn{}
	artists, err = columns.GetArtists(columnsData)
	if err != nil {
		t.Errorf("Test 2 failed, error: %s", err.Error())
	}
	// Test case 2.1: length of artists is equal to 0
	if leng := len(artists); leng != 0 {
		t.Errorf("Test 2 failed, expected %d artist, got %d", 0, leng)
	}

	// Test case 3: flex column contains invalid data
	columnsData = data.GET_ARTIST_FLEX_COLUMS_WITH_INVALID_DATA
	artists, err = columns.GetArtists(columnsData)
	if err == nil {
		t.Errorf("Test 3 expected error, got nil")
	}
	// Test case 3.1: length of artists is equal to 0
	if leng := len(artists); leng != 0 {
		t.Errorf("Test 3.1 failed, expected %d artist, got %d", 0, leng)
	}
}
