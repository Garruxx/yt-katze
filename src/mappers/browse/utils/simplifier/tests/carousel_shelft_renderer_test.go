package tests

import (
	"katze/src/mappers/browse/utils/simplifier"
	"katze/src/models/external"
	"katze/src/utils"
	"testing"
)

func TestCarouselShelfRenderer(t *testing.T) {

	// test case 1 - get the data valid
	carousel := external.MusicCarouselShelfRenderer{}
	err := utils.GetStructFromJson(
		"./data/json/carousel_shelf/valid.json", &carousel,
	)
	if err != nil {
		t.Errorf("test 1 failed, error: %v", err)
	}
	// test case 1.1 - check if the data has an error
	result, err := simplifier.CarouselShelfRenderer(carousel)
	if err != nil {
		t.Errorf("test 1.1 failed, error: %v", err)
	}
	// test case 1.2 - check if the result has a title "Sencillos"
	if result.Title != "Sencillos" {
		t.Errorf(
			"test 1.2 failed, expected title Sencillos, got %v",
			result.Title,
		)
	}
	// test case 1.3 - check if the result has a continuationID
	if result.BrowseID != "browseIDTest" {
		t.Errorf(
			"test 1.3 failed, expected browseID browseIDTest, got %v",
			result.BrowseID,
		)
	}
	// test case 1.4 - check if the result has a params
	if result.Params != "paramsTest" {
		t.Errorf(
			"test 1.4 failed, expected params paramsTest, got %v",
			result.Params,
		)
	}
	// test case 1.5 - check if the content has 4 items
	if len(result.Contents) != 4 {
		t.Errorf(
			"test 1.5 failed, expected 4 items, got %v",
			len(result.Contents),
		)
	}

	// test case 2 - get the data no contents
	carousel = external.MusicCarouselShelfRenderer{}
	err = utils.GetStructFromJson(
		"./data/json/carousel_shelf/no_contents.json", &carousel,
	)
	if err != nil {
		t.Errorf("test 2 failed, error: %v", err)
	}
	// test case 2.1 - check if the data has an error
	result, err = simplifier.CarouselShelfRenderer(carousel)
	if err != nil {
		t.Errorf("test 2.1 failed, error: %v", err)
	}
	// test case 2.2 - check if the result has a title "Sencillos"
	if result.Title != "Sencillos" {
		t.Errorf(
			"test 2.2 failed, expected title Sencillos, got %v",
			result.Title,
		)
	}
	// test case 2.3 - check if the contents length is 0
	if len(result.Contents) != 0 {
		t.Errorf(
			"test 2.3 failed, expected 0 items, got %v",
			len(result.Contents),
		)
	}

	// test case 3 - get the data invalid 
	carousel = external.MusicCarouselShelfRenderer{}
	err = utils.GetStructFromJson(
		"./data/json/carousel_shelf/invalid.json", &carousel,
	)
	if err != nil {
		t.Errorf("test 3 failed, error: %v", err)
	}
	// test case 3.1 - check if the data has an error
	_, err = simplifier.CarouselShelfRenderer(carousel)
	if err == nil {
		t.Errorf("test 3.1 failed, error: %v", err)
	}
}
