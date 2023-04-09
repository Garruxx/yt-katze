package main

import (
	"encoding/json"
	"katze/logger"
	searchMaps "katze/src/mappers/search"
	"katze/src/services/search"
	"os"
)

func main() {

	result, err := search.MusicList("agua con chia sous sol", "EgWKAQIIAWoKEAMQBBAJEAoQBQ%3D%3D", "")
	if err != nil {
		panic(logger.Errorf("error: %v", err))
	}

	resultMapped, err := searchMaps.MusicList(result)
	if err != nil {
		panic(logger.Errorf("error: %v", err))
	}

	json, err := json.MarshalIndent(resultMapped, "", "	")
	if err != nil {
		panic(err)
	}
	os.WriteFile("response.json", json, 0777)

}
