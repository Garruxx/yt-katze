package tests

import (
	"katze/src/graphql/schema/common"
	"testing"

	"github.com/graphql-go/graphql"
)

func TestBestMatch(t *testing.T) {

	// Test case 1, type field
	typeField := common.BestMatchType.Fields()["type"]
	if typeField == nil {
		t.Fatalf("Test case 1 failed: expected type field to be non-nil")
	}
	// Test case 1.1, type field type
	if typeField.Type != graphql.String {
		t.Fatalf("Test case 1.1 failed: expected type field type to be BestMatchTypeEnum")
	}
	
	// Test case 2, albumType field
	albumTypeField := common.BestMatchType.Fields()["albumType"]
	if albumTypeField == nil {
		t.Fatalf("Test case 2 failed: expected albumType field to be non-nil")
	}
	// Test case 2.1, albumType field type
	if albumTypeField.Type != graphql.String {
		t.Fatalf("Test case 2.1 failed: expected albumType field type to be AlbumTypeEnum")
	}

	// Test case 3, title field
	titleField := common.BestMatchType.Fields()["title"]
	if titleField == nil {
		t.Fatalf("Test case 3 failed: expected title field to be non-nil")
	}
	// Test case 3.1, title field type
	if titleField.Type != graphql.String {
		t.Fatalf("Test case 3.1 failed: expected title field type to be string")
	}

	// Test case 4, id field
	idField := common.BestMatchType.Fields()["id"]
	if idField == nil {
		t.Fatalf("Test case 4 failed: expected id field to be non-nil")
	}
	// Test case 4.1, id field type
	if idField.Type != graphql.String {
		t.Fatalf("Test case 4.1 failed: expected id field type to be string")
	}

	// Test case 5, thumbnails field
	thumbnailsField := common.BestMatchType.Fields()["thumbnails"]
	if thumbnailsField == nil {
		t.Fatalf("Test case 5 failed: expected thumbnails field to be non-nil")
	}
	// Test case 5.1, thumbnails field type
	if thumbnailsField.Type != common.ThumbnailsType {
		t.Fatalf("Test case 5.1 failed: expected thumbnails field type to be ThumbnailsType")
	}

	// Test case 6, watchId field
	watchIdField := common.BestMatchType.Fields()["watchId"]
	if watchIdField == nil {
		t.Fatalf("Test case 6 failed: expected watchId field to be non-nil")
	}
	// Test case 6.1, watchId field type
	if watchIdField.Type != graphql.String {
		t.Fatalf("Test case 6.1 failed: expected watchId field type to be string")
	}

	// Test case 7, artists field
	artistsField := common.BestMatchType.Fields()["artists"]
	if artistsField == nil {
		t.Fatalf("Test case 7 failed: expected artists field to be non-nil")
	}
	// Test case 7.1, artists field type
	if artistsField.Type != common.ArtistsType {
		t.Fatalf("Test case 7.1 failed: expected artists field type to be ArtistType %v", artistsField.Type)
	}

	// Test case 8, album field
	albumField := common.BestMatchType.Fields()["album"]
	if albumField == nil {
		t.Fatalf("Test case 8 failed: expected album field to be non-nil")
	}
	// Test case 8.1, album field type
	if albumField.Type != graphql.String {
		t.Fatalf("Test case 8.1 failed: expected album field type to be AlbumType")
	}

	// Test case 9, duration field
	durationField := common.BestMatchType.Fields()["duration"]
	if durationField == nil {
		t.Fatalf("Test case 9 failed: expected duration field to be non-nil")
	}
	// Test case 9.1, duration field type
	if durationField.Type != graphql.String {
		t.Fatalf("Test case 9.1 failed: expected duration field type to be string")
	}

	// Test case 10, explicit field
	explicitField := common.BestMatchType.Fields()["explicit"]
	if explicitField == nil {
		t.Fatalf("Test case 10 failed: expected explicit field to be non-nil")
	}
	// Test case 10.1, explicit field type
	if explicitField.Type != graphql.Boolean {
		t.Fatalf("Test case 10.1 failed: expected explicit field type to be boolean")
	}

}
