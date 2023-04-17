package tests

import (
	"katze/src/mappers/browse/utils/simplifier"
	"katze/src/models/external"
	"katze/src/utils"
	"testing"
)

func TestMusicDetailHeaderRenderer(t *testing.T) {

	// Test case 1 valid data
	headerData := external.MusicDetailHeaderRenderer{}
	err := utils.GetStructFromJson(
		"./data/json/detail_header/valid.json",
		&headerData,
	)
	if err != nil {
		t.Errorf("test 1 failed error: %s", err)
	}
	// test case 1.1 check if the data has an error
	result, err := simplifier.MusicDetailHeaderRenderer(headerData)
	if err != nil {
		t.Errorf("test 1.1 failed error: %s", err)
	}
	// test case 1.2 check has a title "Mi berrinche"
	if result.Title != "Mi berrinche" {
		t.Errorf(
			"test 1.2 failed expected a title, got %v",
			result.Title,
		)
	}
	// test case 1.3 check has a duration "25 minutos"
	if result.Duration != "25 minutos" {
		t.Errorf(
			"test 1.3 failed expected a duration, got %v",
			result.Duration,
		)
	}
	// test case 1.4 check has a year 2024
	if result.Year != 2024 {
		t.Errorf(
			"test 1.4 failed expected a year, got %v",
			result.Year,
		)
	}
	// test case 1.5 check has a 4 thumbnails
	if len(result.Thumbnails) != 4 {
		t.Errorf(
			"test 1.5 failed expected 4 thumbnails, got %v",
			len(result.Thumbnails),
		)
	}
	// test case 1.6 check has a 1 artist
	if len(result.Artists) != 1 {
		t.Errorf(
			"test 1.6 failed expected 1 artist, got %v",
			len(result.Artists),
		)
	}

	// Test case 2 invalid data no title
	headerData = external.MusicDetailHeaderRenderer{}
	err = utils.GetStructFromJson(
		"./data/json/detail_header/invalid_no_title.json",
		&headerData,
	)
	if err != nil {
		t.Errorf("test 2 failed error: %s", err)
	}
	// test case 2.1 check if the data has an error
	_, err = simplifier.MusicDetailHeaderRenderer(headerData)
	if err == nil {
		t.Errorf("test 2.1 failed expected an error, got nil")
	}

	// Test case 3 valid no thumbnails
	headerData = external.MusicDetailHeaderRenderer{}
	err = utils.GetStructFromJson(
		"./data/json/detail_header/valid_no_thumbnails.json",
		&headerData,
	)
	if err != nil {
		t.Errorf("test 3 failed error: %s", err)
	}
	// test case 3.1 check if the data has an error
	result, err = simplifier.MusicDetailHeaderRenderer(headerData)
	if err != nil {
		t.Errorf("test 3.1 failed error: %s", err)
	}
	// test case 3.2 check has a 0 thumbnails
	if len(result.Thumbnails) != 0 {
		t.Errorf(
			"test 3.2 failed expected 0 thumbnails, got %v",
			len(result.Thumbnails),
		)
	}

	// Test case 4 invalid data
	headerData = external.MusicDetailHeaderRenderer{}
	err = utils.GetStructFromJson(
		"./data/json/detail_header/invalid.json",
		&headerData,
	)
	if err != nil {
		t.Errorf("test 4 failed error: %s", err)
	}
	// test case 4.1 check if the data has an error
	_, err = simplifier.MusicDetailHeaderRenderer(headerData)
	if err == nil {
		t.Errorf("test 4.1 failed expected an error, got nil")
	}
}
