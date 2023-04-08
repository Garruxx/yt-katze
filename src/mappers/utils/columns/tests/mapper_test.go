package tests

import (
	"katze/src/mappers/utils/columns"
	"katze/src/mappers/utils/columns/tests/data"
	"testing"
)

func TestMapper(t *testing.T) {

	// Test case 1: valid data
	flexColumn := data.MAPPER_FLEX_COLUMNS_VALID
	result, err := columns.Mapper(flexColumn)
	if err != nil {
		t.Errorf("Test 1 failed: %v", err)
	}
	// Test case 1.1: result text is equal to expected "test"
	if text := result[0].Items[0].Text; text != "test" {
		t.Errorf("Test 1.1 failed: expected %s, got %v", text, err)
	}
	// Test case 1.2: result browseID is equal to expected "testBrowseID"
	if id := result[0].Items[0].BrowseID; id != "testBrowseID" {
		t.Errorf("Test 1.2 failed: expected testBrowseID, got %v", id)
	}
	// Test case 1.3: result pageType is equal to expected "testPageType"
	if pageType := result[0].Items[0].PageType; pageType != "testPageType" {
		t.Errorf("Test 1.3 failed: expected testPageType, got %v", pageType)
	}
	// Test case 1.4: result text is equal to expected "test1"
	if text := result[1].Items[0].Text; text != "test1" {
		t.Errorf("Test 1.4 failed: expected test1, got %v", text)
	}
	// Test case 1.5: result length is equal to expected 2
	if leng := len(result); leng != 2 {
		t.Errorf("Test 1.5 failed: expected 2, got %v", leng)
	}
	// Test case 1.6: result length is equal to expected 1
	if leng := len(result[0].Items); leng != 1 {
		t.Errorf("Test 1 failed: expected 1, got %v", leng)
	}

	// Test case 2: invalid data empty an text
	flexColumn = data.MAPPER_FLEX_COLUMNS_WITH_AN_EMPTY_TEXT
	result, err = columns.Mapper(flexColumn)
	if err == nil {
		t.Error("Test 2 failed: expected error, got nil")
	}
	// Test case 2.1: result length is diferent to 0
	if leng := len(result); leng != 0 {
		t.Errorf("Test 2.1 failed: expected 0, got %v", leng)
	}

	// Test case 3: invalid data empty all text
	flexColumn = data.MAPPER_FLEX_COLUMNS_WITH_EMPTY_TEXT
	result, err = columns.Mapper(flexColumn)
	if err == nil {
		t.Error("Test 3 failed: expected error, got nil")
	}
	// Test case 3.1: result length is diferent to 0
	if leng := len(result); leng != 0 {
		t.Errorf("Test 3 failed: expected 0, got %v", leng)
	}

	// Test case 4: invalid data empty navigation endpoint
	flexColumn = data.MAPPER_FLEX_COLUMNS_WITHOUT_NAVIGATION_ENDPOINT
	_, err = columns.Mapper(flexColumn)
	if err != nil {
		t.Errorf("Test 4 failed: %v", err)
	}

	// Test case 5: invalid data empty
	flexColumn = data.MAPPER_FLEX_COLUMNS_EMPTY
	result, err = columns.Mapper(flexColumn)
	if err == nil {
		t.Errorf("Test 5 failed: %v", err)
	}
	// Test case 5.1: result length is diferent to 0
	if leng := len(result); leng != 0 {
		t.Errorf("Test 5.1 failed: expected 0, got %v", leng)
	}

	// Test case 6: valid only one item
	flexColumn = data.MAPPER_FLEX_COLUMNS_ONLY_ONE_COLUMN
	result, err = columns.Mapper(flexColumn)
	if err != nil {
		t.Errorf("Test 6 failed: %v", err)
	}
	// Test case 6.1: result text is equal to expected "test"
	if text := result[0].Items[0].Text; text != "test" {
		t.Errorf("Test 6.1 failed: expected test, got %v", text)
	}

	// Test case 7: valid three columns
	flexColumn = data.MAPPER_FLEX_COLUMNS_THREE_COLUMNS
	result, err = columns.Mapper(flexColumn)
	if err != nil {
		t.Errorf("Test 7 failed: %v", err)
	}

	// Test case 7.1: result length is equal to expected 3
	if leng := len(result); leng != 3 {
		t.Errorf("Test 7.1 failed: expected 3, got %v", leng)
	}
}
