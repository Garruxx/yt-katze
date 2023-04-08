package tests

import (
	"katze/src/mappers/utils/columns"
	"katze/src/mappers/utils/columns/tests/data"
	"testing"
)

func TestColumnMapper(t *testing.T) {

	// Test case 1: Test with valid data
	flexColumnData := data.FLEX_COLUMN_DATA
	flexColumn, err := columns.ColumnMapper(flexColumnData)
	if err != nil {
		t.Errorf("Test 1 failed, error: %s", err.Error())
	}
	// Test case 1.1: Test with valid data, text is equal to expected "Test"
	if text := flexColumn.Text; text != "Test" {
		t.Errorf("Test 1.1 failed, expected: %s, got: %s", "Test", text)
	}
	// Test case 1.2: Test with valid data,
	//browse id is equal to expected "TestID"
	if id := flexColumn.BrowseID; id != "TestID" {
		t.Errorf("Test 1.2 failed, expected: %s, got: %s", "TestID", id)
	}
	// Test case 1.3: Test with valid data,
	//pageType is equal to expected "PageTypeTest"
	if pageType := flexColumn.PageType; pageType != "PageTypeTest" {
		t.Errorf(
			"Test 1.3 failed, expected: %s, got: %s",
			"PageTypeTest", pageType,
		)
	}

	// Test case 2: Test with empty navigation endpoint
	flexColumnData = data.FLEX_COLUMN_DATA_EMPTY_NAVIGATION_ENDPOINT
	flexColumn, err = columns.ColumnMapper(flexColumnData)
	if err != nil {
		t.Errorf("Test 2 failed, error: %s", err)
	}
	// Test case 2.1: Test with empty navigation endpoint, text is equal to expected "Test"
	if text := flexColumn.Text; text != "Test" {
		t.Errorf("Test 2 failed, expected: %s, got: %s", "Test", text)
	}

	// Test case 3: Test with empty text
	flexColumnData = data.FLEX_COLUMN_DATA_EMPTY_TEXT
	_, err = columns.ColumnMapper(flexColumnData)
	if err == nil {
		t.Errorf("Test 3 failed, expected error, got nil")
	}

	// Test case 4: Test with empty browse id
	flexColumnData = data.FLEX_COLUMN_DATA_EMPTY_BROWSE_ID
	_, err = columns.ColumnMapper(flexColumnData)
	if err == nil {
		t.Errorf("Test 4 failed, expected error, got nil")
	}

	// Test case 5: Test with empty page type
	flexColumnData = data.FLEX_COLUMN_DATA_EMPTY_PAGE_TYPE
	_, err = columns.ColumnMapper(flexColumnData)
	if err == nil {
		t.Errorf("Test 5 failed, expected error, got nil")
	}

	// Test case 6: Test with empty browse id and page type
	flexColumnData = data.FLEX_COLUMN_DATA_EMPTY_BROWSE_ID_AND_PAGE_TYPE
	_, err = columns.ColumnMapper(flexColumnData)
	if err == nil {
		t.Errorf("Test 6 failed, expected error, got nil")
	}

	// Test case 7: Test with music id
	flexColumnData = data.FLEX_COLUMN_DATA_WITH_MUSIC_ID
	flexColumn, err = columns.ColumnMapper(flexColumnData)
	if err != nil {
		t.Errorf("Test 7 failed, error:%v", err)
	}
	// Test case 7.1: Test with music id, text is equal to expected "test"
	if id := flexColumn.WatchID; id != "test" {
		t.Errorf("Test 7.1 failed, expected test, got %s", id)
	}
	// Test case 7.2: Test with music id, pageType is equal to expected "test"
	if text := flexColumn.Text; text != "test" {
		t.Errorf("Test 7.2 failed, expected test, got %s", text)
	}
	// Test case 7.3: Test with music id, browseId is not empty
	if id := flexColumn.BrowseID; id != "" {
		t.Errorf("Test 7.3 failed, WatchId was not expected, got: %s", id)
	}
	// Test case 7.4: Test with music id, pageType is equal to expected "Music"
	if pageType := flexColumn.PageType; pageType != "Music" {
		t.Errorf(
			"Test 7.4 failed, pagetype music was expected, got: %s",
			pageType,
		)
	}

	// Test case 8: Test with empty data
	flexColumnData = data.FLEX_COLUMN_DATA_EMPTY
	_, err = columns.ColumnMapper(flexColumnData)
	if err == nil {
		t.Errorf("Test 8 failed, expected error, got nil")
	}
}
