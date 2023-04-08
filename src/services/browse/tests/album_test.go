package tests

import (
	"katze/src/services/browse"
	"testing"
)


//Album just requieres an albumID
//the visitorID is optional
func TestAlbum(t *testing.T) {

	// Test case 1 - valid albumID and visitorID
	albumID := "UC7_Cx5UF7PvY7mwAm0Iq7rw"
	visitorID := "CgswTi1jc3JLMnVuVSir06GhBg%3D%3D"
	result, err := browse.Album(albumID, &visitorID)
	if err != nil {
		t.Errorf("Test case 1 failed: %v", err)
	}
	if len(result.Contents.SingleColumnBrowseResultsRenderer.Tabs) == 0 {
		t.Errorf("Test case 1 failed: expected non-empty list of tabs")
	}

	// Test case 2 - invalid albumID
	albumID = "Tirilil"
	result, err = browse.Album(albumID, &visitorID)
	if err == nil {
		t.Errorf("Test case 2 failed expected error but got nil: %v", err)
	}

	// Test case 3 - invalid visitorID
	visitorID = "Tirilil"
	albumID = "UC7_Cx5UF7PvY7mwAm0Iq7rw"
	result, err = browse.Album(albumID, &visitorID)
	if err != nil {
		t.Errorf("Test case 3 failed: %v", err)
	}
	if len(result.Contents.SingleColumnBrowseResultsRenderer.Tabs) == 0 {
		t.Errorf("Test case 3 failed: expected non-empty list of tabs")
	}
}
