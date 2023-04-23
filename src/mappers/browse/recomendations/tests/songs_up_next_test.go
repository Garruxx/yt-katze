package tests

import (
	"katze/src/mappers/browse/recomendations"
	"katze/src/models/external"
	"katze/src/utils"
	"testing"
)

func TestSongsUpNexts(t *testing.T) {
	// Test case 1: valid input values
	recomendationsData := external.MusicRecomendations{}
	err := utils.GetStructFromJson(
		"./data/json/valid.json", &recomendationsData,
	)
	if err != nil {
		t.Fatalf("Test case 1 failed: %v", err)
	}
	// Test case 1.1 should not be error
	result, err := recomendations.SongUpNext(recomendationsData)
	if err != nil {
		t.Fatalf("Test case 1.1 failed: %v", err)
	}
	// Test case 1.2 should have 20 elements
	if len(result) != 20 {
		t.Errorf("Test case 1.2 failed: expected 20 elements")
	}
	// Test case 1.3 the first element should be title "titleTest"
	if result[0].Title != "titleTest" {
		t.Errorf("Test case 1.3 failed: expected titleTest")
	}
	// Test case 1.4 the first element should be one artist
	if len(result[0].Artists) != 1 {
		t.Errorf("Test case 1.4 failed: expected 1 artist")
	}
	// Test case 1.5 the first element should be artist "artistTest"
	if result[0].Artists[0].Name != "artistTest" {
		t.Errorf("Test case 1.5 failed: expected artistTest")
	}
	// Test case 1.6 the first element should be one album
	if result[0].AlbumTitle == nil {
		t.Errorf("Test case 1.6 failed: expected 1 album")
	}
	// Test case 1.7 the first element should be album "albumTest"
	if *result[0].AlbumTitle != "albumTest" {
		t.Errorf("Test case 1.6 failed: expected albumTest")
	}

	// Test case 2: invalid error
	recomendationsData = external.MusicRecomendations{}
	err = utils.GetStructFromJson("./data/json/error.json", &recomendationsData)
	if err != nil {
		t.Fatalf("Test case 2 failed: %v", err)
	}
	// Test case 2.1 should be error
	result, err = recomendations.SongUpNext(recomendationsData)
	if err == nil {
		t.Errorf("Test case 2.1 failed: expected error")
	}

	// Test case 3: invalid data
	recomendationsData = external.MusicRecomendations{}
	err = utils.GetStructFromJson(
		"./data/json/invalid_no_response_context.json", &recomendationsData,
	)
	if err != nil {
		t.Fatalf("Test case 3 failed: %v", err)
	}
	// Test case 3.1 should be error
	result, err = recomendations.SongUpNext(recomendationsData)
	if err == nil {
		t.Fatalf("Test case 3.1 failed: expected error, got: %v", err)
	}

	// Test case 4: invalid data
	recomendationsData = external.MusicRecomendations{}
	err = utils.GetStructFromJson(
		"./data/json/invalid_no_contents.json", &recomendationsData,
	)
	if err != nil {
		t.Fatalf("Test case 4 failed: %v", err)
	}
	// Test case 4.1 should be error
	result, err = recomendations.SongUpNext(recomendationsData)
	if err == nil {
		t.Fatalf("Test case 4.1 failed: expected error, got: %v", err)
	}

}
