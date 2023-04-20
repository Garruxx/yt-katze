package tests

import (
	"katze/src/graphql/resolvers/search"
	"katze/src/graphql/resolvers/tests/utils"
	"katze/src/models/lists"
	"testing"
)

func TestMusicsPagination(t *testing.T) {

	// Test case 1 - valid continuationID and visitorID
	continuationID := "EqQDEg1BZ3VhIGNvbiBjaGlhGpIDRWdXS0FRSUlBVWdVYWdnUUF4QUVFQWtRQ29JQkMyTXRWRWRXVEMxeU5ubE5nZ0VMY21neUxYZExkMU5qYjJPQ0FRczBSMVZQT1VWS1JXMDFiNElCQ3pSd1J6QlRNbXBSTUdoWmdnRUxYeTFHZHpCZlFWVmlVR2VDQVF0QlNGOW1VakZaZHpKRU9JSUJDMUZLTjJJMlJYcFJkVTVyZ2dFTGJIbENTMHRZTUd0dVVtdUNBUXMwVTNGUVF6SjZORGRsYjRJQkMwVXRiVk5PV1VzMllYTlpnZ0VMWWtFNFMzcDBPRUl4V1ZHQ0FRc3RiakU0YmtKU2J6QnpTWUlCQzFCWlp6Wk1SRzlpUWpGQmdnRUxlbTFCT1ZWT2NIRkhaRm1DQVFzellYVkNhMkZIVlhOMGQ0SUJDMkpzUmxGVGJuRndWRzR3Z2dFTGQwZFBSM2hETVRjeE9HLUNBUXRNVUhCbk56UnpRbWRuTUlJQkMyVjZlbEZFWWpobVZYZE5nZ0VMWVRKb2VucDFiVjlCWlZFJTNEGPHq0C4%3D"
	visitorID := "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D"
	result, err := utils.GqlResolver[lists.Music](
		map[string]any{
			"continuationId": continuationID,
			"visitorId":      visitorID,
		},
		search.MusicsPagination,
	)
	if err != nil {
		t.Errorf("Test case 1 failed: %v", err)
	}
	// Test case 1.1 verify that the result is not nil
	if len(result.Songs) == 0 {
		t.Errorf("Test case 1.1 failed: expected non-nil music list")
	}

	// Test case 2 - invalid continuationID
	continuationID = "Tirilil"
	visitorID = "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D"
	result, err = utils.GqlResolver[lists.Music](
		map[string]any{
			"continuationId": continuationID,
			"visitorId":      visitorID,
		},
		search.MusicsPagination,
	)
	if err == nil {
		t.Errorf("Test case 2 failed expected an error, got: %v", err)
	}
	// Test case 2.1 verify that the result is nil
	if len(result.Songs) != 0 {
		t.Errorf("Test case 2.1 failed: expected nil music list")
	}

	// Test case 3 - invalid visitorID
	continuationID = "EqQDEg1BZ3VhIGNvbiBjaGlhGpIDRWdXS0FRSUlBVWdVYWdnUUF4QUVFQWtRQ29JQkMyTXRWRWRXVEMxeU5ubE5nZ0VMY21neUxYZExkMU5qYjJPQ0FRczBSMVZQT1VWS1JXMDFiNElCQ3pSd1J6QlRNbXBSTUdoWmdnRUxYeTFHZHpCZlFWVmlVR2VDQVF0QlNGOW1VakZaZHpKRU9JSUJDMUZLTjJJMlJYcFJkVTVyZ2dFTGJIbENTMHRZTUd0dVVtdUNBUXMwVTNGUVF6SjZORGRsYjRJQkMwVXRiVk5PV1VzMllYTlpnZ0VMWWtFNFMzcDBPRUl4V1ZHQ0FRc3RiakU0YmtKU2J6QnpTWUlCQzFCWlp6Wk1SRzlpUWpGQmdnRUxlbTFCT1ZWT2NIRkhaRm1DQVFzellYVkNhMkZIVlhOMGQ0SUJDMkpzUmxGVGJuRndWRzR3Z2dFTGQwZFBSM2hETVRjeE9HLUNBUXRNVUhCbk56UnpRbWRuTUlJQkMyVjZlbEZFWWpobVZYZE5nZ0VMWVRKb2VucDFiVjlCWlZFJTNEGPHq0C4%3D"
	visitorID = "Tirilil"
	result, err = utils.GqlResolver[lists.Music](
		map[string]any{
			"continuationId": continuationID,
			"visitorId":      visitorID,
		},
		search.MusicsPagination,
	)
	if err != nil {
		t.Errorf("Test case 3 failed: %v", err)
	}
	// Test case 3.1 verify that the result is no nil
	if len(result.Songs) == 0 {
		t.Errorf("Test case 3.1 failed: expected  music list")
	}

	// Test case 4 - invalid continuationID type
	_, err = utils.GqlResolver[lists.Music](
		map[string]any{
			"continuationId": 1,
			"visitorId":      visitorID,
		},
		search.MusicsPagination,
	)
	if err == nil {
		t.Errorf("Test case 4 failed expected an error, got: %v", err)
	}

	// Test case 5 - invalid visitorID type
	_, err = utils.GqlResolver[lists.Music](
		map[string]any{
			"continuationId": continuationID,
			"visitorId":      1,
		},
		search.MusicsPagination,
	)
	if err == nil {
		t.Errorf("Test case 5 failed expected an error, got: %v", err)
	}
}
