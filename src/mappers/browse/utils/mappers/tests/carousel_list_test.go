package tests

import (
	"katze/src/mappers/browse/utils/mappers"
	"katze/src/mappers/browse/utils/mappers/tests/data"
	"testing"
)

func TestCarouselShelf(t *testing.T) {
	// Test case 1: Test with valid data
	carouselShelf := data.CAROUSEL_LIST_VALID
	result, err := mappers.CarouselList(carouselShelf)
	if err != nil {
		t.Errorf("test 1 failed, error: %v", err)
	}
	// Test case 1.1: Check if the result has a two items
	if leng := len(result.Items); leng != 2 {
		t.Errorf("test 1.1 failed, expected 2 items, got %v", leng)
	}

	// Test case 2: Test with no contents
	carouselShelf = data.CAROUSEL_LIST_VALID_EMPTY
	_, err = mappers.CarouselList(carouselShelf)
	if err != nil {
		t.Errorf("test 2 failed, error: %v", err)
	}

}
