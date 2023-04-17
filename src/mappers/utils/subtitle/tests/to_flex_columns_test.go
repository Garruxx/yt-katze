package tests

import (
	"katze/src/mappers/utils/subtitle"
	"katze/src/mappers/utils/subtitle/tests/data"
	"testing"
)

func TestToFlexColumns(t *testing.T) {

	// Test case 1: Test with valid data
	var Subtitle = data.TO_FLEX_COLUMN_SUBTITLE_VALID
	flexColumn, err := subtitle.ToFlexColumns(Subtitle)
	if err != nil {
		t.Errorf("Test 1 failed, error: %s", err)
	}
	// Test case 1.1: Test with valid text flex column length 3
	if len(flexColumn) != 3 {
		t.Errorf(
			"Test 1.1 failed, expected: %d, got: %d",
			3, len(flexColumn),
		)
	}
	// Test case 1.2: Test with valid text "testText"
	if flexColumn[0].Text != "testText" {
		t.Errorf(
			"Test 1.2 failed, expected: %s, got: %s",
			"testText", flexColumn[0].Text,
		)
	}
	// Test case 1.3: Test with valid text "testText1"
	if flexColumn[1].Text != "testText1" {
		t.Errorf(
			"Test 1.3 failed, expected: %s, got: %s",
			"testText1", flexColumn[1].Text,
		)
	}
	// Test case 1.4: Test with valid browseID "testID"
	if *flexColumn[1].BrowseID != "testID" {
		t.Errorf(
			"Test 1.4 failed, expected: %s, got: %s",
			"testID", *flexColumn[0].BrowseID,
		)
	}
	// Test case 1.5: Test with valid pageType "testPageType"
	if *flexColumn[1].PageType != "testPageType" {
		t.Errorf(
			"Test 1.5 failed, expected: %s, got: %s",
			"testPageType", *flexColumn[0].PageType,
		)
	}
	// Test case 1.6: Test with valid text "testText2"
	if flexColumn[2].Text != "testText2" {
		t.Errorf(
			"Test 1 failed, expected: %s, got: %s",
			"testText2", flexColumn[2].Text,
		)
	}

	// Test case 2: Test with just text
	Subtitle = data.TO_FLEX_COLUMN_SUBTITLE_VALID_JUST_TEXT
	flexColumn, err = subtitle.ToFlexColumns(Subtitle)
	if err != nil {
		t.Errorf("Test 2 failed, error: %s", err)
	}
	// Test case 2.1: Test with valid text flex column length 1
	if len(flexColumn) != 1 {
		t.Errorf(
			"Test 2.1 failed, expected: %d, got: %d",
			1, len(flexColumn),
		)
	}
	// Test case 2.2: Test with valid text "testText"
	if flexColumn[0].Text != "testText" {
		t.Errorf(
			"Test 2 failed, expected: %s, got: %s",
			"testText", flexColumn[0].Text,
		)
	}

	// Test case 3: Test with empty data
	Subtitle = data.TO_FLEX_COLUMN_SUBTITLE_EMPTY
	_, err = subtitle.ToFlexColumns(Subtitle)
	if err == nil {
		t.Errorf("Test 3 failed, expected error, got nil")
	}

	// Test case 4: Test with empty text
	Subtitle = data.TO_FLEX_COLUMN_SUBTITLE_INVALID_EMPTY_TEXT
	_, err = subtitle.ToFlexColumns(Subtitle)
	if err == nil {
		t.Errorf("Test 4 failed, expected error, got nil")
	}

	// Test case 5: Test with an empty text
	Subtitle = data.TO_FLEX_COLUMN_SUBTITLE_INVALID_AN_EMPTY_TEXT
	_, err = subtitle.ToFlexColumns(Subtitle)
	if err == nil {
		t.Errorf("Test 5 failed, expected error, got nil")
	}

	// Test case 6: Test with empty browseID
	Subtitle = data.TO_FLEX_COLUMN_SUBTITLE_INVALID_EMPTY_BROWSE_ID
	_, err = subtitle.ToFlexColumns(Subtitle)
	if err == nil {
		t.Errorf("Test 6 failed, expected error, got nil")
	}

	// Test case 7: Test with an empty pageType
	Subtitle = data.TO_FLEX_COLUMN_SUBTITLE_INVALID_EMPTY_PAGE_TYPE
	_, err = subtitle.ToFlexColumns(Subtitle)
	if err == nil {
		t.Errorf("Test 7 failed, expected error, got nil")
	}
}
