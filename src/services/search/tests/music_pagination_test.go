package tests

import (
	"fmt"
	"katze/src/services/search"
	"testing"
)

func TestMusicPagination(t *testing.T) {

	// Test case 1: continuation ID and visitor ID are valid
	continuationID := "EqQDEg1BZ3VhIGNvbiBjaGlhGpIDRWdXS0FRSUlBVWdVYWdnUUF4QUVFQWtRQ29JQkMyTXRWRWRXVEMxeU5ubE5nZ0VMY21neUxYZExkMU5qYjJPQ0FRczBSMVZQT1VWS1JXMDFiNElCQ3pSd1J6QlRNbXBSTUdoWmdnRUxYeTFHZHpCZlFWVmlVR2VDQVF0QlNGOW1VakZaZHpKRU9JSUJDMUZLTjJJMlJYcFJkVTVyZ2dFTGJIbENTMHRZTUd0dVVtdUNBUXMwVTNGUVF6SjZORGRsYjRJQkMwVXRiVk5PV1VzMllYTlpnZ0VMWWtFNFMzcDBPRUl4V1ZHQ0FRc3RiakU0YmtKU2J6QnpTWUlCQzFCWlp6Wk1SRzlpUWpGQmdnRUxlbTFCT1ZWT2NIRkhaRm1DQVFzellYVkNhMkZIVlhOMGQ0SUJDMkpzUmxGVGJuRndWRzR3Z2dFTGQwZFBSM2hETVRjeE9HLUNBUXRNVUhCbk56UnpRbWRuTUlJQkMyVjZlbEZFWWpobVZYZE5nZ0VMWVRKb2VucDFiVjlCWlZFJTNEGPHq0C4%3D"

	visitorID := "CgswTi1jc3JLMnVuVSir06GhBg%3D%3D"
	result, err := search.MusicPagination(continuationID, visitorID)
	if err != nil {
		fmt.Print(err)
		t.Errorf("Test case 1 failed: %v", err)
	}
	if len(result.ContinuationContents.MusicShelfContinuation.Contents) == 0 {
		t.Errorf("Test case 1 failed: expected non-empty list of tracks")
	}

	// Test case 2: continuation ID is invalid
	continuationID = "invalid"
	result, err = search.MusicPagination(continuationID, visitorID)
	if err == nil {
		t.Errorf("Test case 2 failed: expected error, but got nil")
	}

	// Test case 3: visitor ID is nil
	continuationID = "EqQDEg1BZ3VhIGNvbiBjaGlhGpIDRWdXS0FRSUlBVWdVYWdnUUF4QUVFQWtRQ29JQkMyTXRWRWRXVEMxeU5ubE5nZ0VMY21neUxYZExkMU5qYjJPQ0FRczBSMVZQT1VWS1JXMDFiNElCQ3pSd1J6QlRNbXBSTUdoWmdnRUxYeTFHZHpCZlFWVmlVR2VDQVF0QlNGOW1VakZaZHpKRU9JSUJDMUZLTjJJMlJYcFJkVTVyZ2dFTGJIbENTMHRZTUd0dVVtdUNBUXMwVTNGUVF6SjZORGRsYjRJQkMwVXRiVk5PV1VzMllYTlpnZ0VMWWtFNFMzcDBPRUl4V1ZHQ0FRc3RiakU0YmtKU2J6QnpTWUlCQzFCWlp6Wk1SRzlpUWpGQmdnRUxlbTFCT1ZWT2NIRkhaRm1DQVFzellYVkNhMkZIVlhOMGQ0SUJDMkpzUmxGVGJuRndWRzR3Z2dFTGQwZFBSM2hETVRjeE9HLUNBUXRNVUhCbk56UnpRbWRuTUlJQkMyVjZlbEZFWWpobVZYZE5nZ0VMWVRKb2VucDFiVjlCWlZFJTNEGPHq0C4%3D"
	var nilVisitorID string
	result, err = search.MusicPagination(continuationID, nilVisitorID)
	if err != nil {
		t.Errorf("Test case 3 failed: %v", err)
	}
	if len(result.ContinuationContents.MusicShelfContinuation.Contents) == 0 {
		t.Errorf("Test case 3 failed: expected non-empty list of tracks")
	}
}
