package search

import (
	"encoding/json"
	"fmt"
	"katze/src/models/external"
	"katze/src/services/models"
	"katze/src/services/utils"
)

func General(query string, visitorID string) (
	external.General, error,
) {

	req := models.Request{
		Query:         query,
		GoogVisitorID: &visitorID,
		UrlPath:       "/youtubei/v1/search",
	}
	body, err := utils.Request(req)
	if err != nil {
		err := fmt.Errorf("Error sending request: %v", err)
		return external.General{}, err
	}

	//Decode the response body into a General struct
	var trackPagination external.General
	err = json.Unmarshal(body, &trackPagination)
	if err != nil {
		err := fmt.Errorf("Error unmarshalling response body: %v", err)
		return external.General{}, err
	}

	if trackPagination.Error != nil {
		err := fmt.Errorf("Error getting search result: %v", trackPagination.Error.Status)
		return external.General{}, err
	}

	return trackPagination, nil
}
