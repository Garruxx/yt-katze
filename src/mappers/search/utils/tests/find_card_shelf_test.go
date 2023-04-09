package tests

import (
	"katze/src/mappers/search/utils"
	"katze/src/mappers/search/utils/tests/data"
	"testing"
)

func TestFindCardShelf(t *testing.T) {

	// Test case 1: tabContents is valid
	var tabContents = data.TAB_CONTENT_RENDERER
	var shelf, err = utils.FindCardShelf(tabContents, "test1")
	if err != nil {
		t.Errorf("Test case 1 failed: %v", err)
	}
	// Test case 1.1: shelf title is correct "test1"
	if shelf.Title.Runs[0].Text != "test1" {
		t.Errorf(
			"Test case 1.1 failed: expected %v, got %v",
			"test1", shelf.Title.Runs[0].Text,
		)
	}

	// Test case 2: tabContents is empty
	tabContents = data.TAB_CONTENT_RENDERER_EMPTY
	_, err = utils.FindCardShelf(tabContents, "test1")
	if err == nil {
		t.Errorf("Test case 2 failed: expected error, got nil")
	}

	// Test case 3: query is not found
	tabContents = data.TAB_CONTENT_RENDERER
	_, err = utils.FindCardShelf(tabContents, "test2")
	if err == nil {
		t.Errorf("Test case 3 failed: expected error, got nil")
	}
}
