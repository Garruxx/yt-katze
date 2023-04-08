package tests

import (
	"katze/src/mappers/search/general/utils"
	"katze/src/mappers/search/general/utils/tests/data"
	"testing"
)

func TestValidateItemRenderer(t *testing.T) {

	// Test case 1: valid itemRenderer
	itemRenderer := data.VALIDATE_ITEM_RENDERER_VALID_A_FLEX_COLUMN
	err := utils.ValidateItemRenderer(itemRenderer, "MUSIC_PAGE_TYPE_ARTIST", 1)
	if err != nil {
		t.Errorf("test 1 failed, error: %v", err)
	}

	// Test case 2: valid two flexColumns
	itemRenderer = data.VALIDATE_ITEM_RENDERER_VALID_TWO_FLEX_COLUMN
	err = utils.ValidateItemRenderer(itemRenderer, "MUSIC_PAGE_TYPE_ARTIST", 2)
	if err != nil {
		t.Errorf("Test 2 failed, error: %v", err)
	}
	
	// Test case 3: invalid pageType
	itemRenderer = data.VALIDATE_ITEM_RENDERER_INVALID_PAGE_TYPE
	err = utils.ValidateItemRenderer(itemRenderer, "MUSIC_PAGE_TYPE_ARTIST", 2)
	if err == nil {
		t.Errorf("Test 3 failed, expected error, got %v", err)
	}

	// Test case 4: invalid no flexColumns
	itemRenderer = data.VALIDATE_ITEM_RENDERER_VALID_A_FLEX_COLUMN
	err = utils.ValidateItemRenderer(itemRenderer, "MUSIC_PAGE_TYPE_ARTIST", 2)
	if err == nil {	
		t.Errorf("Test 4 failed, expected error, got: %v", err)
	}

	// Test case 5: invalid no id
	itemRenderer = data.VALIDATE_ITEM_RENDERER_VALID_A_FLEX_COLUMN
	err = utils.ValidateItemRenderer(itemRenderer, "MUSIC_PAGE_TYPE_ARTIST", 2)
	if err == nil {
		t.Errorf("Test 5 failed, expected error, got: %v", err)
	}

	// Test case 6: invalid no length
	itemRenderer = data.VALIDATE_ITEM_RENDERER_VALID_A_FLEX_COLUMN
	err = utils.ValidateItemRenderer(itemRenderer, "MUSIC_PAGE_TYPE_ARTIST", 2)
	if err == nil {
		t.Errorf("Test 6 failed, expected error, got: %v", err)
	}
}
