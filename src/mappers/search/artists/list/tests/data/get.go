package data

import (
	"encoding/json"
	"katze/src/models/external"
	"os"
)

func Get(path string) (external.ArtistList, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return external.ArtistList{}, err
	}
	var data external.ArtistList
	err = json.Unmarshal(file, &data)
	if err != nil {
		return external.ArtistList{}, err
	}
	return data, nil
}
