package utils

import (
	"encoding/json"
	"io/ioutil"
)

func GetStructFromJson(filename string, structure interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &structure)
	if err != nil {
		return err
	}

	return nil
}
