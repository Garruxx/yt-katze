package tests

import (
	"katze/src/mappers/search/general/result/best/formatters"
	"katze/src/mappers/search/general/result/best/formatters/tests/data"
	"testing"
)

func TestArtist(t *testing.T) {

	// Test case 1: valid artist
	var cardShelfRenderer = data.ARTIST_MUSIC_CARD_SHELF_VALID
	artist, err := formatters.Artist(cardShelfRenderer)
	if err != nil {
		t.Errorf("test 1 failed: %s", err)
	}
	if artist.Type != "artist" {
		t.Errorf("test 1 failed: type is not artist")
	}
	if artist.Title != "ArtistNameTest" {
		t.Errorf("test 1 failed: title is not ArtistNameTest")
	}
	if artist.ID != "BrowseIDTest" {
		t.Errorf(
			"test 1 failed: browseID expected BrowseIDTest, got %s",
			artist.ID,
		)
	}
	if artist.Thumbnails[0].URL != "testThumbnailURL" {
		t.Errorf(
			"test 1 failed: thumbnail URL expected testThumbnailURL, got %s",
			artist.Thumbnails[0].URL,
		)
	}
	if artist.Thumbnails[0].Width != 100 {
		t.Errorf(
			"test 1 failed: thumbnail width expected 100, got %d",
			artist.Thumbnails[0].Width,
		)
	}
	if artist.Thumbnails[0].Height != 100 {
		t.Errorf(
			"test 1 failed: thumbnail height expected 100, got %d",
			artist.Thumbnails[0].Height,
		)
	}

	// Test case 2: valid artist with no subtitle
	cardShelfRenderer = data.ARTIST_MUSIC_CARD_SHELF_VALID_NO_SUBTITLE
	artist, err = formatters.Artist(cardShelfRenderer)
	if err != nil {
		t.Errorf("test 2 failed: %s", err)
	}
	if artist.Type != "artist" {
		t.Errorf("test 2 failed: type is not artist")
	}
	if artist.Title != "ArtistNameTest" {
		t.Errorf("test 2 failed: title is not ArtistNameTest")
	}
	if artist.ID != "BrowseIDTest" {
		t.Errorf(
			"test 2 failed: browseID expected BrowseIDTest, got %s",
			artist.ID,
		)
	}
	if artist.Thumbnails[0].URL != "testThumbnailURL" {
		t.Errorf(
			"test 2 failed: thumbnail URL expected testThumbnailURL, got %s",
			artist.Thumbnails[0].URL,
		)
	}
	if artist.Thumbnails[0].Width != 100 {
		t.Errorf(
			"test 2 failed: thumbnail width expected 100, got %d",
			artist.Thumbnails[0].Width,
		)
	}
	if artist.Thumbnails[0].Height != 100 {
		t.Errorf(
			"test 2 failed: thumbnail height expected 100, got %d",
			artist.Thumbnails[0].Height,
		)
	}

	// Test case 3: invalid artist
	cardShelfRenderer = data.ARTIST_MUSIC_CARD_SHELF_INVALID
	_, err = formatters.Artist(cardShelfRenderer)
	if err == nil {
		t.Errorf("test 3 failed: expected error, got nil")
	}

	// Test case 4: invalid no browseEndpoint
	cardShelfRenderer = data.ARTIST_MUSIC_CARD_SHELF_INVALID_NO_BROWSE_ENDPOINT
	_, err = formatters.Artist(cardShelfRenderer)
	if err == nil {
		t.Errorf("test 4 failed: expected error, got nil")
	}

	// Test case 5: invalid no type artist
	cardShelfRenderer = data.ARTIST_MUSIC_CARD_SHELF_INVALID_NO_TYPE_ARTIST
	_, err = formatters.Artist(cardShelfRenderer)
	if err == nil {
		t.Errorf("test 5 failed: expected error, got nil")
	}

	// Test case 6: invalid no browse endpoint
	cardShelfRenderer = data.ARTIST_MUSIC_CARD_SHELF_INVALID_NO_BROWSE_ENDPOINT
	_, err = formatters.Artist(cardShelfRenderer)
	if err == nil {
		t.Errorf("test 6 failed: expected error, got nil")
	}

	// Test case 7: invalid no browse endpoint supported configs
	cardShelfRenderer = data.ARTIST_MUSIC_CARD_SHELF_NO_BROWSE_ENDPOINT_CONTEXT_SUPPORTED_CONFIGS
	_, err = formatters.Artist(cardShelfRenderer)
	if err == nil {
		t.Errorf("test 7 failed: expected error, got nil")
	}

	// Test case 8: invalid no browseID
	cardShelfRenderer = data.ARTIST_MUSIC_CARD_SHELF_INVALID_NO_BROWSE_ID
	_, err = formatters.Artist(cardShelfRenderer)
	if err == nil {
		t.Errorf("test 8 failed: expected error, got nil")
	}
}
