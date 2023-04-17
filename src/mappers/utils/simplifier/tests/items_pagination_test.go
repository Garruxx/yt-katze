package tests

import (
	"katze/src/mappers/utils/simplifier"
	"katze/src/models/external"
	"katze/src/utils"
	"testing"
)

func TestItemsPagination(t *testing.T) {

	// test case 1 - get the valid data struct
	var pagination external.ItemsPagination
	err := utils.GetStructFromJson(
		"./data/json/items_pagination/valid.json", &pagination,
	)
	if err != nil {
		t.Errorf("test 1 failed, error: %v", err)
	}
	// test case 1.1 - check if the data has an error
	result, err := simplifier.ItemsPagination(pagination)
	if err != nil {
		t.Errorf("test 1.1 failed, error: %v", err)
	}
	// test case 1.2 - check if the data has 22 items
	if leng := len(result.Contents); leng != 22 {
		t.Errorf("test 1.2 failed, expected 22 items, got: %v", leng)
	}
	// test case 1.3 - check if the data has the visitorID "visitorIDTest"
	if result.VisitorID != "visitorIDTest" {
		t.Errorf(
			"test 1.3 failed, expected visitorIDTest, got: %v",
			result.VisitorID,
		)
	}
	// test case 1.4 - check if the data has the nextPage "exampleNextPageTest"
	if result.NextPage != "exampleNextPageTest" {
		t.Errorf(
			"test 1.4 failed, expected exampleNextPageTest, got: %v",
			result.NextPage,
		)
	}

	// test case 2 - get the data struct with no continuation contents
	pagination = external.ItemsPagination{}
	err = utils.GetStructFromJson(
		"./data/json/items_pagination/no_continuation_contents.json", &pagination,
	)
	if err != nil {
		t.Errorf("test 2 failed, error: %v", err)
	}
	// test case 2.1 - check if the data has an error
	result, err = simplifier.ItemsPagination(pagination)
	if err != nil {
		t.Errorf("test 2.1 failed, error: %v", err)
	}
	// test case 2.2 - check if the data has 0 items
	if leng := len(result.Contents); leng != 0 {
		t.Errorf("test 2.2 failed, expected 0 items, got: %v", leng)
	}

	// test case 3 - get the data struct with no response context
	pagination = external.ItemsPagination{}
	err = utils.GetStructFromJson(
		"./data/json/items_pagination/no_response_context.json", &pagination,
	)
	if err != nil {
		t.Errorf("test 3 failed, error: %v", err)
	}
	// test case 3.1 - check if the data has an error
	result, err = simplifier.ItemsPagination(pagination)
	if err == nil {
		t.Errorf("test 3.1 failed, expected error, got: %v", err)
	}

	// test case 4 - get the data error
	pagination = external.ItemsPagination{}
	err = utils.GetStructFromJson(
		"./data/json/items_pagination/error.json", &pagination,
	)
	if err != nil {
		t.Errorf("test 4 failed, error: %v", err)
	}
	// test case 4.1 - check if the data has an error
	result, err = simplifier.ItemsPagination(pagination)
	if err == nil {
		t.Errorf("test 4.1 failed, expected error, got: %v", err)
	}

	// test case 5 - get the data struct with no continuation id
	pagination = external.ItemsPagination{}
	err = utils.GetStructFromJson(
		"./data/json/items_pagination/no_continuation_id.json", &pagination,
	)
	if err != nil {
		t.Errorf("test 5 failed, error: %v", err)
	}
	// test case 5.1 - check if the data has an error
	result, err = simplifier.ItemsPagination(pagination)
	if err != nil {
		t.Errorf("test 5.1 failed, error: %v", err)
	}

}
