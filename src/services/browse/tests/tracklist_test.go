package tests

import (
	"katze/src/services/browse"
	"testing"
)


//Tracklist just requieres an tracklistID
//the visitorID is optional
func TestTracklist(t *testing.T) {

	// Test case 1 - valid TracklistID and visitorID
	TracklistID := "UC7_Cx5UF7PvY7mwAm0Iq7rw"
	visitorID := "CgswTi1jc3JLMnVuVSir06GhBg%3D%3D"
	result, err := browse.Tracklist(TracklistID, &visitorID)
	if err != nil {
		t.Errorf("Test case 1 failed: %v", err)
	}
	if len(result.Contents.SingleColumnBrowseResultsRenderer.Tabs) == 0 {
		t.Errorf("Test case 1 failed: expected non-empty list of tabs")
	}

	// Test case 2 - invalid TracklistID
	TracklistID = "Tirilil"
	result, err = browse.Tracklist(TracklistID, &visitorID)
	if err == nil {
		t.Errorf("Test case 2 failed expected error but got nil: %v", err)
	}

	// Test case 3 - invalid visitorID
	visitorID = "Tirilil"
	TracklistID = "UC7_Cx5UF7PvY7mwAm0Iq7rw"
	result, err = browse.Tracklist(TracklistID, &visitorID)
	if err != nil {
		t.Errorf("Test case 3 failed: %v", err)
	}
	if len(result.Contents.SingleColumnBrowseResultsRenderer.Tabs) == 0 {
		t.Errorf("Test case 3 failed: expected non-empty list of tabs")
	}
}
