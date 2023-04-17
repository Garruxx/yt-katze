package tests

import (
	"katze/src/mappers/search/music/list"
	"katze/src/models/external"
	"katze/src/utils"
	"testing"
)

func TestMapper(t *testing.T) {

	// test case 1 - get the data
	var musicData external.MusicList
	err := utils.GetStructFromJson(
		"./data/json/music_data_valid.json", &musicData,
	)
	if err != nil {
		t.Errorf("test 1 failed, error: %v", err)
	}

	// test case 1.1 - check if the data has an error
	result, err := list.Map(musicData)
	if err != nil {
		t.Errorf("test 1.1 failed, error: %v", err)
	}
	// test case 1.2 - check if the result has a
	if leng := len(result.Songs); leng != 20 {
		t.Errorf("test 1.2 failed, expected 20 tracks, got %v", leng)
	}
	// test case 1.3 - check if the result has a continuationID
	if result.ContinuationID == "" {
		t.Errorf("test 1.3 failed, expected a continuationID")
	}
	// test case 1.4 - check if the result has a visitorID
	if result.VisitorID == "" {
		t.Errorf("test 1.4 failed, expected a visitorID")
	}

	// test case 2 - get the data from the file path valid but the data
	// has an error
	musicData = external.MusicList{}
	err = utils.GetStructFromJson(
		"./data/json/music_data_error_invalid.json", &musicData,
	)
	if err != nil {
		t.Errorf("test 2 failed, error: %v", err)
	}
	// test case 2.1 - check if the error exists
	result, err = list.Map(musicData)
	if err == nil {
		t.Errorf("test 2.1 failed, error: %v", err)
	}

	musicData = external.MusicList{}
	// test case 3 - get the data no results
	err = utils.GetStructFromJson(
		"./data/json/music_data_no_results_valid.json", &musicData,
	)
	if err != nil {
		t.Errorf("test 3 failed, error: %v", err)
	}
	// test case 3.1 - check if the error exists
	result, err = list.Map(musicData)
	if err != nil {
		t.Errorf("test 3.1 failed, error: %v", err)
	}
	// test case 3.2 - check if the result has a continuationID
	if result.ContinuationID != "" {
		t.Errorf("test 3.2 failed, expected a continuationID empty")
	}
	// test case 3.3 - check if the result has a visitorID
	if result.VisitorID != "" {
		t.Errorf("test 3.3 failed, expected a visitorID empty")
	}
	// test case 3.4 - check if the result has a tracks length 0
	if leng := len(result.Songs); leng != 0 {
		t.Errorf("test 3.4 failed, expected 0 tracks, got %v", leng)
	}

}
