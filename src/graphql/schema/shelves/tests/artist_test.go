package tests

import (
	"katze/src/graphql/schema/artist"
	"katze/src/graphql/schema/common"
	"katze/src/graphql/schema/shelves"
	"testing"

	"github.com/graphql-go/graphql"
)

func TestArtist(t *testing.T) {

	// Test case 1, name field
	nameField := shelves.ArtistType.Fields()["name"]
	if nameField == nil {
		t.Fatalf("Test case 1 failed: expected name field to be non-nil")
	}
	// Test case 1.1, name field type
	if nameField.Type != graphql.String {
		t.Fatalf("Test case 1.1 failed: expected name field type to be string")
	}

	// Test case 2, thumbnails field
	thumbnailsField := shelves.ArtistType.Fields()["thumbnails"]
	if thumbnailsField == nil {
		t.Fatalf("Test case 2 failed: expected thumbnails field to be non-nil")
	}
	// Test case 2.1, thumbnails field type
	if thumbnailsField.Type != common.ThumbnailsType {
		t.Fatalf(
			"Test case 2.1 failed: expected thumbnails field type to be ThumbnailsType",
		)
	}

	// Test case 3, musicsList field
	musicsListField := shelves.ArtistType.Fields()["musicsList"]
	if musicsListField == nil {
		t.Fatalf("Test case 3 failed: expected musicsList field to be non-nil")
	}
	// Test case 3.1, musicsList field type
	if musicsListField.Type != artist.MusicsListType {
		t.Fatalf(
			"Test case 3.1 failed: expected musicsList field type to be MusicsListType",
		)
	}

	// Test case 4, albumsList field
	albumsListField := shelves.ArtistType.Fields()["albumsList"]
	if albumsListField == nil {
		t.Fatalf("Test case 4 failed: expected albumsList field to be non-nil")
	}
	// Test case 4.1, albumsList field type
	if albumsListField.Type != artist.AlbumsListType {
		t.Fatalf(
			"Test case 4.1 failed: expected albumsList field type to be AlbumsListType",
		)
	}

	// Test case 5, singlesList field
	singlesListField := shelves.ArtistType.Fields()["singlesList"]
	if singlesListField == nil {
		t.Fatalf("Test case 5 failed: expected singlesList field to be non-nil")
	}
	// Test case 5.1, singlesList field type
	if singlesListField.Type != artist.SinglesListType {
		t.Fatalf(
			"Test case 5.1 failed: expected singlesList field type to be SinglesListType",
		)
	}

	// Test case 6, visitorId field
	visitorIdField := shelves.ArtistType.Fields()["visitorId"]
	if visitorIdField == nil {
		t.Fatalf("Test case 6 failed: expected visitorId field to be non-nil")
	}
	// Test case 6.1, visitorId field type
	if visitorIdField.Type != graphql.String {
		t.Fatalf(
			"Test case 6.1 failed: expected visitorId field type to be string",
		)
	}
}
