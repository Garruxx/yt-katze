package tests

import (
	"katze/src/mappers/search/general/result/albums/formatters"
	"katze/src/mappers/search/general/result/albums/formatters/tests/data"
	"testing"
)

func TestAlbum(t *testing.T) {

	// Test case 1 Valid album single
	item := data.ITEM_ALBUM_VALID_SINGLE
	result, err := formatters.Album(item)
	if err != nil {
		t.Errorf("failed test 1 error: %v", err)
	}

	// Test case 1.1 Valid album title
	if result.Title != "TitleTest" {
		t.Errorf(
			"failed test 1.1 error: invalid title expected: %v, got: %v", "TitleTest", result.Title,
		)
	}
	// Test case 1.2 Valid album artists
	if len(*result.Artists) != 2 {
		t.Errorf(
			"failed test 1.2 error: invalid artists length expected: %v, got: %v", "2", len(*result.Artists),
		)
	}
	// Test case 1.3 Valid album thumbnails
	thumbnails := *result.Thumbnails
	if thumbnails[0].URL != "example.com" {
		t.Errorf(
			"failed test 1.3 error: invalid url expected: %v, got: %v", "example.com", thumbnails[0].URL,
		)
	}
	if len(thumbnails) != 2 {
		t.Errorf(
			`failed test 1.3 error: invalid thumbnails length expected:
			%v, got: %v`,
			"2", len(thumbnails),
		)
	}
	// Test case 1.4 Valid album id
	if result.ID != "AlbumTestID" {
		t.Errorf(
			"failed test 1.4 error: invalid id expected: %v, got: %v",
			"AlbumTestID", result.ID,
		)
	}
	// Test case 1.5 Valid album year
	if *result.Year != 2022 {
		t.Errorf(
			"failed test 1.5 error: invalid year expected: %v, got: %v", "2020",
			result.Year,
		)
	}
	// Test case 1.6 Valid album explicit
	if *result.Explicit != true {
		t.Errorf(
			"failed test 1.6 error: invalid explicit expected: %v, got: %v",
			"true", result.Explicit,
		)
	}

	// Test case 2 valid album ep
	item = data.ITEM_ALBUM_VALID_EP
	result, err = formatters.Album(item)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	// Test case 2.1 Valid album ep
	if *result.EP != true {
		t.Errorf(
			"error: faild test 2.1 invalid ep expected: %v, got: %v", "true", result.EP,
		)
	}

	// Test case 3 valid album
	item = data.ITEM_ALBUM_VALID
	result, err = formatters.Album(item)
	if err != nil {
		t.Errorf("failed test 3 error: %v", err)
	}

	// Test case 3.1 Valid album ep false
	if *result.EP != false {
		t.Errorf(
			"failed test 3.1 error: invalid ep expected: %v, got: %v",
			"nil", result.EP,
		)
	}
	// Test case 3.2 Valid album single false
	if *result.Single != false {
		t.Errorf(
			"failed test 3.2 error: invalid single expected: %v, got: %v",
			"nil", result.Single,
		)
	}

	// Test case 4 valid no explicit
	item = data.ITEM_ALBUM_VALID_NO_EXPLICIT
	result, err = formatters.Album(item)
	if err != nil {
		t.Errorf("failed test 4 error: %v", err)
	}
	// Test case 4.1 Valid album explicit false
	if *result.Explicit != false {
		t.Errorf(
			"failed test 4.1 error: invalid explicit expected: %v, got: %v",
			"false", result.Explicit,
		)
	}

	// Test case 5 invalid empty
	item = data.ITEM_ALBUM_INVALID_EMPTY
	_, err = formatters.Album(item)
	if err == nil {
		t.Errorf("failed test 5 error: expected error got: %v", "nil")
	}

	// Test case 6 invalid no id
	item = data.ITEM_ALBUM_INVALID_NO_ID
	_, err = formatters.Album(item)
	if err == nil {
		t.Errorf("failed test 6 error: expected error got: %v", "nil")
	}

	// Test case 7 invalid no flex columns
	item = data.ITEM_ALBUM_INVALID_NO_FLEX_COLUMNS
	_, err = formatters.Album(item)
	if err == nil {
		t.Errorf("failed test 7 error: expected error got: %v", "nil")
	}

	// Test case 8 invalid year
	item = data.ITEM_ALBUM_INVALID_YEAR
	_, err = formatters.Album(item)
	if err == nil {
		t.Errorf("failed test 8 error: expected error got: %v", "nil")
	}

	// Test case 9 invalid no page type
	item = data.ITEM_ALBUM_INVALID_NO_PAGE_TYPE
	_, err = formatters.Album(item)
	if err == nil {
		t.Errorf("failed test 9 error: expected error got: %v", "nil")
	}
}
