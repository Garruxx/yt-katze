package tests

import (
	"katze/src/mappers/utils/mappers/artist"
	"katze/src/mappers/utils/mappers/artist/tests/data"
	"testing"
)

func TestArtist(t *testing.T) {

	// Test case 1: valid artist
	itemRenderer := data.ITEM_ARTIST_VALID
	result, err := artist.Mapper(itemRenderer)
	if err != nil {
		t.Errorf("test 1 failed, error: %v", err)
	}
	// Test case 1.1 check artist name
	if result.Name != "ArtistNameTest" {
		t.Errorf("test 1.1 failed, expected Artist 1, got: %v", result.Name)
	}
	// Test case 1.2 check artist id
	if result.ID != "testID" {
		t.Errorf("test 1.2 failed, expected ArtistIdTest, got: %v", result.ID)
	}

	// Test case 2: invalid itemRenderer empty
	itemRenderer = data.ITEM_ARTIST_INVALID_EMPTY
	_, err = artist.Mapper(itemRenderer)
	if err == nil {
		t.Errorf("test 2 failed, expected error, got: %v", err)
	}

	// Test case 3: invalid pageType
	itemRenderer = data.ITEM_ARTIST_INVALID_PAGE_TYPE
	_, err = artist.Mapper(itemRenderer)
	if err == nil {
		t.Errorf("test 3 failed, expected error, got: %v", err)
	}

	// Test case 4: invalid no flexColumns length
	itemRenderer = data.ITEM_ARTIST_INVALID_FLEX_COLUMN_LENGTH
	_, err = artist.Mapper(itemRenderer)
	if err == nil {
		t.Errorf("test 4 failed, expected error, got: %v", err)
	}

	// Test case 5: invalid no id
	itemRenderer = data.ITEM_ARTIST_INVALID_NO_ID
	_, err = artist.Mapper(itemRenderer)
	if err == nil {
		t.Errorf("test 5 failed, expected error, got: %v", err)
	}
}
