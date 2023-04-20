package tests

import (
	"katze/src/graphql/schema/common"
	"testing"

	"github.com/graphql-go/graphql"
)

func TestCardItem(t *testing.T) {

	// Test case 1, thumbnails field
	thumbnailsField := common.CardItemObjectType.Fields()["thumbnails"]
	if thumbnailsField == nil {
		t.Fatalf("Test case 1 failed: expected thumbnails field to be non-nil")
	}
	// Test case 1.1, thumbnails field type
	if thumbnailsField.Type != common.ThumbnailsType {
		t.Fatalf("Test case 1.1 failed: expected thumbnails field type to be ThumbnailListType")
	}

	// Test case 2, title field
	titleField := common.CardItemObjectType.Fields()["title"]
	if titleField == nil {
		t.Fatalf("Test case 2 failed: expected title field to be non-nil")
	}
	// Test case 2.1, title field type
	if titleField.Type != graphql.String {
		t.Fatalf(
			"Test case 2.1 failed: expected title field type to be TitleType",
		)
	}

	// Test case 3, year field
	yearField := common.CardItemObjectType.Fields()["year"]
	if yearField == nil {
		t.Fatalf("Test case 3 failed: expected year field to be non-nil")
	}
	// Test case 3.1, year field type
	if yearField.Type != graphql.Int {
		t.Fatalf(
			"Test case 3.1 failed: expected year field type to be YearType",
		)
	}

	// Test case 4, browseId field
	browseIdField := common.CardItemObjectType.Fields()["browseId"]
	if browseIdField == nil {
		t.Fatalf("Test case 4 failed: expected browseId field to be non-nil")
	}
	// Test case 4.1, browseId field type
	if browseIdField.Type != graphql.String {
		t.Fatalf(
			`
			  Test case 4.1 failed: expected browseId field type to be
		      BrowseIdType
			`,
		)
	}

	// Test case 5, params field
	paramsField := common.CardItemObjectType.Fields()["params"]
	if paramsField == nil {
		t.Fatalf("Test case 5 failed: expected params field to be non-nil")
	}
	// Test case 5.1, params field type
	if paramsField.Type != graphql.String {
		t.Fatalf(
			"Test case 5.1 failed: expected params field type to be ParamsType",
		)
	}

	// Test case 6, explicit field
	explicitField := common.CardItemObjectType.Fields()["explicit"]
	if explicitField == nil {
		t.Fatalf("Test case 6 failed: expected explicit field to be non-nil")
	}
	// Test case 6.1, explicit field type
	if explicitField.Type != graphql.Boolean {
		t.Fatalf(
			"Test case 6.1 failed: expected explicit field type to be ExplicitType",
		)
	}

}
