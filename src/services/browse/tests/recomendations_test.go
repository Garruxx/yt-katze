package tests

import (
	"katze/src/services/browse"
	"testing"
)

func TestRecomendations(t *testing.T) {

	browseID := "MPTRt_Ssvc78X4qDD-1"
	visitorID := "CgswTi1jc3JLMnVuVSir06GhBg%3D%3D"
	result, err := browse.Recomendations(browseID, visitorID)
	// Test case 1 - valid browseID and visitorID
	if err != nil {
		t.Errorf("Test case 1 failed: %v", err)
	}
	// Test case 1.1 - should have a list of contents
	if len(result.Contents.SectionListRenderer.Contents) == 0 {
		t.Errorf("Test case 1 failed: expected non-empty list of tabs")
	}

}
