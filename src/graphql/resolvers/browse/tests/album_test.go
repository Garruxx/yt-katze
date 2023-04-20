package tests

import (
	"katze/src/graphql/resolvers"
	"katze/src/graphql/resolvers/tests/utils"
	"katze/src/models/shelves"
	"testing"
)

func TestAlbum(t *testing.T) {

	// Test case 1 - valid albumID and visitorID
	albumID := "MPREb_R1STxFWpP62"
	visitorID := "CgtTay1OakNOTWNIOCiAs4eyhBg%3D%3D"
	result, err := utils.GqlResolver[shelves.Album](
		map[string]any{
			"albumId":   albumID,
			"visitorId": visitorID,
		},
		resolvers.Album,
	)
	if err != nil {
		t.Errorf("Test case 1 failed: %v", err)
	}
	if result.Title == "" {
		t.Errorf("Test case 1 failed: expected non-empty list of tabs")
	}

	// Test case 2 - invalid albumID
	albumID = "Tirilil"
	visitorID = "CgtTay1OakNOTWNIOCiAs4eyhBg%3D%3D"
	_, err = utils.GqlResolver[shelves.Album](
		map[string]any{
			"albumId":   albumID,
			"visitorId": visitorID,
		},
		resolvers.Album,
	)
	if err == nil {
		t.Errorf("Test case 2 failed expected error but got nil: %v", err)
	}

	// Test case 3 - invalid visitorID
	visitorID = "Tirilil"
	albumID = "MPREb_R1STxFWpP62"
	_, err = utils.GqlResolver[shelves.Album](
		map[string]any{
			"albumId":   albumID,
			"visitorId": visitorID,
		},
		resolvers.Album,
	)
	if err != nil {
		t.Errorf("Test case 3 failed: %v", err)
	}

	// Test case 4 - invalid visitorID type
	albumID = "MPREb_R1STxFWpP62"
	_, err = utils.GqlResolver[shelves.Album](
		map[string]any{
			"albumId":   albumID,
			"visitorId": 123,
		},
		resolvers.Album,
	)
	if err == nil {
		t.Errorf("Test case 4 failed expected error but got nil: %v", err)
	}

	// Test case 5 - invalid albumID type
	visitorID = "CgtTay1OakNOTWNIOCiAs4eyhBg%3D%3D"
	_, err = utils.GqlResolver[shelves.Album](
		map[string]any{
			"albumId":   4444,
			"visitorId": visitorID,
		},
		resolvers.Album,
	)
}
