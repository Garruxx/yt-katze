package tests

import (
	"katze/src/mappers/utils/simplifier"
	"katze/src/models/external"
	"katze/src/utils"
	"testing"
)

func TestItemsList(t *testing.T) {

	//  Test case 1 open file item list
	var itemList external.ItemsList
	err := utils.GetStructFromJson(
		"./data/json/items_list/valid.json", &itemList,
	)
	if err != nil {
		t.Errorf("test 1 failed: %v", err)
	}
	// Test case 1.1 simplify the data
	result, err := simplifier.ItemsList(itemList)
	if err != nil {
		t.Errorf("test 1.1 failed: %v", err)
	}
	// Test case 1.2 check if the content has the 22 items
	if leng := len(result.Contents); leng != 22 {
		t.Errorf("test 1.2 failed expected 22 items, got: %v", leng)
	}
	// Test case 1.3 check if the content has the nextPage "exampleNextPageTest"
	if result.NextPage != "exampleNextPageTest" {
		t.Errorf(
			"test 1.3 failed expected exampleNextPageTest, got: %v",
			result.NextPage,
		)
	}
	// Test case 1.4 check if the content has the visitorID "visitorIDTest"
	if result.VisitorID != "visitorIDTest" {
		t.Errorf(
			"test 1.4 failed expected visitorIDTest, got: %v",
			result.VisitorID,
		)
	}

	// Test case 2 open item list no tabs
	itemList = external.ItemsList{}
	err = utils.GetStructFromJson(
		"./data/json/items_list/no_tabs.json", &itemList,
	)
	if err != nil {
		t.Errorf("test 2 failed: %v", err)
	}
	// Test case 2.1 simplify the data
	result, err = simplifier.ItemsList(itemList)
	if err == nil {
		t.Errorf("test 2.1 failed expected error, got: %v", err)
	}

	// Test case 3 open item list no tab contents
	itemList = external.ItemsList{}
	err = utils.GetStructFromJson(
		"./data/json/items_list/no_tab_contents.json", &itemList,
	)
	if err != nil {
		t.Errorf("test 3 failed: %v", err)
	}
	// Test case 3.1 verify the error
	result, err = simplifier.ItemsList(itemList)
	if err == nil {
		t.Errorf("test 3.1 failed expected error, got: %v", err)
	}

	// Test case 4 open item list no shelf contents
	itemList = external.ItemsList{}
	err = utils.GetStructFromJson(
		"./data/json/items_list/no_shelf_contents.json", &itemList,
	)
	if err != nil {
		t.Errorf("test 4 failed: %v", err)
	}
	// Test case 4.1 verify no error
	result, err = simplifier.ItemsList(itemList)
	if err != nil {
		t.Errorf("test 4.1 failed expected no error, got: %v", err)
	}
	// Test case 4.2 verify the length of the contents
	if leng := len(result.Contents); leng != 0 {
		t.Errorf("test 4.2 failed expected 0 items, got: %v", leng)
	}
	// Test case 4.3 verify the nextPage
	if result.NextPage != "" {
		t.Errorf("test 4.3 failed expected empty nextPage, got: %v", result.NextPage)
	}
	// Test case 4.4 verify the visitorID
	if result.VisitorID != "" {
		t.Errorf("test 4.4 failed expected empty visitorID, got: %v", result.VisitorID)
	}

	// Test case 5 open item list error
	itemList = external.ItemsList{}
	err = utils.GetStructFromJson(
		"./data/json/items_list/error.json", &itemList,
	)
	if err != nil {
		t.Errorf("test 5 failed: %v", err)
	}

	// test case 6 open item list no response context
	itemList = external.ItemsList{}
	err = utils.GetStructFromJson(
		"./data/json/items_list/no_response_context.json", &itemList,
	)
	if err != nil {
		t.Errorf("test 6 failed: %v", err)
	}
	// Test case 6.1 verify the error
	result, err = simplifier.ItemsList(itemList)
	if err == nil {
		t.Errorf("test 6.1 failed expected error, got: %v", err)
	}

	// Test case 7 open item list no continuation id
	itemList = external.ItemsList{}
	err = utils.GetStructFromJson(
		"./data/json/items_list/no_continuation_id.json", &itemList,
	)
	if err != nil {
		t.Errorf("test 7 failed: %v", err)
	}
	// Test case 7.1 verify the error
	result, err = simplifier.ItemsList(itemList)
	if err != nil {
		t.Errorf("test 7.1 failed expected no error, got: %v", err)
	}
}
