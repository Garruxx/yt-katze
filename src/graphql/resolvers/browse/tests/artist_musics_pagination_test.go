package tests

import (
	"katze/src/graphql/resolvers"
	"katze/src/graphql/resolvers/tests/utils"
	"katze/src/models/lists"
	"testing"
)

func TestArtistMusicPagination(t *testing.T) {
	// Test case 1 - valid continuationID and visitorID
	continuationID := "4qmFsgJFEitWTE9MQUs1dXlfbkQwVjVmRklTVjl1TW12UFRGeEtJVERibUxVcUFNZVJZGhZlZ1pRVkRwRFIyZVNBUU1JdWdRJTNE"
	visitorID := "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D"
	result, err := utils.GqlResolver[lists.Music](
		map[string]any{
			"continuationId": continuationID,
			"visitorId":      visitorID,
		},
		resolvers.ArtistMusicsListPagination,
	)
	if err != nil {
		t.Fatalf("Test case 1 failed: %v", err)
	}
	// Test case 1.1 verify that the result is not nil
	if len(result.Songs) == 0 {
		t.Errorf("Test case 1.1 failed: expected non-nil artist music list")
	}

	// Test case 2 - invalid continuationID
	continuationID = "Tirilil"
	visitorID = "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D"
	_, err = utils.GqlResolver[lists.Music](
		map[string]any{
			"continuationId": continuationID,
			"visitorId":      visitorID,
		},
		resolvers.ArtistMusicsListPagination,
	)
	if err == nil {
		t.Fatalf("Test case 2 failed expected error but got nil: %v", err)
	}

	// Test case 3 - invalid visitorID type 
	continuationID = "4qmFsgJFEitWTE9MQUs1dXlfbkQwVjVmRklTVjl1TW12UFRGeEtJVERibUxVcUFNZVJZGhZlZ1pRVkRwRFIyZVNBUU1JdWdRJTNE"
	_, err = utils.GqlResolver[lists.Music](
		map[string]any{
			"continuationId": continuationID,
			"visitorId":      true,
		},
		resolvers.ArtistMusicsListPagination,
	)
	if err == nil {
		t.Errorf("Test case 3 failed expected error but got nil: %v", err)
	}

	// Test case 4 - invalid visitorID type
	continuationID = "4qmFsgJFEitWTE9MQUs1dXlfbkQwVjVmRklTVjl1TW12UFRGeEtJVERibUxVcUFNZVJZGhZlZ1pRVkRwRFIyZVNBUU1JdWdRJTNE"
	_, err = utils.GqlResolver[lists.Music](
		map[string]any{
			"continuationId": continuationID,
			"visitorId":      4444,
		},
		resolvers.ArtistMusicsListPagination,
	)
	if err == nil {
		t.Errorf("Test case 4 failed expected error but got nil: %v", err)
	}
}
