package tests

import (
	"katze/src/graphql/schema/common"
	"katze/src/graphql/schema/lists"
	"katze/src/graphql/schema/shelves"
	"testing"

	"github.com/graphql-go/graphql"
)

func TestGeneral(t *testing.T) {

	// Test case 1, bestMatch field
	bestMatchField := shelves.GeneralType.Fields()["bestMatch"]
	if bestMatchField == nil {
		t.Fatalf("Test case 1 failed: expected bestMatch field to be non-nil")
	}
	// Test case 1.1, bestMatch field type
	if bestMatchField.Type != common.BestMatchType {
		t.Fatalf(
			"Test case 1.1 failed: expected bestMatch field type to be BestMatchType",
		)
	}

	// Test case 2, tracks field
	tracksField := shelves.GeneralType.Fields()["tracks"]
	if tracksField == nil {
		t.Fatalf("Test case 2 failed: expected tracks field to be non-nil")
	}
	// Test case 2.1, tracks field type
	if tracksField.Type != lists.MusicType {
		t.Fatalf(
			"Test case 2.1 failed: expected tracks field type to be TracksType",
		)
	}

	// Test case 3, albums field
	albumsField := shelves.GeneralType.Fields()["albums"]
	if albumsField == nil {
		t.Fatalf("Test case 3 failed: expected albums field to be non-nil")
	}
	// Test case 3.1, albums field type
	if albumsField.Type != lists.AlbumType {
		t.Fatalf(
			"Test case 3.1 failed: expected albums field type to be AlbumsType",
		)
	}

	// Test case 4, artists field
	artistsField := shelves.GeneralType.Fields()["artists"]
	if artistsField == nil {
		t.Fatalf("Test case 4 failed: expected artists field to be non-nil")
	}
	// Test case 4.1, artists field type
	if artistsField.Type != lists.ArtistType {
		t.Fatalf(
			"Test case 4.1 failed: expected artists field type to be ArtistsType",
		)
	}

	// Test case 5, visitorId field
	visitorIdField := shelves.GeneralType.Fields()["visitorId"]
	if visitorIdField == nil {
		t.Fatalf("Test case 5 failed: expected visitorId field to be non-nil")
	}
	// Test case 5.1, visitorId field type
	if visitorIdField.Type != graphql.String {
		t.Fatalf(
			"Test case 5.1 failed: expected visitorId field type to be String",
		)
	}
}
