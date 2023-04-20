package tests

import (
	"katze/src/graphql/schema/common"
	"testing"

	"github.com/graphql-go/graphql"
)

func TestContinuation(t *testing.T) {

	// Test case 1, query field
	queryField := common.ContinuationType.Fields()["query"]
	if queryField == nil {
		t.Fatalf("Test case 1 failed: expected query field to be non-nil")
	}
	// Test case 1.1, query field type
	if queryField.Type != graphql.String {
		t.Fatalf(
			"Test case 1.1 failed: expected query field type to be QueryType",
		)
	}

	// Test case 2, params field
	paramsField := common.ContinuationType.Fields()["params"]
	if paramsField == nil {
		t.Fatalf("Test case 2 failed: expected params field to be non-nil")
	}
	// Test case 2.1, params field type
	if paramsField.Type != graphql.String {
		t.Fatalf(
			"Test case 2.1 failed: expected params field type to be string",
		)
	}
}
