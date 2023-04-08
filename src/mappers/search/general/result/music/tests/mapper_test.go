package tests

import (
	"katze/src/mappers/search/general/result/music"
	"katze/src/mappers/search/general/result/music/tests/data"
	"testing"
)

func TestMusic(t *testing.T) {

	// test case 1 - valid data
	musicShelf := data.MUSIC_SHELF_RENDERER_VALID_MUSIC
	result, err := music.Mapper(musicShelf)
	if err != nil {
		t.Errorf("test 1 failed error: %v", err)
	}
	// test case 1.1 - check if the result is not empty and has 2 songs
	if leng := len(result.Songs); leng != 2 {
		t.Errorf("test 1.1 failed, expected 2 song but got %d", leng)
	}
	// test case 1.2 - check if the result continuation query is test1
	if query := result.Continuation.Query; query != "test1" {
		t.Errorf("test 1.2 failed, expected test1 but got %s", query)
	}
	// test case 1.3 - check if the result continuation params is test1
	if params := *result.Continuation.Params; params != "test1" {
		t.Errorf("test 1.3 failed, expected test1 but got %s", params)
	}

	// test case 2 - empty musicShelfRenderer
	musicShelf = data.MUSIC_SHELF_RENDERER_INVALID_EMPTY
	_, err = music.Mapper(musicShelf)
	if err == nil {
		t.Error("test 2 failed expected error, got: nil")
	}

	// test case 3 - invalid no songs
	musicShelf = data.MUSIC_SHELF_RENDERER_INVALID_NO_SONGS
	_, err = music.Mapper(musicShelf)
	if err == nil {
		t.Error("test 3 failed expected error, got: nil")
	}

	// test case 4 - invalid no contents
	musicShelf = data.MUSIC_SHELF_RENDERER_INVALID_NO_CONTENTS
	_, err = music.Mapper(musicShelf)
	if err == nil {
		t.Error("test 4 failed expected error, got: nil")
	}

	// test case 5 - invalid no bottom endpoint
	musicShelf = data.MUSIC_SHELF_RENDERER_INVALID_NO_BOTTOM_ENDPOINT
	_, err = music.Mapper(musicShelf)
	if err == nil {
		t.Error("test 5 failed expected error, got: nil")
	}

	// test case 6 - invalid title
	musicShelf = data.MUSIC_SHELF_RENDERER_INVALID_TITLE
	_, err = music.Mapper(musicShelf)
	if err == nil {
		t.Error("test 6 failed expected error, got: nil")
	}

}
