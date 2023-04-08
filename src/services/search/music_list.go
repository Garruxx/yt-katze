package search

import (
	"encoding/json"
	"fmt"
	"katze/src/models/external"
	"katze/src/services/models"
	"katze/src/services/utils"
)

func MusicList(query string, param string, visitorID *string) (
	external.MusicList, error,
) {

	req := models.Request{
		Params:        param,
		Query:         query,
		GoogVisitorID: visitorID,
		UrlPath:       "/youtubei/v1/search",
	}
	body, err := utils.Request(req)
	if err != nil {
		err := fmt.Errorf("error sending request: %v", err)
		return external.MusicList{}, err
	}

	//Decode the response body into a Tracklist struct
	var tracklist external.MusicList
	err = json.Unmarshal(body, &tracklist)
	if err != nil {
		err := fmt.Errorf("error unmarshalling response body: %v", err)
		return external.MusicList{}, err
	}

	if tracklist.Error != nil {
		err := fmt.Errorf("error getting tracklist: %v", tracklist.Error.Status)
		return external.MusicList{}, err
	}

	return tracklist, nil
}
