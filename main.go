package main

import (
	"encoding/json"
	"katze/logger"
	searchMaps "katze/src/mappers/search"
	"katze/src/services/search"
	"os"
)

func main() {

	result, err := search.General("agua con chia sous sol", "")
	if err != nil {
		panic(logger.Errorf("error: %v", err))
	}

	resultMapped, err := searchMaps.GeneralResult(result)
	if err != nil {
		panic(logger.Errorf("error: %v", err))
	}

	json, err := json.MarshalIndent(resultMapped, "", "	")
	if err != nil {
		panic(err)
	}
	os.WriteFile("response.json", json, 0777)

}
