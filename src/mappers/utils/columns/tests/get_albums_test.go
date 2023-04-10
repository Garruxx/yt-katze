package tests

import (
	"katze/src/mappers/utils/columns"
	"katze/src/mappers/utils/columns/tests/data"
	"testing"
)

func TestGetAlbums(t *testing.T) {

	// Test case 1: Test with valid data
	FlexColumn := data.GET_ALBUM_DATATEST_VALID
	albums, err := columns.GetAlbums(FlexColumn)
	if err != nil {
		t.Errorf("Test 1 failed, error: %s", err)
	}
	// Test case 1.1: Test with valid data, length of albums is equal to 2
	if leng := len(albums); leng != 2 {
		t.Errorf("Test 1.1 failed, expected: %d, got: %d", 2, leng)
	}

	// Test case 1.2: Test with valid data,
	// first album title is equal to expected "testText"
	if title := albums[0].Title; title != "testText" {
		t.Errorf(
			"Test 1 failed, expected: %s, got: %s",
			"testText", title,
		)
	}
	// Test case 1.3: Test with valid data,
	// first album ID is equal to expected "testID"
	if id := albums[0].ID; id != "testID" {
		t.Errorf("Test 1.3 failed, expected: %s, got: %s", "testID", id)
	}
	// Test case 1.4: Test with valid data,
	// second album title is equal to expected "testText1"
	if title := albums[1].Title; title != "testText1" {
		t.Errorf(
			"Test 1.4 failed, expected: %s, got: %s",
			"testText1", title,
		)
	}
	// Test case 1.5: Test with valid data,
	// second album ID is equal to expected "testID1"
	if id := albums[1].ID; id != "testID1" {
		t.Errorf("Test 1 failed, expected: %s, got: %s", "testID1", id)
	}

	// Test case 2: Test with empty data
	FlexColumn = data.GET_ALBUM_DATATEST_EMPTY
	albums, err = columns.GetAlbums(FlexColumn)
	if err != nil {
		t.Errorf("Test 2 failed, error: %s", err)
	}
	// Test case 2.1: Test with empty data, length of albums is equal to 1
	if leng := len(albums); leng != 1 {
		t.Errorf(
			"Test 2 failed, expected: %d, got: %d",
			0, len(albums),
		)
	}

	// Test case 3: Test with invalid data, just text
	FlexColumn = data.GET_ALBUM_DATATEST_INVALID_JUST_TEXT
	_, err = columns.GetAlbums(FlexColumn)
	if err == nil {
		t.Error("Test 3 failed, expected error but got, nil")
	}

	// Test case 4: Test with invalid data, browseID is empty
	FlexColumn = data.GET_ALBUM_DATATEST_INVALID_BROWSE_ID_EMPTY
	_, err = columns.GetAlbums(FlexColumn)
	if err == nil {
		t.Error("Test 4 failed, expected error but got, nil")
	}

	// Test case 5: Test with invalid data, watchID is not empty
	FlexColumn = data.GET_ALBUM_DATATEST_INVALID_WATCH_ID_NOT_EMPTY
	_, err = columns.GetAlbums(FlexColumn)
	if err == nil {
		t.Error("Test 5 failed, expected error but got, nil")
	}

	// Test case 6: Test with no albums
	FlexColumn = data.GET_ALBUM_DATATEST_NO_ALBUMS
	albums, err = columns.GetAlbums(FlexColumn)
	if err != nil {
		t.Errorf("Test 6 failed, error: %s", err)
	}
	// Test case 6.1: Test with no albums, length of albums is equal to 1
	if leng := len(albums); leng != 1 {
		t.Errorf("Test 6.1 failed, expected: %d, got: %d", 0, leng)
	}
	// Test case 6.2: Test with no albums, title Unknown
	if title := albums[0].Title; title != "Unknown" {
		t.Errorf(
			"Test 6.2 failed, expected: %s, got: %s",
			"Unknown", title,
		)
	}
	// Test case 6.3: Test with no albums, ID Unknown
	if id := albums[0].ID; id != "Unknown" {
		t.Errorf("Test 6.3 failed, expected: %s, got: %s", "Unknown", id)
	}

}
