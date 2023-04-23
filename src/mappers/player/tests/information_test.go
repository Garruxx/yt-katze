package tests

import (
	"katze/src/mappers/player"
	"katze/src/models/external"
	"katze/src/utils"
	"testing"
)

func TestInformation(t *testing.T) {

	// Test case 1: valid input values
	infoData := external.PlayerInformation{}
	err := utils.GetStructFromJson(
		"./data/json/information_valid.json", &infoData,
	)
	if err != nil {
		t.Fatalf("Test case 1 failed: %v", err)
	}

	// Test case 1.1 valid data
	result, err := player.Information(infoData)
	if err != nil {
		t.Fatalf("Test case 1.1 failed: %v", err)
	}
	// Test case 1.2 the related browseID should be the same
	if result.Related.BrowseID != "browseIdRelatedTest" {
		t.Errorf("Test case 1.2 failed: expected browseIdRelatedTest")
	}
	// Test case 1.3 the lyrics browseID should be the same
	if result.Lyrics.BrowseID != "browseIdLyricsTest" {
		t.Errorf("Test case 1.3 failed: expected browseIdLyricsTest")
	}

	// Test case 2: invalid error
	infoData = external.PlayerInformation{}
	err = utils.GetStructFromJson("./data/json/error.json", &infoData)
	if err != nil {
		t.Fatalf("Test case 2 failed: %v", err)
	}

	// Test case 3 invalid data
	infoData = external.PlayerInformation{}
	err = utils.GetStructFromJson(
		"./data/json/information_invalid.json", &infoData,
	)
	if err != nil {
		t.Fatalf("Test case 3 failed: %v", err)
	}
	result, err = player.Information(infoData)
	if err != nil {
		t.Errorf("Test case 3.1 failed: error: %v", err)
	}
	// Test case 3.1 the related browseID should be empty
	if result.Related.BrowseID != "" {
		t.Errorf("Test case 3.2 failed: expected empty string")
	}
	// Test case 3.2 the lyrics browseID should be empty
	if result.Lyrics.BrowseID != "" {
		t.Errorf("Test case 3.3 failed: expected empty string")
	}

}
