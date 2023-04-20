package tests

import (
	"katze/src/graphql/schema/artist"
	"katze/src/graphql/schema/common"
	"testing"

	"github.com/graphql-go/graphql"
)

func TestSingleList(t *testing.T) {

	// Test case 1, singles field
	singlesField := artist.SinglesListType.Fields()["singles"]
	if singlesField == nil {
		t.Fatalf("Test case 1 failed, expected singles field no be no-nil")
	}
	// Test case 1.1 singles field type
	if singlesField.Type != common.CardItemsType {
		t.Fatalf(
			"Test case 1.1 failed, extected singles field type artist.SinglesType",
		)
	}

	// Test case 2, params field
	paramsField := artist.SinglesListType.Fields()["params"]
	if paramsField == nil {
		t.Fatalf("Test case 2 failed, expected params field no be no-nil")
	}
	// Test case 2.1 params field type
	if paramsField.Type != graphql.String {
		t.Fatalf(
			"Test case 2.1 failed, extected params field type graphql.String",
		)
	}

	// Test case 3, continuationId field
	continuationIdField := artist.SinglesListType.Fields()["continuationId"]
	if continuationIdField == nil {
		t.Fatalf(
			"Test case 3 failed, expected continuationId field no be no-nil",
		)
	}
	// Test case 3.1 continuationId field type
	if continuationIdField.Type != graphql.String {
		t.Fatalf(
			"Test case 3.1 failed, extected continuationId field type graphql.String",
		)
	}

}
