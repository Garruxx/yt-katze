package data

import (
	"encoding/json"
	"katze/src/models/external"
	"os"
)

func Get(path string) (external.General, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return external.General{}, err
	}
	var data external.General
	err = json.Unmarshal(file, &data)
	if err != nil {
		return external.General{}, err
	}
	return data, nil
}
