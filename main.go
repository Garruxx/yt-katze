package main

import (
	"encoding/json"
	"katze/src/mappers/search/general"
	"katze/src/services/search"
	"os"
)

func main() {

	result, err := search.General("Agua con chia", "")
	if err != nil {
		panic(err)
	}

	resultMapped, err := general.Map(result)
	if err != nil {
		panic(err)
	}

	json, err := json.MarshalIndent(resultMapped, "", "	")
	if err != nil {
		panic(err)
	}
	os.WriteFile("response.json", json, 777)

}
