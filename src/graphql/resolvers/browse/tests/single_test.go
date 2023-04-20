package tests

import (
	"katze/src/graphql/resolvers"
	"katze/src/graphql/resolvers/tests/utils"
	"katze/src/models/shelves"
	"testing"
)

func TestSingle(t *testing.T) {

	// Test case 1 - valid singleId and visitorID
	result, err := utils.GqlResolver[shelves.Single](
		map[string]any{
			"singleId":  "MPREb_EP66MtSPSW7",
			"visitorId": "CgtTay1OakNOTWNIOCiAs4eyhBg%3D%3D",
		},
		resolvers.Single,
	)
	if err != nil {
		t.Fatalf("Test case 1 failed: %v", err)
	}
	if result.Title == "" {
		t.Errorf("Test case 1 failed: expected non-empty list of tabs")
	}

	// Test case 2 - invalid singleId
	_, err = utils.GqlResolver[shelves.Single](
		map[string]any{
			"singleId":  "tirirlil",
			"visitorId": "CgtTay1OakNOTWNIOCiAs4eyhBg%3D%3D",
		},
		resolvers.Single,
	)
	if err == nil {
		t.Errorf("Test case 2 failed expected error but got nil: %v", err)
	}

	// Test case 3 - invalid visitorID
	_, err = utils.GqlResolver[shelves.Single](
		map[string]any{
			"singleId":  "MPREb_EP66MtSPSW7",
			"visitorId": "tirilil",
		},
		resolvers.Single,
	)
	if err != nil {
		t.Errorf("Test case 3 failed: %v", err)
	}

	// Test case 4 - invalid visitorID type
	_, err = utils.GqlResolver[shelves.Single](
		map[string]any{
			"singleId":  "MPREb_EP66MtSPSW7",
			"visitorId": 123,
		},
		resolvers.Single,
	)
	if err == nil {
		t.Errorf("Test case 4 failed expected error but got nil: %v", err)
	}

	// Test case 5 - invalid singleId type
	_, err = utils.GqlResolver[shelves.Single](
		map[string]any{
			"singleId":  123,
			"visitorId": "CgtTay1OakNOTWNIOCiAs4eyhBg%3D%3D",
		},
		resolvers.Single,
	)
	if err == nil {
		t.Errorf("Test case 5 failed expected error but got nil: %v", err)
	}
}
