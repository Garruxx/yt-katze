package tests

import (
	"katze/src/graphql/resolvers"
	"katze/src/graphql/resolvers/tests/utils"
	"katze/src/models/lists"
	"testing"
)

func TestAlbumList(t *testing.T) {

	// Test case 1 - valid query, params and visitorID

	result, err := utils.GqlResolver[lists.Albums](
		map[string]any{
			"query":     "album",
			"params":    "EgWKAQIYAWoKEAMQBBAJEAoQBQ%3D%3D",
			"visitorId": "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D",
		},
		resolvers.AlbumsList,
	)
	if err != nil {
		t.Fatalf("Test case 1 failed: %v", err)
	}
	// Test case 1.1 verify that the result is not nil
	if len(result.Albums) == 0 {
		t.Errorf("Test case 1.1 failed: expected non-nil album list")
	}

	// Test case 2 - invalid params
	result, err = utils.GqlResolver[lists.Albums](
		map[string]any{
			"query":     "album",
			"params":    "Tirilil",
			"visitorId": "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D",
		},
		resolvers.AlbumsList,
	)
	if err != nil {
		t.Errorf("Test case 2 failed error: %v", err)
	}
	// Test case 2.1 verify that the result is nil
	if len(result.Albums) != 0 {
		t.Errorf("Test case 2.1 failed: expected nil album list")
	}

	// Test case 3 - invalid visitorID (no error)
	result, err = utils.GqlResolver[lists.Albums](
		map[string]any{
			"query":     "album",
			"params":    "EgWKAQIYAWoKEAMQBBAJEAoQBQ%3D%3D",
			"visitorId": "Tirilil",
		},
		resolvers.AlbumsList,
	)
	if err != nil {
		t.Errorf("Test case 3 failed error, got: %v", err)
	}

	// Test case 4 - invalid query type
	result, err = utils.GqlResolver[lists.Albums](
		map[string]any{
			"query":     22,
			"params":    "EgWKAQIYAWoKEAMQBBAJEAoQBQ%3D%3D",
			"visitorId": "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D",
		},
		resolvers.AlbumsList,
	)
	if err == nil {
		t.Errorf("Test case 4 failed: expected error")
	}

	// Test case 5 - invalid params type
	result, err = utils.GqlResolver[lists.Albums](
		map[string]any{
			"query":     "album",
			"params":    22,
			"visitorId": "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D",
		},
		resolvers.AlbumsList,
	)
	if err == nil {
		t.Errorf("Test case 5 failed: expected error")
	}

	// Test case 6 - invalid visitorID type
	result, err = utils.GqlResolver[lists.Albums](
		map[string]any{
			"query":     "album",
			"params":    "EgWKAQIYAWoKEAMQBBAJEAoQBQ%3D%3D",
			"visitorId": 22,
		},
		resolvers.AlbumsList,
	)
	if err == nil {
		t.Errorf("Test case 6 failed: expected error")
	}
}
