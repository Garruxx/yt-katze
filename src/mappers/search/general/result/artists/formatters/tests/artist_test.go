package tests

import (
	"katze/src/mappers/search/general/result/artists/formatters"
	"katze/src/mappers/search/general/result/artists/formatters/tests/data"
	"testing"
)

func TestArtist(t *testing.T) {

	// Test case 1: valid artist
	itemRenderer := data.ITEM_ARTIST_VALID
	artist, err := formatters.Artist(itemRenderer)
	if err != nil {
		t.Errorf("test 1 failed, error: %v", err)
	}
	// Test case 1.1 check artist name
	if artist.Name != "ArtistNameTest" {
		t.Errorf("test 1.1 failed, expected Artist 1, got: %v", artist.Name)
	}
	// Test case 1.2 check artist id
	if artist.ID != "ArtistIdTest" {
		t.Errorf("test 1.2 failed, expected ArtistIdTest, got: %v", artist.ID)
	}

	// Test case 2: invalid itemRenderer empty
	itemRenderer = data.ITEM_ARTIST_INVALID_EMPTY
	_, err = formatters.Artist(itemRenderer)
	if err == nil {
		t.Errorf("test 2 failed, expected error, got: %v", err)
	}

	// Test case 3: invalid pageType 
	itemRenderer = data.ITEM_ARTIST_INVALID_PAGE_TYPE
	_, err = formatters.Artist(itemRenderer)
	if err == nil {
		t.Errorf("test 3 failed, expected error, got: %v", err)
	}

	// Test case 4: invalid no flexColumns length
	itemRenderer = data.ITEM_ARTIST_INVALID_FLEX_COLUMN_LENGTH
	_, err = formatters.Artist(itemRenderer)
	if err == nil {
		t.Errorf("test 4 failed, expected error, got: %v", err)
	}

	// Test case 5: invalid no id
	itemRenderer = data.ITEM_ARTIST_INVALID_NO_ID
	_, err = formatters.Artist(itemRenderer)
	if err == nil {
		t.Errorf("test 5 failed, expected error, got: %v", err)
	}
}
