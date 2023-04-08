package tests

import (
	"katze/src/services/search"
	"testing"
)

func TestGeneralResult(t *testing.T) {

	// Test case 1: search query and visitorID is valid
	query := "The Beatles"
	visitorID := "CgswTi1jc3JLMnVuVSir06GhBg%3D%3D"

	result, err := search.General(query, visitorID)
	if err != nil {
		t.Errorf("Test case 1 failed: %v", err)
	}
	if len(result.Contents.TabbedSearchResultsRenderer.Tabs) == 0 {
		t.Errorf("Test case 1 failed: expected non-empty list of tabs")
	}

	// Test case 2: visitorID is invalid
	visitorID = "Tirilil"
	result, err = search.General(query, visitorID)
	if err != nil {
		t.Errorf("Test case 2 failed: %v", err)
	}

	// Test case 3: visitorID is nil
	var nilVisitorID string
	result, err = search.General(query, nilVisitorID)
	if err != nil {
		t.Errorf("Test case 3 failed: %v", err)
	}
	if len(result.Contents.TabbedSearchResultsRenderer.Tabs) == 0 {
		t.Errorf("Test case 3 failed: expected non-empty list of tabs")
	}
}
