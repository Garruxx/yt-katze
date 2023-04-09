package tests

import (
	"katze/src/mappers/search/general"
	"katze/src/mappers/search/general/tests/data"
	"testing"
)

func TestMap(t *testing.T) {

	// test case 1 - get the data from the file path valid
	generalData, err := data.Get("./data/json/general_data_valid.json")
	if err != nil {
		t.Errorf("test 1 failed, error: %v", err)
	}
	// test case 1.1 - check if the best match exists
	result, err := general.Mapper(generalData)
	if err != nil {
		t.Errorf("test 1.1 failed, error: %v", err)
	}
	// test case 1.2 - check if the best match exists
	if result.BestMatch.WatchID == "" {
		t.Errorf("test 1.2 failed, error: best match is empty")
	}
	// test case 1.3 - check if the music exists 3 songs
	if leng := len(result.Tracks.Songs); leng != 3 {
		t.Errorf("test 1.3 failed, expected 3 songs, got %v", leng)
	}
	// test case 1.4 - check if the music exists 3 albums
	if leng := len(result.Albums.Albums); leng != 3 {
		t.Errorf("test 1.4 failed, expected 3 albums, got %v", leng)
	}
	// test case 1.5 - check if the music exists 3 artists
	if leng := len(result.Artists.Artists); leng == 3 {
		t.Errorf("test 1.5 failed, expected 3 artists, got %v", leng)
	}

	// test case 2 - get the data from the file path valid but the data
	// contains an error
	generalData, err = data.Get("./data/json/general_data_error_invalid.json")
	if err != nil {
		t.Errorf("test 2 failed, error: %v", err)
	}
	// check if the error exists
	result, err = general.Mapper(generalData)
	if err == nil {
		t.Errorf("test 2.1 failed, error: %v", err)
	}

	// test case 3 - get the data from the file path valid, no artists
	generalData, err = data.Get(
		"./data/json/general_data_valid_no_artists.json",
	)
	if err != nil {
		t.Errorf("test 3 failed, error: %v", err)
	}
	// test case 3.1 check if the error exists
	result, err = general.Mapper(generalData)
	if err != nil {
		t.Errorf("test 3.1 failed, error: %v", err)
	}
	// test case 3.2 - check if the artist is empty
	if leng := len(result.Artists.Artists); leng != 0 {
		t.Errorf("test 3.2 failed, expected 0 artists, got %v", leng)
	}

	// test case 4 - get the data from the file path valid, no songs
	generalData, err = data.Get("./data/json/general_data_valid_no_songs.json")
	if err != nil {
		t.Errorf("test 4 failed, error: %v", err)
	}
	// test case 4.1 check if the error exists
	result, err = general.Mapper(generalData)
	if err != nil {
		t.Errorf("test 4.1 failed, error: %v", err)
	}
	// test case 4.2 - check if the songs is empty
	if leng := len(result.Tracks.Songs); leng != 0 {
		t.Errorf("test 4.2 failed, expected 0 songs, got %v", leng)
	}

	// test case 5 - get the data from the file path valid, no albums
	generalData, err = data.Get(
		"./data/json/general_data_valid_no_albums.json",
	)
	if err != nil {
		t.Errorf("test 5 failed, error: %v", err)
	}
	// test case 5.1 check if the error exists
	result, err = general.Mapper(generalData)
	if err != nil {
		t.Errorf("test 5.1 failed, error: %v", err)
	}
	// test case 5.2 - check if the albums is empty
	if leng := len(result.Albums.Albums); leng != 0 {
		t.Errorf("test 5.2 failed, expected 0 albums, got %v", leng)
	}

	// test case 6 - get the data from the file path valid, no best match
	generalData, err = data.Get(
		"./data/json/general_data_valid_no_best_match.json",
	)
	if err != nil {
		t.Errorf("test 6 failed, error: %v", err)
	}
	// test case 6.1 check if the error exists
	result, err = general.Mapper(generalData)
	if err != nil {
		t.Errorf("test 6.1 failed, error: %v", err)
	}
	// test case 6.2 - check if the best match is empty
	if result.BestMatch.WatchID != "" {
		t.Errorf("test 6.2 failed, expected empty best match, got %v",
			result.BestMatch.WatchID,
		)
	}
}
