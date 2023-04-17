package search

import (
	"encoding/json"
	"fmt"
	"katze/src/models/external"
	"katze/src/services/models"
	"katze/src/services/utils"
)

// ItemsList returns a list of items from a search query
func ItemsList(query string, param string, visitorID string) (
	external.ItemsList, error,
) {

	req := models.Request{
		Params:        param,
		Query:         query,
		GoogVisitorID: &visitorID,
		UrlPath:       "/youtubei/v1/search",
	}
	body, err := utils.Request(req)
	if err != nil {
		err := fmt.Errorf("error sending request: %v", err)
		return external.ItemsList{}, err
	}

	//Decode the response body into a itemslist struct
	var ItemsList external.ItemsList
	err = json.Unmarshal(body, &ItemsList)
	if err != nil {
		err := fmt.Errorf("error unmarshalling response body: %v", err)
		return external.ItemsList{}, err
	}

	if ItemsList.Error != nil {
		err := fmt.Errorf("error getting tracklist: %v", ItemsList.Error.Status)
		return external.ItemsList{}, err
	}

	return ItemsList, nil
}
