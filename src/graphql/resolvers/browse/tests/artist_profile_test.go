package tests

import (
	"katze/src/graphql/resolvers"
	"katze/src/graphql/resolvers/tests/utils"
	"katze/src/models/shelves"
	"testing"
)

func TestArtistProfile(t *testing.T) {

	// Test case 1 - valid artistID and visitorID
	artistID := "UCPC0L1d253x-KuMNwa05TpA"
	visitorID := "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D"
	result, err := utils.GqlResolver[shelves.Artist](
		map[string]any{
			"artistId":  artistID,
			"visitorId": visitorID,
		},
		resolvers.ArtistProfile,
	)
	if err != nil {
		t.Fatalf("Test case 1 failed: %v", err)
	}
	// Test case 1.1 verify that the result is not nil
	if result.Name == "" {
		t.Errorf("Test case 1.1 failed: expected non-nil artist name")
	}

	// Test case 2 - invalid artistID
	artistID = "Tirilil"
	visitorID = "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D"
	_, err = utils.GqlResolver[shelves.Artist](
		map[string]any{
			"artistId":  artistID,
			"visitorId": visitorID,
		},
		resolvers.ArtistProfile,
	)
	if err == nil {
		t.Errorf("Test case 2 failed expected error but got nil: %v", err)
	}

	// Test case 3 - invalid visitorID type
	artistID = "UCPC0L1d253x-KuMNwa05TpA"
	_, err = utils.GqlResolver[shelves.Artist](
		map[string]any{
			"artistId":  artistID,
			"visitorId": 22,
		},
		resolvers.ArtistProfile,
	)
	if err == nil {
		t.Errorf("Test case 3 failed expected error but got nil: %v", err)
	}

	// Test case 4 - invalid artistId
	visitorID = "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D"
	_, err = utils.GqlResolver[shelves.Artist](
		map[string]any{
			"artistId":  22,
			"visitorId": visitorID,
		},
		resolvers.ArtistProfile,
	)
	if err == nil {
		t.Errorf("Test case 4 failed expected error but got nil: %v", err)
	}
}
