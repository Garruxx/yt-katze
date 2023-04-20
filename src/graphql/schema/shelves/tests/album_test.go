package tests

import (
	"katze/src/graphql/schema/common"
	"katze/src/graphql/schema/shelves"
	"testing"

	"github.com/graphql-go/graphql"
)

func TestAlbumType(t *testing.T) {

	// Test case 1, title field
	titleField := shelves.AlbumType.Fields()["title"]
	if titleField == nil {
		t.Fatalf("Test case 1 failed, expected title field to be non-nil")
	}
	// Test case 1.1, title field type
	if titleField.Type != graphql.String {
		t.Fatalf("Test case 1.1 failed, expected title field type to be TitleType")
	}

	// Test case 2, artist field
	artistField := shelves.AlbumType.Fields()["artists"]
	if artistField == nil {
		t.Fatalf("Test case 2 failed, expected artist field to be non-nil")
	}
	// Test case 2.1, artist field type
	if artistField.Type != common.ArtistsType {
		t.Fatalf(
			"Test case 2.1 failed, expected artist field type to be ArtistType",
		)
	}

	// Test case 3, thumbnails field
	thumbnailsField := shelves.AlbumType.Fields()["thumbnails"]
	if thumbnailsField == nil {
		t.Fatalf("Test case 3 failed, expected thumbnails field to be non-nil")
	}
	// Test case 3.1, thumbnails field type
	if thumbnailsField.Type != common.ThumbnailsType {
		t.Fatalf(
			"Test case 3.1 failed, expected thumbnails field type to be ThumbnailsType",
		)
	}

	// Test case 4, tracks field
	tracksField := shelves.AlbumType.Fields()["tracks"]
	if tracksField == nil {
		t.Fatalf("Test case 4 failed, expected tracks field to be non-nil")
	}
	// Test case 4.1, tracks field type
	if tracksField.Type != common.SongsType {
		t.Fatalf(
			"Test case 4.1 failed, expected tracks field type to be TracksType",
		)
	}

	// Test case 5, duration field
	durationField := shelves.AlbumType.Fields()["duration"]
	if durationField == nil {
		t.Fatalf("Test case 5 failed, expected duration field to be non-nil")
	}
	// Test case 5.1, duration field type
	if durationField.Type != graphql.String {
		t.Fatalf(
			"Test case 5.1 failed, expected duration field type to be DurationType",
		)
	}

	// Test case 6, year field
	yearField := shelves.AlbumType.Fields()["year"]
	if yearField == nil {
		t.Fatalf("Test case 6 failed, expected year field to be non-nil")
	}
	// Test case 6.1, year field type
	if yearField.Type != graphql.Int {
		t.Fatalf(
			"Test case 6.1 failed, expected year field type to be YearType",
		)
	}

	// Test case 7, visitorId field
	visitorIdField := shelves.AlbumType.Fields()["visitorId"]
	if visitorIdField == nil {
		t.Fatalf("Test case 7 failed, expected visitorId field to be non-nil")
	}
	// Test case 7.1, visitorId field type
	if visitorIdField.Type != graphql.String {
		t.Fatalf(
			"Test case 7.1 failed, expected visitorId field type to be VisitorIdType",
		)
	}
}
