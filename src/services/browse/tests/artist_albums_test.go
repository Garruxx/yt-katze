package tests

import (
	"katze/src/services/browse"
	"testing"
)

// ArtistAlbums requieres a browseID, params and visitorID
// nothing is optional

func TestArtistAlbums(t *testing.T) {

	// Test case 1: all is valid
	browseID := "UCoIOOL7QKuBhQHVKL8y7BEQ"
	visitorID := "CgswTi1jc3JLMnVuVSir06GhBg%3D%3D"
	params :=
		"6gPoAUdxc0JXcHdCQ3BrQkNpUjVkRjl3WVdkbFgzTnVZWEJ6YUc5MFgyMTFjMmxqWDNCaFoyVmZjbVZuYVc5dVlXd1NIMjk2Wm1OQmVHWnhZMGx2WlRSRU9EaFFkVGR1YUZwMFZHZGplbVJqZUdjYVVBQUFaWE10TkRFNUFBRkRUd0FCUTA4QUFRQkdSVzExYzJsalgyUmxkR0ZwYkY5aGNuUnBjM1FBQVFGREFBQUJBQUVBQUFFQkFGVkRiMGxQVDB3M1VVdDFRbWhSU0ZaTFREaDVOMEpGVVFBQjh0cXpxZ29HUUFGSUFGQWk%3D"

	result, err := browse.ArtistAlbums(browseID, params, visitorID)
	if err != nil {
		t.Errorf("Test case 1 failed: %v", err)
	}
	if len(result.Contents.SingleColumnBrowseResultsRenderer.Tabs) == 0 {
		t.Errorf("Test case 1 failed: expected non-empty list of tabs")
	}

	// Test case 2: invalid browseID
	browseID = "Tirilil"
	result, err = browse.ArtistAlbums(browseID, params, visitorID)
	if err == nil {
		t.Errorf("Test case 2 failed expected error but got nil: %v", err)
	}

	// Test case 3: invalid visitorID
	visitorID = "Tirilil"
	browseID = "UCoIOOL7QKuBhQHVKL8y7BEQ"
	result, err = browse.ArtistAlbums(browseID, params, visitorID)
	if err != nil {
		t.Errorf("Test case 3 failed: %v", err)
	}

	// Test case 4: invalid params
	params = "Tirilil"
	browseID = "UCoIOOL7QKuBhQHVKL8y7BEQ"
	visitorID = "CgswTi1jc3JLMnVuVSir06GhBg%3D%3D"
	result, err = browse.ArtistAlbums(browseID, params, visitorID)
	if err != nil {
		t.Errorf("Test case 4 failed: %v", err)
	}
}
