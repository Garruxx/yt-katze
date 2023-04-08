package tests

import (
	"katze/src/mappers/search/utils"
	"katze/src/mappers/search/utils/tests/data"
	"katze/src/models/external"
	"testing"
)

func TestGetBestResult(t *testing.T) {

	// Test case 1: tabContents is valid
	var tabContents = data.GET_BEST_RESULT
	var bestResult, err = utils.GetBestResult(tabContents)
	if err != nil {
		t.Errorf("Test case 1 failed: %v", err)
	}
	// Test case 1.1: bestResult title expected "test"
	if title := bestResult.Title.Runs[0].Text; title != "test" {
		t.Errorf(
			"Test case 1.1 failed: expected %v, got %v",
			"test", title,
		)
	}

	// Test case 2: tabContents is empty
	tabContents = []external.MagentaContent{}
	_, err = utils.GetBestResult(tabContents)
	if err == nil {
		t.Errorf("Test case 2 failed: expected error, got nil")
	}

	// Test case 3: there is no best outcome
	tabContents = data.GET_BEST_RESULT_NO_BEST_OUTCOME
	_, err = utils.GetBestResult(tabContents)
	if err == nil {
		t.Errorf("Test case 3 failed: expected error, got nil")
	}
}
