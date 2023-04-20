package tests

import (
	"katze/src/services/browse"
	"testing"
)

// ArtistAlbums requieres tTracklisteID, params and visitorID
// nothing is optional

func TestArtTracklistums(t *testing.T) {

	// Test case 1: all is valid
	browseID := "UCoIOOL7QKuBhQHVKL8y7BEQ"
	visitorID := "CgtpQ1Q4NV9sUHhQNCiJjYKiBg%3D%3D"
	params :=
		"6gPjAUdxY0JXcGdCQ3BVQkNpUjVkRjl3WVdkbFgzTnVZWEJ6YUc5MFgyMTFjMmxqWDNCaFoyVmZjbVZuYVc5dVlXd1NIMWw0UVhSTFlYWlJWbmxyVkRSRU9EaFFkVGR1YUZwMWFrVXpZa2RsVW1jYVRBQUFaVzRBQVVOUEFBRkRUd0FCQUVaRmJYVnphV05mWkdWMFlXbHNYMkZ5ZEdsemRBQUJBVU1BQUFFQUFRQUFBUUVBVlVOdlNVOVBURGRSUzNWQ2FGRklWa3RNT0hrM1FrVlJBQUh5MnJPcUNnWkFBVWdBVUNJ"

	result, err := browse.ArtistTracklist(browseID, params, visitorID)
	if err != nil {
		t.Errorf("Test case 1 failed: %v", err)
	}
	if len(result.Contents.SingleColumnBrowseResultsRenderer.Tabs) == 0 {
		t.Errorf("Test case 1 failed: expected non-empty list of tabs")
	}
	if len(result.Contents.SingleColumnBrowseResultsRenderer.Tabs[0].TabRenderer.Content.SectionListRenderer.Contents) == 0 {
		t.Errorf("Test case 1 failed: expected non-empty list of contents")
	}

	// Test case 2: invalid browseID
	browseID = "Tirilil"
	result, err = browse.ArtistTracklist(browseID, params, visitorID)
	if err == nil {
		t.Errorf("Test case 2 failed expected error but got nil: %v", err)
	}

	// Test case 3: invalid visitorID
	visitorID = "Tirilil"
	browseID = "UCoIOOL7QKuBhQHVKL8y7BEQ"
	result, err = browse.ArtistTracklist(browseID, params, visitorID)
	if err != nil {
		t.Errorf("Test case 3 failed: %v", err)
	}

	// Test case 4: invalid params
	params = "Tirilil"
	browseID = "UCoIOOL7QKuBhQHVKL8y7BEQ"
	visitorID = "CgswTi1jc3JLMnVuVSir06GhBg%3D%3D"
	result, err = browse.ArtistTracklist(browseID, params, visitorID)
	if err != nil {
		t.Errorf("Test case 4 failed: %v", err)
	}
}
