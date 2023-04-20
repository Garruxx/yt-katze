package tests

import (
	"katze/src/graphql/resolvers/search"
	"katze/src/graphql/resolvers/tests/utils"
	"katze/src/models/lists"
	"testing"
)

func TestMusicsList(t *testing.T) {

	// Test case 1 - valid query, params and visitorID
	result, err := utils.GqlResolver[lists.Music](
		map[string]any{
			"query":     "Lana del rey",
			"params":    "EgWKAQIIAWoIEAMQBBAJEAo%3D",
			"visitorId": "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D",
		},
		search.MusicsList,
	)
	if err != nil {
		t.Errorf("Test case 1 failed: %v", err)
	}
	// Test case 1.1 verify that the result is not nil
	if len(result.Songs) == 0 {
		t.Errorf("Test case 1.1 failed: expected non-nil music list")
	}

	// Test case 2 - invalid params
	result, err = utils.GqlResolver[lists.Music](
		map[string]any{
			"query":     "Lana del rey",
			"params":    "Tirilil",
			"visitorId": "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D",
		},
		search.MusicsList,
	)
	if err != nil {
		t.Errorf(
			"Test case 2 failed error: %v", err,
		)
	}
	// Test case 2.1 verify that the result is nil
	if len(result.Songs) != 0 {
		t.Errorf("Test case 2.1 failed: expected nil music list")
	}

	// Test case 3 - invalid visitorID
	result, err = utils.GqlResolver[lists.Music](
		map[string]any{
			"query":     "Lana del rey",
			"params":    "EgWKAQIIAWoIEAMQBBAJEAo%3D",
			"visitorId": "Tirilil",
		},
		search.MusicsList,
	)
	if err != nil {
		t.Errorf("Test case 3 failed error, got: %v", err)
	}
	// Test case 3.1 verify that the result is nil
	if len(result.Songs) == 0 {
		t.Errorf("Test case 3.1 failed: expected music list")
	}

	// Test case 4 - invalid query type
	_, err = utils.GqlResolver[lists.Music](
		map[string]any{
			"query":     123,
			"params":    "EgWKAQIIAWoIEAMQBBAJEAo%3D",
			"visitorId": "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D",
		},
		search.MusicsList,
	)
	if err == nil {
		t.Errorf("Test case 4 failed: expected error")
	}

	// Test case 5 - invalid params type
	_, err = utils.GqlResolver[lists.Music](
		map[string]any{
			"query":     "Lana del rey",
			"params":    123,
			"visitorId": "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D",
		},
		search.MusicsList,
	)
	if err == nil {
		t.Errorf("Test case 5 failed: expected error")
	}

	// Test case 6 - invalid visitorID type
	_, err = utils.GqlResolver[lists.Music](
		map[string]any{
			"query":     "Lana del rey",
			"params":    "EgWKAQIIAWoIEAMQBBAJEAo%3D",
			"visitorId": 123,
		},
		search.MusicsList,
	)
}
