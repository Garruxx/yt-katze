package tests

import (
	"katze/src/mappers/search/general/result/albums"
	"katze/src/mappers/search/general/result/albums/test/data"
	"testing"
)

func TestAlbums(t *testing.T) {

	// test case 1 valid data
	shelfRenderer := data.MUSIC_SHELF_RENDERER_VALID_ALBUMS
	result, err := albums.Mapper(shelfRenderer)
	if err != nil {
		t.Error("test 1 failed, error:", err)
	}
	// test case 1.1 check if the albums list is not empty
	if len(result.Albums) == 0 {
		t.Error("test 1.1 failed, albums list is empty")
	}
	// test case 1.2 check if the continuation is valid
	if result.Continuation.Query == "" || result.Continuation.Params == nil {
		t.Error("test 1.2 failed, continuation is not valid")
	}

	// test case 2 empty data
	shelfRenderer = data.MUSIC_SHELF_RENDERER_INVALID_EMPTY
	_, err = albums.Mapper(shelfRenderer)
	if err == nil {
		t.Error("test 2 failed, error is nil")
	}

	// test case 3 invalid no albums
	shelfRenderer = data.MUSIC_SHELF_RENDERER_INVALID_NO_ALBUMS
	_, err = albums.Mapper(shelfRenderer)
	if err == nil {
		t.Error("test 3 failed, error is nil")
	}

	// test case 4 invalid empty data
	shelfRenderer = data.MUSIC_SHELF_RENDERER_INVALID_NO_CONTENTS
	_, err = albums.Mapper(shelfRenderer)
	if err == nil {
		t.Error("test 4 failed, error is nil")
	}

}
