package tests

import (
	"katze/src/mappers/search/general/result/best/utils"
	"katze/src/mappers/search/general/result/best/utils/tests/data"
	"testing"
)

func TestShelfTypeValidator(t *testing.T) {

	// Test case 1: cardShelfRenderer is valid
	cardShelfRenderer := data.SHELF_VALIDATOR_VALID
	shelfType, err := utils.ShelfTypeValidator(cardShelfRenderer)
	if err != nil {
		t.Errorf("test 1 failed: %v", err)
	}
	// Test case 1.1: bestMatch type is TEST_PAGE_TYPE
	if shelfType != "TEST_PAGE_TYPE" {
		t.Errorf(
			"test 1.1 failed: expected %v, got %v",
			"TEST_PAGE_TYPE", shelfType,
		)
	}

	// Test case 2: cardShelfRenderer is valid music
	cardShelfRenderer = data.SHELF_VALIDATOR_VALID_MUSIC
	shelfType, err = utils.ShelfTypeValidator(cardShelfRenderer)
	if err != nil {
		t.Errorf("test 2 failed: %v", err)
	}
	// Test case 2.1: bestMatch type is TEST_PAGE_TYPE_MUSIC
	if shelfType != "TEST_PAGE_TYPE_MUSIC" {
		t.Errorf(
			"test 2.1 failed: expected %v, got %v",
			"TEST_PAGE_TYPE_MUSIC", shelfType,
		)
	}

	// Test case 3: cardShelfRenderer is empty
	cardShelfRenderer = data.SHELF_VALIDATOR_INVALID_EMPTY
	_, err = utils.ShelfTypeValidator(cardShelfRenderer)
	if err == nil {
		t.Errorf("test 3 failed: %v", err)
	}

	// Test case 4: cardShelfRenderer invalid browseEndpoint is empty
	cardShelfRenderer = data.SHELF_VALIDATOR_INVALID_NO_BROWSE_ENDPOINT
	_, err = utils.ShelfTypeValidator(cardShelfRenderer)
	if err == nil {
		t.Errorf("test 4 failed: %v", err)
	}

	// Test case 5: cardShelfRenderer invalid pageType is empty
	cardShelfRenderer = data.SHELF_VALIDATOR_INVALID_EMPTY
	_, err = utils.ShelfTypeValidator(cardShelfRenderer)
	if err == nil {
		t.Errorf("test 5 failed: %v", err)
	}

	// Test case 6: cardShelfRenderer invalid no browse id
	cardShelfRenderer = data.SHELF_VALIDATOR_INVALID_NO_BROWSE_ID
	_, err = utils.ShelfTypeValidator(cardShelfRenderer)
	if err == nil {
		t.Errorf("test 7 failed: %v", err)
	}

	// Test case 7: cardShelfRenderer invalid
	cardShelfRenderer = data.SHELF_VALIDATOR_INVALID
	_, err = utils.ShelfTypeValidator(cardShelfRenderer)
	if err == nil {
		t.Errorf("test 7 failed: %v", err)
	}
}
