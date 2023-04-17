package tests

import (
	"katze/src/mappers/browse/utils"
	"katze/src/mappers/browse/utils/tests/data"
	"testing"
)

func TestFindCarouselShelf(t *testing.T) {

	// test case 1 valid data
	carousels := data.FIND_CAROUSEL_SHELF_VALID
	result, err := utils.FindCarouselShelf(carousels, "titleTest")
	if err != nil {
		t.Errorf("test 1 failed, error: %v", err)
	}
	if result.Title != "titleTest" {
		t.Errorf(
			"test 1 failed, expected title titleTest, got %v",
			result.Title,
		)	
	}

	// test case 2 invalid data no contents
	carousels = data.FIND_CAROUSEL_SHELF_INVALID_EMPTY
	_, err = utils.FindCarouselShelf(carousels, "titleTest")
	if err == nil {
		t.Errorf("test 2 failed, expected error, got: %v", err)
	}

	// test case 3 invalid title 
	carousels = data.FIND_CAROUSEL_SHELF_VALID
	_, err = utils.FindCarouselShelf(carousels, "titleTestInvalid")
	if err == nil {
		t.Errorf("test 3 failed, error: %v", err)
	}
}
