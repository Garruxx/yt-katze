package tests

import (
	"katze/src/graphql/schema/common"
	"testing"

	"github.com/graphql-go/graphql"
)

func TestSong(t *testing.T) {

	// Test case 1, title field
	titleField := common.SongObjectType.Fields()["title"]
	if titleField == nil {
		t.Fatalf("Test case 1 failed: expected title field to be non-nil")
	}
	// Test case 1.1, title field type
	if titleField.Type != graphql.String {
		t.Fatalf("Test case 1.1 failed: expected title field type to be string")
	}

	// Test case 2, id field
	idField := common.SongObjectType.Fields()["id"]
	if idField == nil {
		t.Fatalf("Test case 2 failed: expected id field to be non-nil")
	}
	// Test case 2.1, id field type
	if idField.Type != graphql.String {
		t.Fatalf("Test case 2.1 failed: expected id field type to be string")
	}

	// Test case 3, artist field
	artistField := common.SongObjectType.Fields()["artists"]
	if artistField == nil {
		t.Fatalf("Test case 3 failed: expected artist field to be non-nil")
	}
	// Test case 3.1, artist field type
	if artistField.Type != common.ArtistsType {
		t.Fatalf(
			"Test case 3.1 failed: expected artist field type to be ArtistType",
		)
	}

	// Test case 4, album field
	albumField := common.SongObjectType.Fields()["album"]
	if albumField == nil {
		t.Fatalf("Test case 4 failed: expected album field to be non-nil")
	}
	// Test case 4.1, album field type
	if albumField.Type != graphql.String {
		t.Fatalf(
			"Test case 4.1 failed: expected album field type to be AlbumType",
		)
	}

	// Test case 5, albumId field
	albumIDField := common.SongObjectType.Fields()["albumId"]
	if albumIDField == nil {
		t.Fatalf("Test case 5 failed: expected albumID field to be non-nil")
	}
	// Test case 5.1, albumId field type
	if albumIDField.Type != graphql.String {
		t.Fatalf(
			"Test case 5.1 failed: expected albumID field type to be string",
		)
	}

	// Test case 6, duration field
	durationField := common.SongObjectType.Fields()["duration"]
	if durationField == nil {
		t.Fatalf("Test case 6 failed: expected duration field to be non-nil")
	}

	// Test case 6.1, duration field type
	if durationField.Type != graphql.String {
		t.Fatalf(
			"Test case 6.1 failed: expected duration field type to be string",
		)
	}

	// Test case 7, explicit field
	explicitField := common.SongObjectType.Fields()["explicit"]
	if explicitField == nil {
		t.Fatalf("Test case 7 failed: expected explicit field to be non-nil")
	}
	// Test case 7.1, explicit field type
	if explicitField.Type != graphql.Boolean {
		t.Fatalf(
			"Test case 7.1 failed: expected explicit field type to be boolean",
		)
	}

	// Test case 8, thumbnails field
	thumbnailsField := common.SongObjectType.Fields()["thumbnails"]
	if thumbnailsField == nil {
		t.Fatalf("Test case 8 failed: expected thumbnails field to be non-nil")
	}
	// Test case 8.1, thumbnails field type
	if thumbnailsField.Type != common.ThumbnailsType {
		t.Fatalf(
			"Test case 8.1 failed: expected thumbnails field type to be ThumbnailsType",
		)
	}

	// Test case 9, trackNumber field
	trackNumberField := common.SongObjectType.Fields()["trackNumber"]
	if trackNumberField == nil {
		t.Fatalf("Test case 9 failed: expected trackNumber field to be non-nil")
	}
	// Test case 9.1, trackNumber field type
	if trackNumberField.Type != graphql.Int {
		t.Fatalf(
			"Test case 9.1 failed: expected trackNumber field type to be int",
		)
	}

}
