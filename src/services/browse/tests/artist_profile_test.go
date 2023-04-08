package tests

import (
	"katze/src/services/browse"
	"testing"
)

func TestArtistProfile(t *testing.T) {

	// Test case 1 - valid artistID and visitorID
	browseID := "UCoIOOL7QKuBhQHVKL8y7BEQ"
	visitorID := "CgswTi1jc3JLMnVuVSir06GhBg%3D%3D"
	result, err := browse.ArtistProfile(browseID, &visitorID)
	if err != nil {
		t.Errorf("Test case 1 failed: %v", err)
	}
	if len(result.Contents.SingleColumnBrowseResultsRenderer.Tabs) == 0 {
		t.Errorf("Test case 1 failed: expected non-empty list of tabs")
	}

	// Test case 2 - invalid artistID
	browseID = "Tirilil"
	visitorID = "CgswTi1jc3JLMnVuVSir06GhBg%3D%3D"
	result, err = browse.ArtistProfile(browseID, &visitorID)
	if err == nil {
		t.Errorf("Test case 2 failed expected error but got nil: %v", err)
	}

	// Test case 3 - invalid visitorID
	visitorID = "Tirilil"
	browseID = "UCoIOOL7QKuBhQHVKL8y7BEQ"
	result, err = browse.ArtistProfile(browseID, &visitorID)
	if err != nil {
		t.Errorf("Test case 3 failed: %v", err)
	}
	if len(result.Contents.SingleColumnBrowseResultsRenderer.Tabs) == 0 {
		t.Errorf("Test case 3 failed: expected non-empty list of tabs")
	}

	// Test case 4 - nil visitorID
	var nilVisitorID *string
	browseID = "UCoIOOL7QKuBhQHVKL8y7BEQ"
	result, err = browse.ArtistProfile(browseID, nilVisitorID)
	if err != nil {
		t.Errorf("Test case 4 failed: %v", err)
	}
	if len(result.Contents.SingleColumnBrowseResultsRenderer.Tabs) == 0 {
		t.Errorf("Test case 4 failed: expected non-empty list of tabs")
	}
}
