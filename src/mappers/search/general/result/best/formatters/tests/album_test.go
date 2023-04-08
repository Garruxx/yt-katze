package tests

import (
	"katze/src/mappers/search/general/result/best/formatters"
	"katze/src/mappers/search/general/result/best/formatters/tests/data"
	"testing"
)

func TestAlbum(t *testing.T) {

	// Test case 1: cardShelfRenderer is valid
	cardShelfRenderer := data.ALBUM_MUSIC_CARD_SHELF_VALID
	album, err := formatters.Album(cardShelfRenderer)
	if err != nil {
		t.Errorf("test 1 failed: %v", err)
	}

	if album.Type != "album" {
		t.Errorf(
			"test 1 failed: expected %v, got %v",
			"album", album.Type,
		)
	}

	// Title is "AlbumTitleTest"
	if album.Title != "AlbumTitleTest" {
		t.Errorf(
			"test 1 failed: expected %v, got %v",
			"AlbumTitleTest", album.Title,
		)
	}

	// ID is "BrowseIDTest"
	if album.ID != "BrowseIDTest" {
		t.Errorf(
			"test 1 failed: videoID expected %v, got %v",
			"BrowseIDTest", album.ID,
		)
	}

	// Thumbnails 1st element URL is "testThumbnailURL"
	if album.Thumbnails == nil {
		t.Errorf(
			"test 1 failed: thumbnails expected %v, got %v",
			nil, album.Thumbnails,
		)
	}
	if album.Thumbnails[0].URL != "testThumbnailURL" {
		t.Errorf(
			"test 1 failed: thumbnail URL expected %v, got %v",
			"testThumbnailURL", album.Thumbnails[0].URL,
		)
	}

	// Thumbnails 1st element Width is 100 and Height is 100
	if album.Thumbnails[0].Width != 100 {
		t.Errorf(
			"test 1 failed: thumbnail width expected %v, got %v",
			100, album.Thumbnails[0].Width,
		)
	}
	if album.Thumbnails[0].Height != 100 {
		t.Errorf(
			"test 1 failed: thumbnail height expected %v, got %v",
			100, album.Thumbnails[0].Height,
		)
	}

	// Artists 1st element Name is "testArtistName"
	if album.Artists[0].Name != "testArtistName" {
		t.Errorf(
			"test 1 failed: artist name expected %v, got %v",
			"testArtistName", album.Artists[0].Name,
		)
	}

	// Artists 1st element ID is "testArtistID"
	if album.Artists[0].ID != "testArtistID" {
		t.Errorf(
			"test 1 failed: artist ID expected %v, got %v",
			"testArtistID", album.Artists[0].ID,
		)
	}

	// Just have an artist
	if len(album.Artists) != 1 {
		t.Errorf(
			"test 1 failed: just one artist expected, got %v",
			len(album.Artists),
		)
	}

	// Test case 2: cardShelfRenderer is empty
	cardShelfRenderer = data.ALBUM_MUSIC_CARD_SHELF_EMPTY
	_, err = formatters.Album(cardShelfRenderer)
	if err == nil {
		t.Errorf("test 2 failed: expected error, got nil")
	}

	// Test case 3: cardShelfRenderer no title
	cardShelfRenderer = data.ALBUM_MUSIC_CARD_SHELF_NO_TITLE
	_, err = formatters.Album(cardShelfRenderer)
	if err == nil {
		t.Errorf("test 3 failed: expected error, got nil")
	}

	// Test case 4: cardShelfRenderer no Album ID
	cardShelfRenderer = data.ALBUM_MUSIC_CARD_SHELF_NO_ALBUM_ID
	_, err = formatters.Album(cardShelfRenderer)
	if err == nil {
		t.Errorf("test 4 failed: expected error, got nil")
	}

	// Test case 5: cardShelfRenderer no thumbnails
	cardShelfRenderer = data.ALBUM_MUSIC_CARD_SHELF_NO_THUMBNAILS
	_, err = formatters.Album(cardShelfRenderer)
	if err != nil {
		t.Errorf("test 5 failed: %v", err)
	}

	// Test case 6: cardShelfRenderer no artists
	cardShelfRenderer = data.ALBUM_MUSIC_CARD_SHELF_NO_ARTIST
	_, err = formatters.Album(cardShelfRenderer)
	if err != nil {
		t.Errorf("test 6 failed: expected error, got nil")
	}
}
