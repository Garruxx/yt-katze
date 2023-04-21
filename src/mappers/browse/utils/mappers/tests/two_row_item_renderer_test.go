package tests

import (
	"katze/src/mappers/browse/utils/mappers"
	"katze/src/models/external"
	"katze/src/utils"
	"testing"
)

func TestTwoRowItemRenderer(t *testing.T) {

	// Test case 1: Get valid data
	twoRowItem := external.MusicCarouselShelfRendererContent{}
	err := utils.GetStructFromJson(
		"./data/json/two_row_items/renderer_valid.json", &twoRowItem,
	)
	if err != nil {
		t.Errorf("test 1 failed, error: %v", err)
	}
	// Test case 1.1: Check if the data has an error
	result, err := mappers.TwoRowItemRenderer(twoRowItem)
	if err != nil {
		t.Errorf("test 1.1 failed, error: %v", err)
	}
	// Test case 1.2: Check if the result has a title "Sakura"
	if result.Title != "Sakura" {
		t.Errorf(
			"test 1.2 failed, expected title Sakura, got %v",
			result.Title,
		)
	}
	// Test case 1.3: Check if the result has a year 2024
	if result.Year != 2024 {
		t.Errorf(
			"test 1.3 failed, expected year 2024, got %v",
			result.Year,
		)
	}
	// Test case 1.4: Check if the result has a browseID "browseIDTest"
	if result.ID != "browseIDTest" {
		t.Errorf(
			"test 1.4 failed, expected browseID browseIDTest, got %v",
			result.ID,
		)
	}
	// Test case 1.5: Check if the result has a params "paramsTest"
	if result.Params != "paramsTest" {
		t.Errorf(
			"test 1.5 failed, expected params paramsTest, got %v",
			result.Params,
		)
	}
	// Test case 1.6: Check if the result has a explicit true
	if result.Explicit != true {
		t.Errorf(
			"test 1.6 failed, expected explicit true, got %v",
			result.Explicit,
		)
	}

	// Test case 2: Get no title runs
	twoRowItem = external.MusicCarouselShelfRendererContent{}
	err = utils.GetStructFromJson(
		"./data/json/two_row_items/renderer_invalid_no_title_runs.json", &twoRowItem,
	)
	if err != nil {
		t.Errorf("test 2 failed, error: %v", err)
	}

	// Test case 2.1: Check if the data has an error
	_, err = mappers.TwoRowItemRenderer(twoRowItem)
	if err == nil {
		t.Errorf("test 2.1 failed, expected an error, got: %v", err)
	}

	// Test case 3: Get no subtitle runs
	twoRowItem = external.MusicCarouselShelfRendererContent{}
	err = utils.GetStructFromJson(
		"./data/json/two_row_items/renderer_invalid_no_subtitle_runs.json", &twoRowItem,
	)
	if err != nil {
		t.Errorf("test 3 failed, error: %v", err)
	}
	// Test case 3.1: Check if the data has an error
	_, err = mappers.TwoRowItemRenderer(twoRowItem)
	if err == nil {
		t.Errorf("test 3.1 failed, expected an error, got: %v", err)
	}

	// Test case 4: Get valid no explicit
	twoRowItem = external.MusicCarouselShelfRendererContent{}
	err = utils.GetStructFromJson(
		"./data/json/two_row_items/renderer_valid_no_explicit.json", &twoRowItem,
	)
	if err != nil {
		t.Errorf("test 4 failed, error: %v", err)
	}
	// Test case 4.1: Check if the data has an error
	result, err = mappers.TwoRowItemRenderer(twoRowItem)
	if err != nil {
		t.Errorf("test 4.1 failed, error: %v", err)
	}
	// Test case 4.2: Check if the result has a explicit false
	if result.Explicit != false {
		t.Errorf(
			"test 4.2 failed, expected explicit false, got %v",
			result.Explicit,
		)
	}
}
