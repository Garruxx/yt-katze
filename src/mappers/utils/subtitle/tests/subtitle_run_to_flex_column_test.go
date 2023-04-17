package tests

import (
	"katze/src/mappers/utils/subtitle"
	"katze/src/mappers/utils/subtitle/tests/data"
	"testing"
)

func TestRunToFlexColumn(t *testing.T) {

	// Test case 1: Test with valid data
	flexColumnData := data.TO_COLUMN_RUN_VALID
	expectedText := "testText"
	subtitleRun, err := subtitle.RunToFlexColumn(flexColumnData)
	if err != nil {
		t.Errorf("Test 1 failed, error: %s", err)
	}
	// Test case 1.1: Test with valid text "testText"
	if subtitleRun.Text != "testText" {
		t.Errorf(
			"Test 1.1 failed, expected: %s, got: %s",
			"testText", subtitleRun.Text,
		)
	}
	// Test case 1.2: Test with valid browse ID "testID"
	if *subtitleRun.BrowseID != "testID" {
		t.Errorf(
			"Test 1.2 failed, expected: %s, got: %s",
			"testID", *subtitleRun.BrowseID,
		)
	}
	// Test case 1.3: Test with valid page type "testPageType"
	if *subtitleRun.PageType != "testPageType" {
		t.Errorf(
			"Test 1.3 failed, expected: %s, got: %s",
			"testPageType", *subtitleRun.PageType,
		)
	}

	// Test case 2: Test with empty data
	flexColumnData = data.TO_COLUMN_RUN_INVALID_EMPTY
	_, err = subtitle.RunToFlexColumn(flexColumnData)
	if err == nil {
		t.Errorf("Test 2 failed, expected error, got nil")
	}

	// Test case 3: Test with empty text
	flexColumnData = data.TO_COLUMN_RUN_INVALID_EMPTY_TEXT
	_, err = subtitle.RunToFlexColumn(flexColumnData)
	if err == nil {
		t.Errorf("Test 3 failed, expected error, got nil")
	}

	// Test case 4: Test with empty browse ID
	flexColumnData = data.TO_COLUMN_RUN_INVALID_EMPTY_BROWSE_ID
	_, err = subtitle.RunToFlexColumn(flexColumnData)
	if err == nil {
		t.Errorf("Test 4 failed, expected error, got nil")
	}

	// Test case 5: Test with empty page type
	flexColumnData = data.TO_COLUMN_RUN_INVALID_EMPTY_PAGE_TYPE
	_, err = subtitle.RunToFlexColumn(flexColumnData)
	if err == nil {
		t.Errorf("Test 5 failed, expected error, got nil")
	}

	// Test case 6: Test with just text
	flexColumnData = data.TO_COLUMN_RUN_VALID_JUST_TEXT
	expectedText = "testText"
	subtitleRun, err = subtitle.RunToFlexColumn(flexColumnData)
	if err != nil {
		t.Errorf("Test 6 failed, error: %s", err)
	}
	if subtitleRun.Text != expectedText {
		t.Errorf(
			"Test 6 failed, expected: %s, got: %s",
			expectedText, subtitleRun.Text,
		)
	}
}
