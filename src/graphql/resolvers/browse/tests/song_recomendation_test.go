package tests

import (
	"katze/src/graphql/resolvers"
	"katze/src/graphql/resolvers/tests/utils"
	"katze/src/models/music"
	"testing"
)

func TestSongsRecomendation(t *testing.T) {
	// Test case 1 - valid songID and visitorID
	visitorID := "CgswTi1jc3JLMnVuVSir06GhBg%3D%3D"
	songID := "H9NJenpBV2I"
	result, err := utils.GqlResolver[[]music.Song](
		map[string]any{
			"songId":  songID,
			"visitorId": visitorID,
		},
		resolvers.SongsRecomendation,
	)
	if err != nil {
		t.Fatalf("Test case 1 failed: %v", err)
	}
	// Test case 1.1 verify that the result is not nil
	if len(result) == 0 {
		t.Errorf(
			"Test case 1.1 failed: expected non-nil song recomendation list",
		)
	}

	// Test case 2 - invalid songID
	songID = "Tirilil"
	_, err = utils.GqlResolver[[]music.Song](
		map[string]any{
			"songId":  songID,
			"visitorId": visitorID,
		},
		resolvers.SongsRecomendation,
	)
	if err == nil {
		t.Errorf("Test case 2 failed expected error but got nil: %v", err)
	}

}
