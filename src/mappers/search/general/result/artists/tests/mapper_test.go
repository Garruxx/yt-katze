package tests

import (
	"katze/src/mappers/search/general/result/artists"
	"katze/src/mappers/search/general/result/artists/tests/data"
	"testing"
)

func TestMapper(t *testing.T) {

	// Test case 1 valid data
	shelfRenderer := data.MUSIC_SHELF_RENDERER_VALID_ARTISTS
	result, err := artists.Mapper(shelfRenderer)
	if err != nil {
		t.Error("test 1 failed, error:", err)
	}
	// test case 1.1 check if the artists list is not empty
	if len(result.Artists) == 0 {
		t.Error("test 1.1 failed, artists list is empty")
	}
	// test case 1.2 check if the continuation qury is valid
	if query := result.Continuation.Query; query != "test1" {
		t.Errorf(
			"test 1.2 failed, continuation query expected: test1, got: %s",
			query,
		)
	}
	// test case 1.3 check if the continuation params is valid
	if params := *result.Continuation.Params; params != "test1" {
		t.Errorf(
			"test 1.3 failed, continuation params expected: test1, got: %s",
			params,
		)
	}
	// test case 1.4 check if the id is correct
	if id := result.Artists[0].ID; id != "testID" {
		t.Errorf("test 1.4 failed, id expected: testID, got: %s", id)
	}

	// Test case 2 empty data
	shelfRenderer = data.MUSIC_SHELF_RENDERER_INVALID_EMPTY
	_, err = artists.Mapper(shelfRenderer)
	if err == nil {
		t.Error("test 2 failed, expected error, got: nil")
	}

	// Test case 3 invalid no artists
	shelfRenderer = data.MUSIC_SHELF_RENDERER_INVALID_NO_ARTISTS
	_, err = artists.Mapper(shelfRenderer)
	if err == nil {
		t.Error("test 3 failed, error is nil")
	}

	// Test case 4 invalid no contents
	shelfRenderer = data.MUSIC_SHELF_RENDERER_INVALID_NO_CONTENTS
	_, err = artists.Mapper(shelfRenderer)
	if err == nil {
		t.Error("test 4 failed, error is nil")
	}

}
