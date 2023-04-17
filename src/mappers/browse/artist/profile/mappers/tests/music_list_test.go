package tests

import (
	"katze/src/mappers/browse/artist/profile/mappers"
	"katze/src/models/external"
	"katze/src/utils"
	"testing"
)

func TestMusicList(t *testing.T) {
	// test case 1, get valid data 4 items
	musicListData := external.TentacledMusicShelfRenderer{}
	err := utils.GetStructFromJson(
		"./data/json/music_list/valid.json", &musicListData,
	)
	if err != nil {
		t.Errorf("test 1 failed, error: %v", err)
	}
	// test case 1.1, check if the data has an error
	result, err := mappers.MusicList(musicListData)
	if err != nil {
		t.Errorf("test 1.1 failed, error: %v", err)
	}
	// test case 1.2, check if the result has 4 items
	if len(result.Songs) != 4 {
		t.Errorf("test 1.2 failed, expected 4 items, got %v", len(result.Songs))
	}
	// test case 1.3, check if the result has a title "Espacio"
	if title := result.Songs[0].Title; title != "Espacio" {
		t.Errorf(
			"test 1.3 failed, expected title Espacio, got %v",
			title,
		)
	}

}
