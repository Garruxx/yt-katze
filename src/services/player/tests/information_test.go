package tests

import (
	"katze/src/services/player"
	"testing"
)

// Tracklist just requieres an videoID
// the visitorID is optional
func TestTracklist(t *testing.T) {

	// Test case 1 - valid videoID and visitorID
	videoID := "H9NJenpBV2I"
	visitorID := "CgswTi1jc3JLMnVuVSir06GhBg%3D%3D"
	result, err := player.Information(videoID, &visitorID)
	if err != nil {
		t.Errorf("Test case 1 failed: %v", err)
	}

	tabs := result.Contents.SingleColumnMusicWatchNextResultsRenderer.
		TabbedRenderer.WatchNextTabbedResultsRenderer.Tabs
	// Test case 1.1 - len of tabs should be 3
	if len(tabs) < 3 {
		t.Fatalf("Test case 1 failed: expected  2 or 3 tabs")
	}

	// Test case 1.2 - third tab should be browseEndpoint
	if tabs[2].TabRenderer.Endpoint.BrowseEndpoint == nil {
		t.Fatalf("Test case 1.2 failed: expected browseEndpoint")
	}
	// Test case 1.3 - third tab should be a endpoint
	if tabs[2].TabRenderer.Endpoint.BrowseEndpoint.BrowseID == "" {
		t.Errorf("Test case 1.3 failed: expected endpoint")
	}

	// Test case 2 - invalid videoID
	videoID = "Tirilil"
	visitorID = "CgswTi1jc3JLMnVuVSir06GhBg%3D%3D"
	result, err = player.Information(videoID, &visitorID)
	if err != nil {
		t.Fatalf("Test case 2 failed error: %v", err)
	}
	tabs = result.Contents.SingleColumnMusicWatchNextResultsRenderer.
		TabbedRenderer.WatchNextTabbedResultsRenderer.Tabs
	// Test case 2.1 - len of tabs should be 3
	if len(tabs) < 3 {
		t.Fatalf("Test case 2.1 failed: expected 3 tabs")
	}
	// Test case 2.2 - second tab should not be a browse endpoint
	if tabs[1].TabRenderer.Endpoint.BrowseEndpoint != nil {
		t.Errorf("Test case 2.2 failed: expected no browseEndpoint")
	}

	// Test case 3 - invalid visitorID
	visitorID = "Tirilil"
	videoID = "H9NJenpBV2I"
	result, err = player.Information(videoID, &visitorID)
	if err != nil {
		t.Fatalf("Test case 3 failed: %v", err)
	}
	tabs = result.Contents.SingleColumnMusicWatchNextResultsRenderer.
		TabbedRenderer.WatchNextTabbedResultsRenderer.Tabs

	if len(tabs) < 3 {
		t.Fatalf("Test case 3 failed: expected non-empty list of tabs")
	}
	if tabs[2].TabRenderer.Endpoint.BrowseEndpoint == nil {
		t.Errorf("Test case 3 failed: expected browseEndpoint")
	}
}
