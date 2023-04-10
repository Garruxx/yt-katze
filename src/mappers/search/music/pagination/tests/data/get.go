package data

import (
	"encoding/json"
	"katze/src/models/external"
	"os"
)

func Get(path string) (external.MusicPagination, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return external.MusicPagination{}, err
	}
	var data external.MusicPagination
	err = json.Unmarshal(file, &data)
	if err != nil {
		return external.MusicPagination{}, err
	}
	return data, nil
}
