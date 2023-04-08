package tests

import (
	"katze/src/mappers/search/general/tests/data"
	"testing"
)

func TestGet(t *testing.T) {

	// test case 1 - get the data from the file path valid
	result, err := data.Get("../json/general_data_valid.json")
	if err != nil {
		t.Errorf("test 1 failed, error: %v", err)
	}
	// test case 1.1 - check if the data is empty
	if result.Contents == nil {
		t.Errorf("test 1.1 failed, error: contents is empty")
	}
	// test case 1.2 - check if the response context is nil
	if result.ResponseContext == nil {
		t.Errorf("test 1.2 failed, error: responseContext is empty")
	}

	// test case 2 - get the data from the file path valid but the data
	// contains an error
	result, err = data.Get("../json/general_data_error_invalid.json")
	if err != nil {
		t.Errorf("test 2 failed, error: %v", err)
	}
	// test case 2.1 - check if the data is empty
	if result.Contents != nil {
		t.Errorf("test 2.1 failed, error: contents is empty")
	}
	// test case 2.2 - check if has an error
	if result.Error == nil {
		t.Errorf("test 2.2 failed, error: error is empty")
	}

	// test case 3 - get the data from the file path invalid
	result, err = data.Get("../nil.json")
	if err == nil {
		t.Errorf("test 3 failed, error: %v", err)
	}
	// test case 3.1 - check if the data is empty
	if result.Contents != nil {
		t.Errorf("test 3.1 failed, error: contents is not empty")
	}

}
