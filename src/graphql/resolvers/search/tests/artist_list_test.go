package tests

import (
	"katze/src/graphql/resolvers"
	"katze/src/graphql/resolvers/tests/utils"
	"katze/src/models/lists"
	"testing"
)

func TestArtistList(t *testing.T) {

	// Test case 1 - valid query, params and visitorID
	result, err := utils.GqlResolver[lists.Artists](
		map[string]any{
			"query":     "Lana del rey",
			"params":    "EgWKAQIgAWoKEAMQBBAJEAoQBQ%3D%3D",
			"visitorId": "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D",
		},
		resolvers.ArtistList,
	)
	if err != nil {
		t.Fatalf("Test case 1 failed: %v", err)
	}
	// Test case 1.1 verify that the result is not nil
	if len(result.Artists) == 0 {
		t.Errorf("Test case 1.1 failed: expected non-nil artist list")
	}

	// Test case 2 - invalid params
	result, err = utils.GqlResolver[lists.Artists](
		map[string]any{
			"query":     "Lana del rey",
			"params":    "Tirilil",
			"visitorId": "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D",
		},
		resolvers.ArtistList,
	)
	if err != nil {
		t.Fatalf("Test case 2 failed error: %v", err)
	}
	// Test case 2.1 verify that the result is nil
	if len(result.Artists) != 0 {
		t.Errorf("Test case 2.1 failed: expected nil artist list")
	}

	// Test case 3 - invalid visitorID
	result, err = utils.GqlResolver[lists.Artists](
		map[string]any{
			"query":     "Lana del rey",
			"params":    "EgWKAQIgAWoKEAMQBBAJEAoQBQ%3D%3D",
			"visitorId": "Tirilil",
		},
		resolvers.ArtistList,
	)
	if err != nil {
		t.Errorf("Test case 3 failed error, got: %v", err)
	}

	// Test case 4 - invalid query type
	result, err = utils.GqlResolver[lists.Artists](
		map[string]any{
			"query":     123,
			"params":    "EgWKAQIgAWoKEAMQBBAJEAoQBQ%3D%3D",
			"visitorId": "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D",
		},
		resolvers.ArtistList,
	)
	if err == nil {
		t.Errorf("Test case 4 failed error, got: %v", err)
	}

	// Test case 5 - invalid params type
	result, err = utils.GqlResolver[lists.Artists](
		map[string]any{
			"query":     "Lana del rey",
			"params":    123,
			"visitorId": "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D",
		},
		resolvers.ArtistList,
	)
	if err == nil {
		t.Errorf("Test case 5 failed error, got: %v", err)
	}

	// Test case 6 - invalid visitorID type
	result, err = utils.GqlResolver[lists.Artists](
		map[string]any{
			"query":     "Lana del rey",
			"params":    "EgWKAQIgAWoKEAMQBBAJEAoQBQ%3D%3D",
			"visitorId": 123,
		},
		resolvers.ArtistList,
	)
	if err == nil {
		t.Errorf("Test case 6 failed error, got: %v", err)
	}
}
