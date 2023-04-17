package tests

import (
	"katze/src/mappers/browse/tracklist"
	"katze/src/models/external"
	"katze/src/utils"
	"testing"
)

func TestMap(t *testing.T) {

	// Tests case 1 - valid data Album
	tracklistData := external.Tracklist{}
	err := utils.GetStructFromJson(
		"./data/json/valid_album.json",
		&tracklistData,
	)
	if err != nil {
		t.Errorf("test 1 failed error: %s", err)
	}
	// test case 1.1 check if the data has an error
	result, err := tracklist.Map(tracklistData)
	if err != nil {
		t.Errorf("test 1.1 failed error: %s", err)
	}
	// test case 1.2 check has a title "Espacio"
	if result.Title != "Espacio" {
		t.Errorf(
			"test 1.2 failed expected a title Espacio, got %v",
			result.Title,
		)
	}
	// test case 1.3 check has a songs length 4
	if len(result.Tracks) != 4 {
		t.Errorf(
			"test 1.3 failed expected a songs length 4, got %v",
			len(result.Tracks),
		)
	}
	// test case 1.4 check has a year 2024
	if result.Year != 2024 {
		t.Errorf(
			"test 1.4 failed expected a year 2024, got %v",
			result.Year,
		)
	}
	// test case 1.5 check has a 1 artist
	if len(result.Artists) != 1 {
		t.Errorf(
			"test 1.5 failed expected 1 artist, got %v",
			len(result.Artists),
		)
	}
	// test case 1.6 check has a visitorID "visitorIDTest"
	if result.VisitorID != "visitorIDTest" {
		t.Errorf(
			"test 1.6 failed expected a visitorIDTest, got %v",
			result.VisitorID,
		)
	}
	// test case 1.7 check has duration 22 minutos y 4 segundos
	if result.Duration != "22 minutos y 4 segundos" {
		t.Errorf(
			"test 1.7 failed expected duration 22 minutos y 4 segundos, got %v",
			result.Duration,
		)
	}

	// Tests case 2 - valid data single
	tracklistData = external.Tracklist{}
	err = utils.GetStructFromJson(
		"./data/json/valid_single.json",
		&tracklistData,
	)
	if err != nil {
		t.Errorf("test 2 failed error: %s", err)
	}
	// test case 2.1 check if the data has an error
	result, err = tracklist.Map(tracklistData)
	if err != nil {
		t.Errorf("test 2.1 failed error: %s", err)
	}
	// test case 2.2 check has a title "Mi berrinche"
	if result.Title != "Ya No" {
		t.Errorf(
			"test 2.2 failed expected a title Ya No, got %v",
			result.Title,
		)
	}
	// test case 2.3 check has a songs length 1
	if len(result.Tracks) != 1 {
		t.Errorf(
			"test 2.3 failed expected a songs length 1, got %v",
			len(result.Tracks),
		)
	}
	// test case 2.4 check has a year 2024
	if result.Year != 2024 {
		t.Errorf(
			"test 2.4 failed expected a year 2024, got %v",
			result.Year,
		)
	}
	// test case 2.5 check has a 2 artist
	if len(result.Artists) != 2 {
		t.Errorf(
			"test 2.5 failed expected 2 artist, got %v",
			len(result.Artists),
		)
	}
	// test case 2.6 check has a visitorID "visitorIDTest"
	if result.VisitorID != "visitorIDTest" {
		t.Errorf(
			"test 2.6 failed expected a visitorIDTest, got %v",
			result.VisitorID,
		)
	}
	// test case 2.7 check has duration 2244 minutos
	if result.Duration != "2244 minutos" {
		t.Errorf(
			"test 2.7 failed expected duration 2244 minutos, got %v",
			result.Duration,
		)
	}
}
