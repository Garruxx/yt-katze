package tests

import (
	"katze/src/services/search"
	"testing"
)

func TestItemsList(t *testing.T) {

	// Test case 1: all is valid
	query := "The Beatles"
	param := "EgWKAQIIAWoIEAMQBBAJEAo%3D"
	visitorID := "CgswTi1jc3JLMnVuVSir06GhBg%3D%3D"

	result, err := search.ItemsList(query, param, visitorID)
	if err != nil {
		t.Errorf("Test case 1 failed: %v", err)
	}
	if len(result.Contents.TabbedSearchResultsRenderer.Tabs) == 0 {
		t.Errorf("Test case 1 failed: expected non-empty list of tabs")
	}

	// Test case 2: param is invalid
	param = "Tirilil"
	result, err = search.ItemsList(query, param, visitorID)
	if err != nil {
		t.Errorf("Test case 2 failed: %v", err)
	}

	// Test case 3: visitorID is nil
	var nilVisitorID string
	result, err = search.ItemsList(query, param, nilVisitorID)
	if err != nil {
		t.Errorf("Test case 3 failed: %v", err)
	}
	if len(result.Contents.TabbedSearchResultsRenderer.Tabs) == 0 {
		t.Errorf("Test case 3 failed: expected non-empty list of tabs")
	}

	// Test case 4: invalid visitorID
	visitorID = "Tirilil"
	result, err = search.ItemsList(query, param, visitorID)
	if err != nil {
		t.Errorf("Test case 4 failed: %v", err)
	}
	if len(result.Contents.TabbedSearchResultsRenderer.Tabs) == 0 {
		t.Errorf("Test case 4 failed: expected non-empty list of tabs")
	}

}
