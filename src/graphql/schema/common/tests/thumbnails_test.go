package tests

import (
	"katze/src/graphql/schema/common"
	"testing"

	"github.com/graphql-go/graphql"
)

func TestThumbnails(t *testing.T) {

	// Test case 1, width field
	widthField := common.ThumbnailObjectType.Fields()["width"]
	if widthField == nil {
		t.Fatalf("Test case 1 failed: expected width field to be non-nil")
	}
	// Test case 1.1, width field type
	if widthField.Type != graphql.Int {
		t.Fatalf("Test case 1.1 failed: expected width field type to be int")
	}

	// Test case 2, height field
	heightField := common.ThumbnailObjectType.Fields()["height"]
	if heightField == nil {
		t.Fatalf("Test case 2 failed: expected height field to be non-nil")
	}
	// Test case 2.1, height field type
	if heightField.Type != graphql.Int {
		t.Fatalf("Test case 2.1 failed: expected height field type to be int")
	}

	// Test case 3, url field
	urlField := common.ThumbnailObjectType.Fields()["url"]
	if urlField == nil {
		t.Fatalf("Test case 3 failed: expected url field to be non-nil")
	}
	// Test case 3.1, url field type
	if urlField.Type != graphql.String {
		t.Fatalf("Test case 3.1 failed: expected url field type to be string")
	}

}
