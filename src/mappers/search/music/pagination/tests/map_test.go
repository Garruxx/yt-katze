package test

import (
	"katze/src/mappers/search/music/pagination"
	"katze/src/mappers/search/music/pagination/tests/data"
	"testing"
)

func TestMap(t *testing.T) {

	// test case 1 - get valid music pagination data
	itemsSongs, err := data.Get(
		"./data/json/pagination_data_valid.json",
	)
	if err != nil {
		t.Errorf("test 1.1 failed error getting data: %v", err)
	}
	// test case 1.2 - map valid music pagination data
	result, err := pagination.Map(itemsSongs)
	if err != nil {
		t.Errorf("test 1.2 failed error mapping data: %v", err)
	}
	// test case 1.3 - validate length 10
	if leng := len(result.Tracks); leng != 10 {
		t.Errorf("test 1.3 failed expected 2 got %v", leng)
	}
	// test case 1.4 - validate continuation id
	if id := result.ContinuationID; id != "testContinuationID" {
		t.Errorf(
			"test 1.4 failed expected testContinuationID got %v", id,
		)
	}
	// test case 1.5 - validate visitor id
	if visitorID := result.VisitorID; visitorID != "testVisitorData" {
		t.Errorf(
			"test 1.5 failed expected testVisitorData got %v", visitorID,
		)
	}

	// test case 2 - get invalid music pagination data error
	itemsSongs, err = data.Get(
		"./data/json/pagination_data_error_invalid.json",
	)
	if err != nil {
		t.Errorf("test 2.1 failed error getting data: %v", err)
	}
	// test case 2.2 - map invalid music pagination data error
	_, err = pagination.Map(itemsSongs)
	if err == nil {
		t.Errorf("test 2.2 failed expected error got nil")
	}

	// test case 3 - get invalid music pagination data no songs
	itemsSongs, err = data.Get(
		"./data/json/pagination_data_no_songs_valid.json",
	)
	if err != nil {
		t.Errorf("test 3.1 failed error getting data: %v", err)
	}
	// test case 3.2 - map invalid music pagination data no songs
	result, err = pagination.Map(itemsSongs)
	if err != nil {
		t.Errorf("test 3.2 failed error mapping data: %v", err)
	}
	// test case 3.3 - validate length 0
	if leng := len(result.Tracks); leng != 0 {
		t.Errorf("test 3.3 failed expected 0 got %v", leng)
	}
}
