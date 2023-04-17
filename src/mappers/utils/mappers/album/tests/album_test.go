package tests

import (
	"katze/src/mappers/utils/mappers/album"
	"katze/src/mappers/utils/mappers/album/tests/data"
	"testing"
)

func TestAlbum(t *testing.T) {

	// Test case 1: valid album
	itemRenderer := data.ITEM_ALBUM_VALID
	result, err := album.Mapper(itemRenderer)
	if err != nil {
		t.Errorf("test 1 failed, error: %v", err)
	}
	// Test case 1.1 check album name
	if result.Title != "AlbumTitleTest" {
		t.Errorf(
			"test 1.1 failed, expected Album name %s got: %v",
			"AlbumTitleTest", result.Title,
		)
	}
	// Test case 1.2 check album id
	if result.ID != "testID" {
		t.Errorf("test 1.2 failed, expected AlbumIdTest, got: %v", result.ID)
	}

	// Test case 2: invalid itemRenderer empty
	itemRenderer = data.ITEM_ALBUM_INVALID_EMPTY
	_, err = album.Mapper(itemRenderer)
	if err == nil {
		t.Errorf("test 2 failed, expected error, got: %v", err)
	}

	// Test case 3: invalid pageType
	itemRenderer = data.ITEM_ALBUM_INVALID_PAGE_TYPE
	_, err = album.Mapper(itemRenderer)
	if err == nil {
		t.Errorf("test 3 failed, expected error, got: %v", err)
	}

	// Test case 4: invalid no flexColumns length
	itemRenderer = data.ITEM_ALBUM_INVALID_FLEX_COLUMN_LENGTH
	_, err = album.Mapper(itemRenderer)
	if err == nil {
		t.Errorf("test 4 failed, expected error, got: %v", err)
	}

	// Test case 5: invalid no id
	itemRenderer = data.ITEM_ALBUM_INVALID_NO_ID
	_, err = album.Mapper(itemRenderer)
	if err == nil {
		t.Errorf("test 5 failed, expected error, got: %v", err)
	}
}
