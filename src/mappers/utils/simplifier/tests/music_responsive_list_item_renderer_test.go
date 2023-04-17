package tests

import (
	"katze/src/mappers/utils/simplifier"
	"katze/src/mappers/utils/simplifier/tests/data"
	"testing"
)

func TestMusicResponsiveListItemRenderer(t *testing.T) {

	// Test case 1 Valid data, type artist
	item := data.MUSIC_RESPONSIVE_LIST_ITEM_RENDERER_ARTIST
	result, err := simplifier.MusicResponsiveListItemRenderer(item)
	if err != nil {
		t.Errorf("failed test 1 error: %v", err)
	}

	// Test case 1.1: pageType is correct "MUSIC_PAGE_TYPE_ARTIST"
	if result.PageType != "MUSIC_PAGE_TYPE_ARTIST" {
		t.Errorf(
			"test 1.1 failed error: invalid pageType expected: %v, got: %v",
			"MUSIC_PAGE_TYPE_ARTIST", result.PageType,
		)
	}

	// Test case 1.2: url is correct "example.url"
	if url := result.Thumbnails[0].URL; url != "example.url" {
		t.Errorf(
			"test 1.2 failed error: invalid url expected: %v, got: %v",
			"example.url", url,
		)
	}

	// Test case 1.3: flexColumns length is correct 2
	if leng := len(result.FlexColumns); leng != 2 {
		t.Errorf(
			` failed test 1.3 error: invalid flexColumns length expected: %v,
			  got: %v`,
			2, leng,
		)
	}

	// Test case 1.4: result is no explicit
	if result.Explicit != false {
		t.Errorf(
			"failed test 1.4 error: invalid explicit expected: %v, got: %v",
			false, result.Explicit,
		)
	}

	// Test case 1.5: id is not empty
	if result.ID == "" {
		t.Errorf("error: invalid id expected %v, got: %v", "testID", result.ID)
	}

	// Test case 2 Valid data, type album
	item = data.MUSIC_RESPONSIVE_LIST_ITEM_RENDERER_ALBUM
	result, err = simplifier.MusicResponsiveListItemRenderer(item)
	if err != nil {
		t.Errorf("failed test 2 error: %v", err)
	}

	// Test case 2.1: pageType is correct "MUSIC_PAGE_TYPE_ALBUM"
	if pageType := result.PageType; pageType != "MUSIC_PAGE_TYPE_ALBUM" {
		t.Errorf(
			"failed test 2.1 error: invalid pageType expected: %v, got: %v",
			"MUSIC_PAGE_TYPE_ALBUM", pageType,
		)
	}

	// Test case 2.2: url is correct "example.url"
	if url := result.Thumbnails[0].URL; url != "example.url" {
		t.Errorf(
			"failed test 2.2 error: invalid url expected: %v, got: %v",
			"example.url", url,
		)
	}

	// Test case 2.3: flexColumns length is correct 2
	if leng := len(result.FlexColumns); leng != 2 {
		t.Errorf(
			`failed test 2.3 error: invalid flexColumns length expected: %v, 	
			 got: %v`,
			2, leng,
		)
	}

	// Test case 2.4: result is explicit (true)
	if result.Explicit != true {
		t.Errorf(
			"failed test 2.4 error: invalid explicit expected: %v, got: %v",
			true, result.Explicit,
		)
	}

	// Test case 2.5: id is not empty
	if result.ID == "" {
		t.Errorf(
			"failed test 2.5 error: invalid id expected: %v, got: %v",
			"example.id", result.ID,
		)
	}

	// Test case 3 Valid data, type song
	item = data.MUSIC_RESPONSIVE_LIST_ITEM_RENDERER_SONG
	result, err = simplifier.MusicResponsiveListItemRenderer(item)
	if err != nil {
		t.Errorf("failed test 1 error: %v", err)
	}
	// Test case 3.1: pageType is correct "MUSIC_TYPE"
	if result.PageType != "MUSIC_TYPE" {
		t.Errorf(
			"failed test 3.1 error: invalid pageType expected: %v, got: %v",
			"MUSIC_TYPE", result.PageType,
		)
	}
	// Test case 3.2: url is correct "example.url"
	if url := result.Thumbnails[0].URL; url != "example.url" {
		t.Errorf(
			"failed test 3.2 error: invalid url expected: %v, got: %v",
			"example.url", url,
		)
	}
	// Test case 3.3: flexColumns length is correct 2
	if leng := len(result.FlexColumns); leng != 2 {
		t.Errorf(
			`failed test 3.3 error: invalid flexColumns length expected: %v,    
			 got: %v`,
			2, leng,
		)
	}
	// Test case 3.4: result is explicit (true)
	if result.Explicit != true {
		t.Errorf(
			"failed test 3.4 error: invalid explicit expected: %v, got: %v",
			false, result.Explicit,
		)
	}
	// Test case 3.5: id is not empty
	if result.ID == "" {
		t.Errorf(
			"failed test 3.5 error: invalid id expected: %v, got: %v",
			"example.id", result.ID,
		)
	}
	// Test case 3.6: id is "example.id"
	if result.ID != "example.id" {
		t.Errorf(
			"failed test 3.6 error: invalid id expected: %v, got: %v",
			"example.test", result.ID,
		)
	}

	// Test case 4 Invalid data
	item = data.MUSIC_RESPONSIVE_LIST_ITEM_RENDERER_INVALID
	_, err = simplifier.MusicResponsiveListItemRenderer(item)
	if err == nil {
		t.Errorf("error: %v", err)
	}

	// Test case 5 empty data
	item = data.MUSIC_RESPONSIVE_LIST_ITEM_RENDERER_EMPTY
	_, err = simplifier.MusicResponsiveListItemRenderer(item)
	if err == nil {
		t.Errorf("test 5 failed error: %v", err)
	}
}
