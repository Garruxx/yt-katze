package tests

import (
	"katze/src/services/browse"
	"testing"
)

// ArtistMusicListPagination just requieres an continuationID
// the visitorID is optional
func TestArtistMusicListPagination(t *testing.T) {

	// Test case 1: valid continuationID and visitorID
	continuationID :=
		"4qmFsgJFEitWTE9MQUs1dXlfbkQwVjVmRklTVjl1TW12UFRGeEtJVERibUxVcUFNZVJZGhZlZ1pRVkRwRFIyZVNBUU1JdWdRJTNE"
	visitorID := "CgswTi1jc3JLMnVuVSir06GhBg%3D%3D"
	result, err := browse.ArtistMusicListPagination(continuationID, &visitorID)
	if err != nil {
		t.Errorf("Test case 1 failed: %v", err)
	}
	if len(
		result.ContinuationContents.MusicPlaylistShelfContinuation.Contents,
	) == 0 {

		t.Errorf("Test case 1 failed: expected non-empty list of tracks")
	}

	// Test case 2: invalid continuationID
	continuationID = "Tirilil"
	visitorID = "CgswTi1jc3JLMnVuVSir06GhBg%3D%3D"
	result, err = browse.ArtistMusicListPagination(continuationID, &visitorID)
	if err == nil {
		t.Errorf("Test case 2 failed expected error but got nil: %v", err)
	}

	// Test case 3: invalid visitorID
	continuationID = "4qmFsgJFEitWTE9MQUs1dXlfbkQwVjVmRklTVjl1TW12UFRGeEtJVERibUxVcUFNZVJZGhZlZ1pRVkRwRFIyZVNBUU1JdWdRJTNE"
	visitorID = "Tirilil"
	result, err = browse.ArtistMusicListPagination(continuationID, &visitorID)
	if err != nil {
		t.Errorf("Test case 3 failed: %v", err)
	}
	if len(
		result.ContinuationContents.MusicPlaylistShelfContinuation.Contents,
	) == 0 {
		t.Errorf("Test case 3 failed: expected non-empty list of tracks")
	}

	// Test case 4: nil visitorID
	var nilVisitorID *string
	continuationID = "4qmFsgJFEitWTE9MQUs1dXlfbkQwVjVmRklTVjl1TW12UFRGeEtJVERibUxVcUFNZVJZGhZlZ1pRVkRwRFIyZVNBUU1JdWdRJTNE"
	result, err = browse.ArtistMusicListPagination(continuationID, nilVisitorID)
	if err != nil {
		t.Errorf("Test case 4 failed: %v", err)
	}
	if len(
		result.ContinuationContents.MusicPlaylistShelfContinuation.Contents,
	) == 0 {
		t.Errorf("Test case 4 failed: expected non-empty list of tracks")
	}
}
