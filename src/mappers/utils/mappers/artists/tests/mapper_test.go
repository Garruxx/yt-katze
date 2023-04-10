package tests

import (
	"katze/src/mappers/utils/mappers/artists"
	"katze/src/mappers/utils/mappers/artists/tests/data"
	"testing"
)

func TestMap(t *testing.T) {

	// test case 1 - items valid two artists
	itemsArtists := data.ITEMS_ARTISTS_VALID_TWO_ARTISTS
	result, err := artists.Mapper(itemsArtists)
	if err != nil {
		t.Errorf("test 1 failed, error: %v", err)
	}
	// test case 1.1 - check if the result has two artists
	if len(result) != 2 {
		t.Errorf("test 1.1 failed, expected 2 artists, got %v", len(result))
	}
	// test case 1.2 - check if the fisrt artist has the correct name
	if name := result[0].Name; name != "ArtistNameTest" {
		t.Errorf("test 1.2 failed, expected ArtistNameTest, got %v", name)
	}

	// test case 2 - items valid one artist
	itemsArtists = data.ITEMS_ARTISTS_VALID_ONE_ARTIST
	result, err = artists.Mapper(itemsArtists)
	if err != nil {
		t.Errorf("test 2 failed, error: %v", err)
	}
	// test case 2.1 - check if the result has one artist
	if leng := len(result); leng != 1 {
		t.Errorf("test 2.1 failed, expected 1 artist, got %v", leng)
	}
	// test case 2.2 - check if the fisrt artist has the correct name
	if name := result[0].Name; name != "ArtistNameTest" {
		t.Errorf("test 2.2 failed, expected ArtistNameTest, got %v", name)
	}

	// test case 3 - items empty
	itemsArtists = data.ITEMS_ARTISTS_VALID_EMPTY
	_, err = artists.Mapper(itemsArtists)
	if err != nil {
		t.Errorf("test 3 failed, error: %v", err)
	}	
	
}
