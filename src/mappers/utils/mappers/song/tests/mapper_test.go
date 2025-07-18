package tests

import (
	"katze/src/mappers/utils/mappers/song"
	"katze/src/mappers/utils/mappers/song/tests/data"
	"testing"
)

func TestSon(t *testing.T) {

	// Test case 1 song valid
	musicItem := data.ITEM_SONG_VALID 
	result, err := song.Mapper(musicItem)
	if err != nil {
		t.Error("test 1 failed, error:", err)
	}
	// Test case 1.1 check if the song id is valid "SongTestID"
	if result.ID != "testID" {
		t.Errorf(
			"test 1.1 failed, song id is not valid expected: %s got: %s", "testID", result.ID,
		)
	}
	// Test case 1.2 check if the song title is valid "TitleTest"
	if result.Title != "TitleTest" {
		t.Errorf(
			"test 1.2 failed, song title is not valid expected: %s got: %s", "TitleTest", result.Title,
		)
	}
	// Test case 1.3 check if the song duration is valid "22:22"
	if duration := *result.Duration; duration != "22:22" {
		t.Errorf(
			"test 1.3 failed, song duration is not valid expected: %s got: %s", "22:22", duration,
		)
	}
	// Test case 1.4 check if the song is explicit
	if explicit := *result.Explicit; explicit != true {
		t.Errorf(
			"test 1.4 failed, song explicit is not valid expected: %t got: %t", true, explicit,
		)
	}
	// Test case 1.5 check if the song artist is valid "artistNameTest"
	if artist := result.Artists[0].Name; artist != "artistNameTest" {
		t.Errorf(
			"test 1.5 failed, song artist is not valid expected: %s got: %s", "artistNameTest", artist,
		)
	}

	// Test case 2 song valid no explicit
	musicItem = data.ITEM_SONG_VALID_NO_EXPLICIT
	result, err = song.Mapper(musicItem)
	if err != nil {
		t.Error("test 2 failed, error:", err)
	}
	// Test case 2.1 check if the song is not explicit
	if explicit := *result.Explicit; explicit != false {
		t.Errorf(
			"test 2.1 failed, song explicit is not valid expected: %t got: %t", false, explicit,
		)
	}

	// Test case 3 song invalid empty
	musicItem = data.ITEM_SONG_INVALID_EMPTY
	_, err = song.Mapper(musicItem)
	if err == nil {
		t.Error("test 3 failed, error is nil")
	}

	// Test case 4 song invalid no id
	musicItem = data.ITEM_SONG_INVALID_NO_VIDEO_ID
	_, err = song.Mapper(musicItem)
	if err == nil {
		t.Error("test 4 failed, error is nil")
	}

	// Test case 5 song invalid no flexColumns
	musicItem = data.ITEM_SONG_INVALID_NO_FLEX_COLUMNS
	_, err = song.Mapper(musicItem)
	if err == nil {
		t.Error("test 5 failed, error is nil")
	}

}
