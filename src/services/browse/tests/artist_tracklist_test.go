package tests

import (
	"katze/src/services/browse"
	"testing"
)

// ArtistTracklist just requieres an browseID ans Params
// the visitorID is optional
func TestArtistTracklist(t *testing.T) {

	// Test case 1 - valid browseID, params and visitorID

	browseID := "VLOLAK5uy_nD0V5fFISV9uMmvPTFxKITDbmLUqAMeRY"
	params := "ggMCCAI%3D"
	visitorID := "CgswTi1jc3JLMnVuVSir06GhBg%3D%3D"
	result, err := browse.ArtistTracklist(browseID, params, &visitorID)
	if err != nil {
		t.Errorf("Test case 1 failed: %v", err)
	}
	if len(result.Contents.SingleColumnBrowseResultsRenderer.Tabs) == 0 {
		t.Errorf("Test case 1 failed: expected non-empty list of tabs")
	}

	// Test case 2 - invalid browseID
	browseID = "Tirilil"
	visitorID = "CgswTi1jc3JLMnVuVSir06GhBg%3D%3D"
	params = "ggMCCAI%3D"
	result, err = browse.ArtistTracklist(browseID, params, &visitorID)
	if err == nil {
		t.Errorf("Test case 2 failed expected error but got nil: %v", err)
	}

	// Test case 3 - invalid visitorID
	visitorID = "Tirilil"
	browseID = "UCoIOOL7QKuBhQHVKL8y7BEQ"
	params = "ggMCCAI%3D"
	result, err = browse.ArtistTracklist(browseID, params, &visitorID)
	if err != nil {
		t.Errorf("Test case 3 failed: %v", err)
	}
	if len(result.Contents.SingleColumnBrowseResultsRenderer.Tabs) == 0 {
		t.Errorf("Test case 3 failed: expected non-empty list of tabs")
	}

	// Test case 4 - invalid params
	params = "Tirilil"
	browseID = "UCoIOOL7QKuBhQHVKL8y7BEQ"
	visitorID = "CgswTi1jc3JLMnVuVSir06GhBg%3D%3D"
	result, err = browse.ArtistTracklist(browseID, params, &visitorID)
	if err != nil {
		t.Errorf("Test case 4 failed expected error but got nil: %v", err)
	}

	// Test case 5 - nil visitorID
	var nilVisitorID *string
	browseID = "UCoIOOL7QKuBhQHVKL8y7BEQ"
	params = "ggMCCAI%3D"
	result, err = browse.ArtistTracklist(browseID, params, nilVisitorID)
	if err != nil {
		t.Errorf("Test case 5 failed: %v", err)
	}
}
