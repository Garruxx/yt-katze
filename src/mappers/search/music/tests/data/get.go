package data

import (
	"encoding/json"
	"katze/src/models/external"
	"os"
)

func Get(path string) (external.MusicList, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return external.MusicList{}, err
	}
	var data external.MusicList
	err = json.Unmarshal(file, &data)
	if err != nil {
		return external.MusicList{}, err
	}
	return data, nil
}
