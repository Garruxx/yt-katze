package tests

import (
	"katze/src/graphql/resolvers"
	"katze/src/graphql/resolvers/tests/utils"
	"katze/src/models/lists"
	"testing"
)

func TestArtistMusicList(t *testing.T) {

	// Test case 1 - valid continuationId and visitorID
	continuationId := "VLOLAK5uy_nD0V5fFISV9uMmvPTFxKITDbmLUqAMeRY"
	params := "ggMCCAI="
	visitorID := "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D"
	result, err := utils.GqlResolver[lists.Music](
		map[string]any{
			"continuationId": continuationId,
			"params":     params,
			"visitorId":  visitorID,
		},
		resolvers.ArtistMusicsList,
	)
	if err != nil {
		t.Fatalf("Test case 1 failed: %v", err)
	}
	// Test case 1.1 verify that the result is not nil
	if len(result.Songs) == 0 {
		t.Errorf("Test case 1.1 failed: expected non-nil artist music list")
	}

	// Test case 2 - invalid continuationId
	continuationId = "Tirilil"
	visitorID = "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D"
	_, err = utils.GqlResolver[lists.Music](
		map[string]any{
			"continuationId": continuationId,
			"params":     params,
			"visitorId":  visitorID,
		},
		resolvers.ArtistMusicsList,
	)
	if err == nil {
		t.Errorf("Test case 2 failed expected error but got nil: %v", err)
	}

	// Test case 3 - invalid visitorID type
	continuationId = "VLOLAK5uy_nD0V5fFISV9uMmvPTFxKITDbmLUqAMeRY"
	_, err = utils.GqlResolver[lists.Music](
		map[string]any{
			"continuationId": continuationId,
			"params":     params,
			"visitorId":  4444,
		},
		resolvers.ArtistMusicsList,
	)
	if err == nil {
		t.Errorf("Test case 3 failed expected error but got nil: %v", err)
	}

	// Test case 4 - invalid params type
	continuationId = "VLOLAK5uy_nD0V5fFISV9uMmvPTFxKITDbmLUqAMeRY"
	_, err = utils.GqlResolver[lists.Music](
		map[string]any{
			"continuationId": continuationId,
			"params":     4444,
			"visitorId":  visitorID,
		},
		resolvers.ArtistMusicsList,
	)
	if err == nil {
		t.Errorf("Test case 4 failed expected error but got nil: %v", err)
	}

	// Test case 5 - invalid continuationId type
	_, err = utils.GqlResolver[lists.Music](
		map[string]any{
			"continuationId": 4444,
			"params":     params,
			"visitorId":  visitorID,
		},
		resolvers.ArtistMusicsList,
	)
	if err == nil {
		t.Errorf("Test case 5 failed expected error but got nil: %v", err)
	}
}
