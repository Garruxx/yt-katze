package tests

import (
	"katze/src/mappers/browse/utils/mappers"
	"katze/src/models/external"
	"katze/src/utils"
	"testing"
)

func TestMischievousSong(t *testing.T) {

	// Test case 1 valid data
	songData := external.MischievousContent{}
	err := utils.GetStructFromJson(
		"./data/json/mischievous_song/valid.json",
		&songData,
	)
	if err != nil {
		t.Errorf("test 1 failed error while parsing json: %v", err)
	}
	// Test case 1.1 - check if the data has an error
	result, err := mappers.MischievousSong(songData)
	if err != nil {
		t.Errorf("test 1.1 failed error while mapping: %v", err)
	}
	// Test case 1.2 - check if the result has a title
	if result.Title != "titleTest" {
		t.Errorf("test 1.2 failed expected titleTest, got %s", result.Title)
	}
	// Test case 1.3 - check if the result has a valid ID
	if result.ID != "videoIDTest" {
		t.Errorf("test 1.3 failed expected idTest, got %s", result.ID)
	}
	// Test case 1.4 - check if the result has a 2 thumbnails
	if len(result.Thumbnails) != 2 {
		t.Errorf(
			"test 1.4 failed, expected 2 thumbnails, got %v",
			len(result.Thumbnails),
		)
	}
	// Test case 1.5 - check if the result has a one artist
	if len(result.Artists) != 1 {
		t.Errorf(
			"test 1.5 failed expected 1 artist, got %v",
			len(result.Artists),
		)
	}
	// Test case 1.6 - check if the result has a valid artist Sous-Sol
	if name := result.Artists[0].Name; name != "Sous-Sol" {
		t.Errorf("test 1.6 failed expected Sous-Sol, got %s", name)
	}
	// Test case 1.7 - check if the result has a valid album title
	if album := *result.AlbumTitle; album != "albumTest" {
		t.Errorf("test 1.7 failed expected albumTest, got %s", album)
	}
	// Test case 1.8 - check if the result has a valid album ID
	if albumID := *result.AlbumID; albumID != "albumIDTest" {
		t.Errorf("test 1.8 failed expected albumIDTest, got %s", albumID)
	}
	// Test case 1.9 - check if the result has explicit
	if *result.Explicit != false {
		t.Errorf("test 1.9 failed expected explicit, got false")
	}

	// Test case 2 invalid data two flex columns
	songData = external.MischievousContent{}
	err = utils.GetStructFromJson(
		"./data/json/mischievous_song/invalid_two_flex_columns.json",
		&songData,
	)
	if err != nil {
		t.Errorf("test 2 failed error while parsing json: %v", err)
	}
	// Test case 2.1 - check if the data has an error
	_, err = mappers.MischievousSong(songData)
	if err == nil {
		t.Errorf("test 2.1 failed expected error, got nil")
	}

	// Test case 3 valid no album
	songData = external.MischievousContent{}
	err = utils.GetStructFromJson(
		"./data/json/mischievous_song/valid_no_album.json",
		&songData,
	)
	if err != nil {
		t.Errorf("test 3 failed error while parsing json: %v", err)
	}
	// Test case 3.1 - check if the data has an error
	result, err = mappers.MischievousSong(songData)
	if err != nil {
		t.Errorf("test 3,1 failed error while mapping: %v", err)
	}
	// Test case 3.2 - check if the result has a valid album title
	if *result.AlbumTitle != "Unknown" {
		t.Errorf("test 3.2 failed expected nil, got %v", *result.AlbumTitle)
	}
	// Test case 3.3 - check if the result has a valid album ID
	if *result.AlbumID != "Unknown" {
		t.Errorf("test 3.3 failed expected nil, got %v", *result.AlbumID)
	}

	// Test case 4 valid no artist
	songData = external.MischievousContent{}
	err = utils.GetStructFromJson(
		"./data/json/mischievous_song/valid_no_artist.json",
		&songData,
	)
	if err != nil {
		t.Errorf("test 4 failed error while parsing json: %v", err)
	}
	// Test case 4.1 - check if the data has an error
	result, err = mappers.MischievousSong(songData)
	if err != nil {
		t.Errorf("test 4.1 failed error while mapping: %v", err)
	}
	// Test case 4.2 - check if the result has a valid artist
	if leng := len(result.Artists); leng != 0 {
		t.Errorf("test 4.2 failed expected 1 artist, got %v", leng)
	}

	// Test case 5 valid explicit
	songData = external.MischievousContent{}
	err = utils.GetStructFromJson(
		"./data/json/mischievous_song/valid_explicit.json",
		&songData,
	)
	if err != nil {
		t.Errorf("test 5 failed error while parsing json: %v", err)
	}
	// Test case 5.1 - check if the data has an error
	result, err = mappers.MischievousSong(songData)
	if err != nil {
		t.Errorf("test 5.1 failed error while mapping: %v", err)
	}
	// Test case 5.2 - check if the result has a valid artist
	if *result.Explicit != true {
		t.Errorf("test 5.2 failed expected true, got %v", *result.Explicit)
	}
}
