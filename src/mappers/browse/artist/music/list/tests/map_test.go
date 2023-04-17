package tests

import (
	"katze/src/mappers/browse/artist/music/list"
	"katze/src/models/external"
	"katze/src/utils"
	"testing"
)

func TestMap(t *testing.T) {

	// Test case 1 valid data
	artistMusicData := external.ArtistTracklist{}
	err := utils.GetStructFromJson(
		"./data/json/valid.json",
		&artistMusicData,
	)
	if err != nil {
		t.Errorf("test 1 failed, error: %v", err)
	}
	// Test case 1.1 check if the data has an error
	result, err := list.Map(artistMusicData)
	if err != nil {
		t.Errorf("test 1.1 failed, error: %v", err)
	}
	// Test case 1.2 check has a continuationID "ContinuationIDtest"
	if result.ContinuationID != "ContinuationIDtest" {
		t.Errorf(
			"test 1.2 failed, expected a continuationID, got %v",
			result.ContinuationID,
		)
	}
	// Test case 1.3 check has a songs length 22
	if len(result.Songs) != 22 {
		t.Errorf(
			"test 1.3 failed, expected a songs length 22, got %v",
			len(result.Songs),
		)
	}
	// Test case 1.4 check has a visitorID "VisitorIDtest"
	if result.VisitorID != "VisitorIDtest" {
		t.Errorf(
			"test 1.4 failed, expected a visitorID, got %v",
			result.VisitorID,
		)
	}

	// Test case 2 valid data no continuations 
	artistMusicData = external.ArtistTracklist{}
	err = utils.GetStructFromJson(
		"./data/json/valid_no_continuations.json",
		&artistMusicData,
	)
	if err != nil {
		t.Errorf("test 2 failed, error: %v", err)
	}
	// Test case 2.1 check if the data has an error
	result, err = list.Map(artistMusicData)
	if err != nil {
		t.Errorf("test 2.1 failed, error: %v", err)
	}
	// Test case 2.2 check has a continuationID ""
	if result.ContinuationID != "" {
		t.Errorf(
			"test 2.2 failed, expected a continuationID, got %v",
			result.ContinuationID,
		)
	}

	// Test case 3 invalid error 
	artistMusicData = external.ArtistTracklist{}
	err = utils.GetStructFromJson(
		"./data/json/invalid_error.json",
		&artistMusicData,
	)
	if err != nil {
		t.Errorf("test 3 failed, error: %v", err)
	}
	// Test case 3.1 check if the data has an error
	_, err = list.Map(artistMusicData)
	if err == nil {
		t.Errorf("test 3.1 failed, error: %v", err)
	}

	
}
