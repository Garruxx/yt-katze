package tests

import (
	"katze/src/mappers/search/artists/list"
	"katze/src/models/external"
	"katze/src/utils"
	"testing"
)

func TestMap(t *testing.T) {

	// test case 1 - get valid data struct
	var artistsList external.ArtistList
	err := utils.GetStructFromJson(
		"./data/json/artist_list_valid.json", &artistsList,
	)
	if err != nil {
		t.Errorf("test 1 failed, error: %v", err)
	}
	// test case 1.1 - check if the data has an error
	result, err := list.Map(artistsList)
	if err != nil {
		t.Errorf("test 1.1 failed, error: %v", err)
	}
	// test case 1.2 - check if the result has a 22 artists
	if leng := len(result.Artists); leng != 22 {
		t.Errorf("test 1.2 failed, expected 22 artist, got %v", leng)
	}
	// test case 1.3 - check if the result has a continuationID
	if result.ContinuationID != "continuationIDTest" {
		t.Errorf(
			"test 1.3 failed, expected continuationIDTest, got %s",
			result.ContinuationID,
		)
	}
	// test case 1.4 - check if the result has a visitorID
	if result.VisitorID != "visitorIDTest" {
		t.Errorf(
			"test 1.4 failed, expected visitorIDTest, got %s",
			result.VisitorID,
		)
	}

	// test case 2 - get the data error
	artistsList = external.ArtistList{}
	err = utils.GetStructFromJson(
		"./data/json/artist_list_error.json", &artistsList,
	)
	if err != nil {
		t.Errorf("test 2 failed, error: %v", err)
	}
	// test case 2.1 - check if the error exists
	result, err = list.Map(artistsList)
	if err == nil {
		t.Errorf("test 2.1 failed, error: %v", err)
	}

	// test case 3 - get the data no results
	artistsList = external.ArtistList{}
	err = utils.GetStructFromJson(
		"./data/json/artist_list_no_results.json", &artistsList,
	)
	if err != nil {
		t.Errorf("test 3 failed, error: %v", err)
	}
	// test case 3.1 - check if the has an error
	result, err = list.Map(artistsList)
	if err != nil {
		t.Errorf("test 3.1 failed, error: %v", err)
	}
	// test case 3.2 - check if the result has a 0 artists
	if leng := len(result.Artists); leng != 0 {
		t.Errorf("test 3.2 failed, expected 0 artist, got %v", leng)
	}

}
