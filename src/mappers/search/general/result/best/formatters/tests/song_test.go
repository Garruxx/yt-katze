package tests

import (
	"katze/src/mappers/search/general/result/best/formatters"
	"katze/src/mappers/search/general/result/best/formatters/tests/data"
	"testing"
)

func TestSong(t *testing.T) {

	// Test case 1: cardShelfRenderer is valid
	cardShelfRenderer := data.SONG_MUSIC_CARD_SHELF_VALID
	song, err := formatters.Song(cardShelfRenderer)
	if err != nil {
		t.Errorf("test 1 failed: %v", err)
	}

	if song.Type != "song" {
		t.Errorf("test 1 failed: expected %v, got %v", "song", song.Type)
	}

	// Title is "test"
	if song.Title != "test" {
		t.Errorf("test 1 failed: expected %v, got %v", "test", song.Title)
	}

	// ID is "testVideoID"
	if song.WatchID != "testVideoID" {
		t.Errorf("test 1 failed: videoID expected %v, got %v", "testVideoID", song.ID)
	}

	// Thumbnails 1st element URL is "testThumbnailURL"
	if song.Thumbnails == nil {
		t.Errorf("test 1 failed: thumbnails expected %v, got %v", nil, song.Thumbnails)
	}
	if song.Thumbnails[0].URL != "testThumbnailURL" {
		t.Errorf("test 1 failed: thumbnail URL expected %v, got %v", "testThumbnailURL", song.Thumbnails[0].URL)
	}

	// Thumbnails 1st element Width is 100 and Height is 100
	if song.Thumbnails[0].Width != 100 {
		t.Errorf("test 1 failed: thumbnail width expected %v, got %v", 100, song.Thumbnails[0].Width)
	}
	if song.Thumbnails[0].Height != 100 {
		t.Errorf("test 1 failed: thumbnail height expected %v, got %v", 100, song.Thumbnails[0].Height)
	}

	// AlbumTitle is "testAlbumTitle"

	if song.AlbumTitle != "testAlbumTitle" {
		t.Errorf("test 1 failed: albumTitle expected %v, got %v", "testAlbumTitle", song.AlbumTitle)
	}

	// AlbumID is "testAlbumID"
	if song.AlbumID != "testAlbumID" {
		t.Errorf("test 1 failed: albumID expected %v, got %v", "testAlbumID", song.AlbumID)
	}

	// Artists 1st element Name is "testArtistName"
	if song.Artists[0].Name != "testArtistName" {
		t.Errorf("test 1 failed: artist name expected %v, got %v", "testArtistName", song.Artists[0].Name)
	}

	// Artists 1st element ID is "testArtistID"
	if song.Artists[0].ID != "testArtistID" {
		t.Errorf("test 1 failed: artist ID expected %v, got %v", "testArtistID", song.Artists[0].ID)
	}

	// Artist length is 2
	if len(song.Artists) != 2 {
		t.Errorf("test 1 failed: artist length expected %v, got %v", 2, len(song.Artists))
	}

	// Duration is 2:22

	if song.Duration != "2:22" {
		t.Errorf(
			"test 1 failed: duration expected %v, got %v",
			"2:22", song.Duration,
		)
	}

	// Explicit is true
	if !song.Explicit {
		t.Errorf(
			"test 1 failed: explicit expected %v, got %v",
			true, song.Explicit,
		)
	}

	// Test case 2: cardShelfRenderer is no type of music
	cardShelfRenderer = data.SONG_MUSIC_CARD_SHELF_NO_TYPE_MUSIC
	_, err = formatters.Song(cardShelfRenderer)
	if err == nil {
		t.Errorf("test 2 failed: expected error, got %v", err)
	}

	// Test case 3: cardShelfRenderer no artist
	cardShelfRenderer = data.SONG_MUSIC_CARD_SHELF_NO_ARTIST
	song, err = formatters.Song(cardShelfRenderer)
	if err != nil {
		t.Errorf("test 3 failed: %v", err)
	}
	if len(song.Artists) != 0 {
		t.Errorf("test 3 failed: artist length expected %v, got %v", 0, len(song.Artists))
	}
	if song.Title != "test" {
		t.Errorf("test 3 failed: title expected %v, got %v", "test", song.Title)
	}

	// test case 4: cardShelfRenderer no album
	cardShelfRenderer = data.SONG_MUSIC_CARD_SHELF_NO_ALBUM
	song, err = formatters.Song(cardShelfRenderer)
	if err != nil {
		t.Errorf("test 4 failed error, got %v", err)
	}
	if song.AlbumTitle != "Unknown" {
		t.Errorf("test 4 failed: albumTitle expected %v, got %v", nil, song.AlbumTitle)
	}
	if song.AlbumID != "Unknown" {
		t.Errorf("test 4 failed: albumID expected %v, got %v", nil, song.AlbumID)
	}

	// test case 5: cardShelfRenderer no duration
	cardShelfRenderer = data.SONG_MUSIC_CARD_SHELF_NO_DURATION
	song, err = formatters.Song(cardShelfRenderer)
	if err != nil {
		t.Errorf("test 5 failed: %v", err)
	}
	if song.Duration != "" {
		t.Errorf(
			"test 5 failed: duration expected %v, got %v",
			nil, song.Duration,
		)
	}

	// test case 6: cardShelfRenderer invalid artist
	cardShelfRenderer = data.SONG_MUSIC_CARD_SHELF_INVALID_ARTIST
	_, err = formatters.Song(cardShelfRenderer)
	if err == nil {
		t.Errorf("test 6 failed: expected error, got %v", err)
	}

	// test case 7: cardShelfRenderer invalid album
	cardShelfRenderer = data.SONG_MUSIC_CARD_SHELF_INVALID_ALBUM
	_, err = formatters.Song(cardShelfRenderer)
	if err == nil {
		t.Errorf("test 7 failed: expected error, got %v", err)
	}

	// test case 8: cardShelfRenderer empty
	cardShelfRenderer = data.SONG_MUSIC_CARD_SHELF_EMPTY
	_, err = formatters.Song(cardShelfRenderer)
	if err == nil {
		t.Errorf("test 8 failed: expected error, got %v", err)
	}
}
