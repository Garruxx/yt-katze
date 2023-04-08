package tests

import (
	"katze/src/mappers/search/general/result/best"
	"katze/src/mappers/search/general/result/best/tests/data"
	"testing"
)

func TestBest(t *testing.T) {

	// Test case 1: bestMatch is valid album
	var cardShelfRenderer = data.BEST_VALID_ALBUM
	bestMatch, err := best.Mapper(cardShelfRenderer)
	if err != nil {
		t.Errorf("test 1 failed: %v", err)
	}
	// Test case 1.1: bestMatch type is album
	if bestMatch.Type != "album" {
		t.Errorf(
			"test 1.1 failed: expected %v, got %v",
			"album", bestMatch.Type,
		)
	}

	// Test case 2: bestMatch is valid artist
	cardShelfRenderer = data.BEST_VALID_ARTIST
	bestMatch, err = best.Mapper(cardShelfRenderer)
	if err != nil {
		t.Errorf("test 2 failed: %v", err)
	}
	// Test case 2.1: bestMatch type is artist
	if bestMatch.Type != "artist" {

		t.Errorf(
			"test 2.1 failed: expected %v, got %v",
			"artist", bestMatch.Type,
		)
	}

	// Test case 3: bestMatch is valid song
	cardShelfRenderer = data.BEST_VALID_SONG
	bestMatch, err = best.Mapper(cardShelfRenderer)
	if err != nil {
		t.Errorf("test 3 failed: %v", err)
	}
	// Test case 3.1: bestMatch type is song
	if bestMatch.Type != "song" {
		t.Errorf(
			"test 3.1 failed: expected %v, got %v",
			"song", bestMatch.Type,
		)
	}

	// Test case 4: bestMatch is empty
	cardShelfRenderer = data.BEST_INVALID_EMPTY
	_, err = best.Mapper(cardShelfRenderer)
	if err != nil {
		t.Errorf("test 4 failed: %v", err)
	}

	// Test case 5: bestMatch is dont know type
	cardShelfRenderer = data.BEST_VALID_DONT_KNOW_TYPE
	_, err = best.Mapper(cardShelfRenderer)
	if err != nil {
		t.Errorf("test 5 failed: %v", err)
	}

}
