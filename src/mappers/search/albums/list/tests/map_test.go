package tests

import (
	"katze/src/mappers/search/albums/list"
	"katze/src/models/external"
	"katze/src/utils"
	"testing"
)

func TestMap(t *testing.T) {

	// test case 1 - get valid albums list
	var itemsAlbums external.AlbumList
	err := utils.GetStructFromJson(
		"./data/json/albums_data_valid.json", &itemsAlbums,
	)
	if err != nil {
		t.Errorf("test 1.1 failed error getting data: %v", err)
	}
	// test case 1.2 - map valid albums list
	result, err := list.Map(itemsAlbums)
	if err != nil {
		t.Errorf("test 1.2 failed error mapping data: %v", err)
	}
	// test case 1.3 - validate length 22
	if leng := len(result.Albums); leng != 22 {
		t.Errorf("test 1.3 failed expected 22 got %v", leng)
	}
	// test case 1.4 - validate continuation id
	if id := result.ContinuationID; id != "continuationIDTest" {
		t.Errorf(
			"test 1.4 failed expected continuationIDTest got %v", id,
		)
	}

	// test case 2 - get invalid albums list error
	itemsAlbums = external.AlbumList{}
	err = utils.GetStructFromJson(
		"./data/json/albums_data_error.json", &itemsAlbums,
	)
	if err != nil {
		t.Errorf("test 2.1 failed error getting data: %v", err)
	}
	// test case 2.2 - map invalid albums list error
	_, err = list.Map(itemsAlbums)
	if err == nil {
		t.Errorf("test 2.2 failed expected error got nil")
	}

	// test case 3 - get invalid albums list no albums
	itemsAlbums = external.AlbumList{}
	err = utils.GetStructFromJson(
		"./data/json/albums_data_no_results.json", &itemsAlbums,
	)
	if err != nil {
		t.Errorf("test 3.1 failed error getting data: %v", err)
	}
	// test case 3.2 - map invalid albums list no albums
	_, err = list.Map(itemsAlbums)
	if err != nil {
		t.Errorf("test 3.2 failed expected no error got %v", err)
	}
}
