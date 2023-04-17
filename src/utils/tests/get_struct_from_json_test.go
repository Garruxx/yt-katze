package tests

import (
	"katze/src/models/external"
	"katze/src/utils"
	"testing"
)

func TestConvertJSONToStruct(t *testing.T) {

	//  Test case 1 open file
	var expected external.General
	err := utils.GetStructFromJson("./json/file.json", &expected)
	if err != nil {
		t.Errorf("test 1 failed: %v", err)
	}
	// Test case 1.1 check if the struct has the correct values visitorData
	if visitorID := expected.ResponseContext.VisitorData; visitorID != "test1" {
		t.Errorf("error expected test1, got'%s'", visitorID)
	}
	// Test case 1.2 check if the struct has the correct values seconds
	if seconds := expected.ResponseContext.MaxAgeSeconds; seconds != 120 {
		t.Errorf("error expected 120, got: %d", seconds)
	}

	// Test case 2 open file invalid path
	err = utils.GetStructFromJson("./json/file2noexist.json", &expected)
	if err == nil {
		t.Errorf("test 2 failed expected error, got: %v", err)
	}
}
