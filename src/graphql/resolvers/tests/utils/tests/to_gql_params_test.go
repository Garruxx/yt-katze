package tests

import (
	"katze/src/graphql/resolvers"
	"katze/src/graphql/resolvers/tests/utils"
	"katze/src/models/lists"
	"katze/src/models/shelves"
	"testing"
)

func TestGqlResolver(t *testing.T) {

	// Test case 1 - valid query, visitorId and resolver
	query := "Lana del rey"
	visitorId := ""
	result, err := utils.GqlResolver[shelves.General](
		map[string]any{
			"query":     query,
			"visitorId": visitorId,
		},
		resolvers.General,
	)
	if err != nil {
		t.Errorf("Test case 1 failed: %v", err)
	}
	// Test case 1.1 verify that the result is not nil
	if result.VisitorID == "" {
		t.Errorf("Test case 1.1 failed: expected non-nil result")
	}

	// Test case 2 - valid query, params, visitorId and resolver
	query = "Lana del rey"
	visitorId = ""
	params := "EgWKAQIIAWoIEAMQBBAJEAo%3D"
	result2, err := utils.GqlResolver[lists.Music](
		map[string]any{
			"query":     query,
			"params":    params,
			"visitorId": visitorId,
		},
		resolvers.MusicsList,
	)
	if err != nil {
		t.Errorf("Test case 2 failed: %v", err)
	}
	// Test case 2.1 verify that the result is not nil
	if len(result2.Songs) == 0 {
		t.Errorf("Test case 2.1 failed: expected non-nil result")
	}
}
