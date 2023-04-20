package tests

import (
	"katze/src/graphql/resolvers"
	"katze/src/graphql/resolvers/tests/utils"
	"katze/src/models/shelves"
	"testing"
)

func TestGeneral(t *testing.T) {

	// Test case 1 - valid query and visitorID

	result, err := utils.GqlResolver[shelves.General](
		map[string]any{
			"query":     "Lana del rey",
			"visitorId": "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D",
		},
		resolvers.General,
	)
	if err != nil {
		t.Errorf("Test case 1 failed: %v", err)
	}
	// Test case 1.1 verify that the result songs is not nil
	if len(result.Tracks.Songs) == 0 {
		t.Errorf("Test case 1.1 failed: expected non-nil music list")
	}
	// Test case 1.2 verify that the result artist is not nil
	if len(result.Artists.Artists) == 0 {
		t.Errorf("Test case 1.2 failed: expected non-nil artist list")
	}
	// Test case 1.3 verify that the result albums is not nil
	if len(result.Albums.Albums) == 0 {
		t.Errorf("Test case 1.3 failed: expected non-nil album list")
	}

	// Test case 2 - invalid visitorID
	result, err = utils.GqlResolver[shelves.General](
		map[string]any{
			"query":     "Lana del rey",
			"visitorId": "tirilil",
		},
		resolvers.General,
	)
	if err != nil {
		t.Errorf("Test case 2 failed error, got: %v", err)
	}
	// Test case 2.1 verify that the result songs is not nil
	if len(result.Tracks.Songs) == 0 {
		t.Errorf("Test case 2.1 failed: expected non-nil music list")
	}
	// Test case 2.2 verify that the result artist is not nil
	if len(result.Artists.Artists) == 0 {
		t.Errorf("Test case 2.2 failed: expected non-nil artist list")
	}
	// Test case 2.3 verify that the result albums is not nil
	if len(result.Albums.Albums) == 0 {
		t.Errorf("Test case 2.3 failed: expected non-nil album list")
	}
	// Test case 2.4 verify that the result has a visitorID
	if result.VisitorID == "" {
		t.Errorf("Test case 2.4 failed: expected visitorID")
	}

	// Test case 3 - query no results
	result, err = utils.GqlResolver[shelves.General](
		map[string]any{
			"query":     "wliuneiruteorteteon",
			"visitorId": "dnUUF4QUVFQWtRQ29JQkMyTXRWRWRXVEMxeU",
		},
		resolvers.General,
	)
	if err != nil {
		t.Errorf("Test case 3 failed error, got: %v", err)
	}
	// Test case 3.1 verify that the result songs is not nil
	if len(result.Tracks.Songs) != 0 {
		t.Errorf("Test case 3.1 failed: expected nil music list")
	}

	// Test case 3.1 invalid query type
	_, err = utils.GqlResolver[shelves.General](
		map[string]any{
			"query":     123,
			"visitorId": "dnUUF4QUVFQWtRQ29JQkMyTXRWRWRXVEMxeU",
		},
		resolvers.General,
	)
	if err == nil {
		t.Errorf("Test case 3.1 failed: expected error")
	}

	// Test case 3.2 invalid visitorID type
	_, err = utils.GqlResolver[shelves.General](
		map[string]any{
			"query":     "Lana del rey",
			"visitorId": 123,
		},
		resolvers.General,
	)
	if err == nil {
		t.Errorf("Test case 3.2 failed: expected error")
	}
}
