package tests

import (
	"katze/src/mappers/search/artists/pagination"
	"katze/src/models/external"
	"katze/src/utils"
	"testing"
)

func TestMap(t *testing.T) {

	// test case 1 - get valid artists pagination data
	var itemsArtists external.ArtistPagination
	err := utils.GetStructFromJson(
		"./data/json/pagination_data_valid.json", &itemsArtists,
	)
	if err != nil {
		t.Errorf("test 1.1 failed error getting data: %v", err)
	}
	// test case 1.2 - map valid artists pagination data
	result, err := pagination.Map(itemsArtists)
	if err != nil {
		t.Errorf("test 1.2 failed error mapping data: %v", err)
	}
	// test case 1.3 - validate length 22
	if leng := len(result.Artists); leng != 22 {
		t.Errorf("test 1.3 failed expected 22 got %v", leng)
	}
	// test case 1.4 - validate continuation id
	if id := result.ContinuationID; id != "continuationIDTest" {
		t.Errorf(
			"test 1.4 failed expected continuationIDTest got %v", id,
		)
	}
	// test case 1.5 - validate visitor id
	if visitorID := result.VisitorID; visitorID != "visitorIDTest" {
		t.Errorf(
			"test 1.5 failed expected visitorIDTest got %v", visitorID,
		)
	}

	// test case 2 - get invalid artists pagination data error
	itemsArtists = external.ArtistPagination{}
	err = utils.GetStructFromJson(
		"./data/json/pagination_data_error.json", &itemsArtists,
	)
	if err != nil {
		t.Errorf("test 2.1 failed error getting data: %v", err)
	}
	// test case 2.2 - map invalid artists pagination data error
	_, err = pagination.Map(itemsArtists)
	if err == nil {
		t.Errorf("test 2.2 failed expected error got nil")
	}

	// test case 3 - get invalid artists pagination data no artists
	itemsArtists = external.ArtistPagination{}
	err = utils.GetStructFromJson(
		"./data/json/pagination_data_no_results.json", &itemsArtists,
	)
	if err != nil {
		t.Errorf("test 3.1 failed error getting data: %v", err)
	}
	// test case 3.2 - map invalid artists pagination data no artists
	result, err = pagination.Map(itemsArtists)
	if err != nil {
		t.Errorf("test 3.2 failed error mapping data: %v", err)
	}
	// test case 3.3 - validate length 0
	if leng := len(result.Artists); leng != 0 {
		t.Errorf("test 3.3 failed expected 0 got %v", leng)
	}
}
