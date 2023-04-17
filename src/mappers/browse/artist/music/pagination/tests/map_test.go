package tests

import (
	"katze/src/mappers/browse/artist/music/pagination"
	"katze/src/models/external"
	"katze/src/utils"
	"testing"
)

func TestMap(t *testing.T) {
	// Tests case 1 - valid data
	paginationData := external.ArtistTracklistPagination{}
	err := utils.GetStructFromJson(
		"./data/json/valid.json",
		&paginationData,
	)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	// test case 1.1 check if the data has an error
	result, err := pagination.Map(paginationData)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	// test case 1.2 check has a continuationID "ContinuationIDtest"
	if result.ContinuationID != "ContinuationIDtest" {
		t.Errorf(
			"expected a continuationID, got %v",
			result.ContinuationID,
		)
	}
	// test case 1.3 check has a visitorID "VisitorIDtest"
	if result.VisitorID != "visitorIDTest" {
		t.Errorf(
			"expected a visitorIDTest, got %v",
			result.VisitorID,
		)
	}
	// test case 1.4 check has a songs length 4
	if len(result.Songs) != 4 {
		t.Errorf(
			"expected a songs length 4, got %v",
			len(result.Songs),
		)
	}

	// Tests case 2 - valid data no continuations
	paginationData = external.ArtistTracklistPagination{}
	err = utils.GetStructFromJson(
		"./data/json/valid_no_continuations.json",
		&paginationData,
	)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	// test case 2.1 check if the data has an error
	result, err = pagination.Map(paginationData)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	// test case 2.2 check has a continuationID ""
	if result.ContinuationID != "" {
		t.Errorf(
			"expected a continuationID, got %v",
			result.ContinuationID,
		)
	}
	// test case 2.3 check has a songs length 4
	if len(result.Songs) != 4 {
		t.Errorf(
			"expected a songs length 4, got %v",
			len(result.Songs),
		)
	}

	// Tests case 3 - invalid data error 
	paginationData = external.ArtistTracklistPagination{}
	err = utils.GetStructFromJson(
		"./data/json/invalid_error.json",
		&paginationData,
	)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	// test case 3.1 check if the data has an error
	result, err = pagination.Map(paginationData)
	if err == nil {
		t.Errorf("expected an error, got nil")
	}
	
}
