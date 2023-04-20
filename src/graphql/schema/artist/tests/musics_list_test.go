package tests

import (
	"katze/src/graphql/schema/artist"
	"katze/src/graphql/schema/common"
	"testing"

	"github.com/graphql-go/graphql"
)

func TestMusicList(t *testing.T) {

	// Test case 1, songs field
	songsField := artist.MusicsListType.Fields()["songs"]
	if songsField == nil {
		t.Fatalf("Test case 1 failed, expected songs field no be no-nil")
	}
	// Test case 1.1 songs field type
	if songsField.Type != common.SongsType {
		t.Fatalf(
			"Test case 1.1 failed, extected songs field type lists.SongsType",
		)
	}

	// Test case 2, params field
	paramsField := artist.MusicsListType.Fields()["params"]
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
	continuationIdField := artist.MusicsListType.Fields()["continuationId"]
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
