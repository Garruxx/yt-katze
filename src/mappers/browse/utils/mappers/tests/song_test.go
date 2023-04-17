package tests

import (
	"katze/src/mappers/browse/utils/mappers"
	"katze/src/models/external"
	"katze/src/utils"
	"testing"
)

func TestSong(t *testing.T) {

	// Test case 1 valid data
	songData := external.MusicPlaylistShelfContinuationContent{}
	err := utils.GetStructFromJson(
		"./data/json/song/valid.json",
		&songData,
	)
	if err != nil {
		t.Errorf("test 1 failed error while parsing json: %v", err)
	}
	// Test case 1.1 - check if the data has an error
	result, err := mappers.Song(songData)
	if err != nil {
		t.Errorf("test 1.1 failed error while mapping: %v", err)
	}
	// Test case 1.2 - check if the result has a title "titleTest"
	if result.Title != "titleTest" {
		t.Errorf("test 1.2 failed expected titleTest, got %s", result.Title)
	}
	// Test case 1.3 - check if the result has a valid ID "videoIDTest"
	if result.ID != "videoIDTest" {
		t.Errorf("test 1.3 failed expected videoIDTest, got %s", result.ID)
	}
	//  Test case 1.4 - check if the result has no explicit
	if *result.Explicit != false {
		t.Errorf("test 1.4 failed expected false, got %v", result.Explicit)
	}
	// Test case 1.5 - check if the result has a 1 artist
	if len(result.Artists) != 1 {
		t.Errorf(
			"test 1.5 failed expected 1 artist, got %v",
			len(result.Artists),
		)
	}
	// Test case 1.6 - check if the result has a valid artist "Michael Jackson"
	if name := result.Artists[0].Name; name != "Michael Jackson" {
		t.Errorf("test 1.6 failed expected Michael Jackson, got %s", name)
	}

	// Test case 2 invalid data empty
	songData = external.MusicPlaylistShelfContinuationContent{}
	result, err = mappers.Song(songData)
	if err == nil {
		t.Errorf("test 2 failed expected error, got nil")
	}

	// Test case 3 invalid data no 3 or 2 flexColumns just 1
	songData = external.MusicPlaylistShelfContinuationContent{}
	err = utils.GetStructFromJson(
		"./data/json/song/invalid_no_flex_columns.json",
		&songData,
	)
	if err != nil {
		t.Errorf("test 3 failed error while parsing json: %v", err)
	}
	// Test case 3.1 - check if the data has an error
	_, err = mappers.Song(songData)
	if err == nil {
		t.Errorf("test 3.1 failed expected error, got nil")
	}

	// Test case 4 valid explicit
	songData = external.MusicPlaylistShelfContinuationContent{}
	err = utils.GetStructFromJson(
		"./data/json/song/valid_explicit.json",
		&songData,
	)
	if err != nil {
		t.Errorf("test 4 failed error while parsing json: %v", err)
	}
	// Test case 4.1 - check if the data has an error
	result, err = mappers.Song(songData)
	if err != nil {
		t.Errorf("test 4.1 failed error while mapping: %v", err)
	}
	// Test case 4.2 - check if the result has a valid explicit
	if *result.Explicit != true {
		t.Errorf("test 4.2 failed expected true, got %v", result.Explicit)
	}

	// Test case 5 valid data with track number
	songData = external.MusicPlaylistShelfContinuationContent{}
	err = utils.GetStructFromJson(
		"./data/json/song/valid_with_track_number.json",
		&songData,
	)
	if err != nil {
		t.Errorf("test 5 failed error while parsing json: %v", err)
	}
	// Test case 5.1 - check if the data has an error
	result, err = mappers.Song(songData)
	if err != nil {
		t.Errorf("test 5.1 failed error while mapping: %v", err)
	}
	// Test case 5.2 - check if the result has a valid track number
	if *result.TrackNumber != 1 {
		t.Errorf("test 5.2 failed expected 1, got %v", result.TrackNumber)
	}
	// Test case 5.3 - check if the result has a title "titleTest"
	if result.Title != "titleTest" {
		t.Errorf("test 5.3 failed expected titleTest, got %s", result.Title)
	}
	// Test case 5.4 - check if the result has a valid ID "videoIDTest"
	if result.ID != "videoIDTest" {
		t.Errorf("test 5.4 failed expected videoIDTest, got %s", result.ID)
	}
}
