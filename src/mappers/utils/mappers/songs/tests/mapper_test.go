package tests

import (
	"katze/src/mappers/utils/mappers/songs"
	"katze/src/mappers/utils/mappers/songs/tests/data"
	"testing"
)

func TestMapper(t *testing.T) {

	// test case 1 - items valid two songs
	itemsSongs := data.ITEMS_SONGS_VALID_TWO_SONGS
	result, err := songs.Mapper(itemsSongs)
	if err != nil {
		t.Errorf("test 1 failed, error: %v", err)
	}
	// test case 1.1 - check if the result has two songs
	if len(result) != 2 {
		t.Errorf("test 1.1 failed, expected 2 songs, got %v", len(result))
	}
	// test case 1.2 - check if the fisrt song has the correct title
	if title := result[0].Title; title != "TitleTest" {
		t.Errorf("test 1.2 failed, expected TitleTest, got %v", title)
	}

	// test case 2 - items valid one song
	itemsSongs = data.ITEMS_SONGS_VALID_ONE_SONG
	result, err = songs.Mapper(itemsSongs)
	if err != nil {
		t.Errorf("test 2 failed, error: %v", err)
	}
	// test case 2.1 - check if the result has one song
	if leng := len(result); leng != 1 {
		t.Errorf("test 2.1 failed, expected 1 song, got %v", leng)
	}
	// test case 2.2 - check if the fisrt song has the correct title
	if title := result[0].Title; title != "TitleTest" {
		t.Errorf("test 2.2 failed, expected TitleTest, got %v", title)
	}

	// test case 3 - items empty
	itemsSongs = data.ITEMS_SONGS_EMPTY
	_, err = songs.Mapper(itemsSongs)
	if err != nil {
		t.Errorf("test 3 failed, error: %v", err)
	}
}
