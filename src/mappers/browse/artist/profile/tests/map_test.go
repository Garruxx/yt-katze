package tests

import (
	"katze/src/mappers/browse/artist"
	"katze/src/models/external"
	"katze/src/utils"
	"testing"
)

func TestMap(t *testing.T) {

	// Test case 1 - valid artist profile
	artistData := external.Artist{}
	err := utils.GetStructFromJson(
		"./data/json/valid_profile.json",
		&artistData,
	)
	if err != nil {
		t.Errorf("test 1 failed error while parsing json: %v", err)
	}
	// Test case 1.1 - check if the data has an error
	result, err := artist.Profile(artistData)
	if err != nil {
		t.Errorf("test 1.1 failed error while parsing json: %v", err)
	}
	// Test case 1.2 - check if the result has a name "Sous Sol"
	if name := result.Name; name != "Sous Sol" {
		t.Errorf("test 1.2 failed expected Sous Sol, got %s", name)
	}
	// Test case 1.3 - check if the result has a valid visitor ID
	if visitorID := result.VisitorID; visitorID != "visitorIDTest" {
		t.Errorf("test 1.3 failed expected visitorIDTest, got %s", visitorID)
	}
	// Test case 1.4 - check if the result has a valid albums length
	if albumsLength := len(result.AlbumsList.Albums); albumsLength != 4 {
		t.Errorf("test 1.4 failed expected 2 albums, got %v", albumsLength)
	}
	// Test case 1.5 - check if the result has a valid songs length
	if songsLength := len(result.MusicList.Songs); songsLength != 4 {
		t.Errorf("test 1.5 failed expected 2 songs, got %v", songsLength)
	}
	// Test case 1.6 - check if the result has a valid singles length
	if singlesLength := len(result.SinglesList.Singles); singlesLength != 4 {
		t.Errorf("test 1.6 failed expected 2 singles, got %v", singlesLength)
	}
	// Test case 1.7 - check if the result has a valid thumbnails length
	if thumbnailsLength := len(result.Thumbnails); thumbnailsLength != 4 {
		t.Errorf(
			"test 1.7 failed expected 4 thumbnails, got %v",
			thumbnailsLength,
		)
	}

	// Test case 2 - invalid artist profile error
	artistData = external.Artist{}
	err = utils.GetStructFromJson(
		"./data/json/invalid_error_profile.json",
		&artistData,
	)
	if err != nil {
		t.Errorf("test 2 failed error while parsing json: %v", err)
	}
	// Test case 2.1 - check if the data has an error
	_, err = artist.Profile(artistData)
	if err == nil {
		t.Errorf("test 2.1 failed expected error, got nil")
	}

	// Test case 3 - valid artist profile with no albums
	artistData = external.Artist{}
	err = utils.GetStructFromJson(
		"./data/json/valid_profile_no_albums.json",
		&artistData,
	)
	if err != nil {
		t.Errorf("test 3 failed error while parsing json: %v", err)
	}
	// Test case 3.1 - check if the data has an error
	result, err = artist.Profile(artistData)
	if err != nil {
		t.Errorf("test 3.1 failed error while parsing json: %v", err)
	}
	// Test case 3.2 - check if the result has a valid albums length
	if albumsLength := len(result.AlbumsList.Albums); albumsLength != 0 {
		t.Errorf("test 3.2 failed expected 0 albums, got %v", albumsLength)
	}
	// Test case 3.3 - check if the result has a valid songs length
	if songsLength := len(result.MusicList.Songs); songsLength != 4 {
		t.Errorf("test 3.3 failed expected 4 songs, got %v", songsLength)
	}
	// Test case 3.4 - check if the result has a valid singles length
	if singlesLength := len(result.SinglesList.Singles); singlesLength != 4 {
		t.Errorf("test 3.4 failed expected 4 singles, got %v", singlesLength)
	}

	// Test case 4 - valid artist profile with no songs
	artistData = external.Artist{}
	err = utils.GetStructFromJson(
		"./data/json/valid_profile_no_songs.json",
		&artistData,
	)
	if err != nil {
		t.Errorf("test 4 failed error while parsing json: %v", err)
	}
	// Test case 4.1 - check if the data has an error
	result, err = artist.Profile(artistData)
	if err != nil {
		t.Errorf("test 4.1 failed error while parsing json: %v", err)
	}
	// Test case 4.2 - check if the result has a valid albums length
	if albumsLength := len(result.AlbumsList.Albums); albumsLength != 4 {
		t.Errorf("test 4.2 failed expected 4 albums, got %v", albumsLength)
	}
	// Test case 4.3 - check if the result has a valid songs length
	if songsLength := len(result.MusicList.Songs); songsLength != 0 {
		t.Errorf("test 4.3 failed expected 0 songs, got %v", songsLength)
	}
}
