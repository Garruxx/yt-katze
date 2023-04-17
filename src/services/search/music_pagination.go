package search

import (
	"encoding/json"
	"fmt"
	"katze/src/models/external"
	"katze/src/services/models"
	"katze/src/services/utils"
)

// ItemsPagination returns a list of items from a continuationID
func ItemsPagination(continuationID string, visitorID string) (
	external.ItemsPagination, error,
) {

	req := models.Request{
		GoogVisitorID: &visitorID,
		UrlPath:       "/youtubei/v1/search",
		UrlQueries:    fmt.Sprintf("continuation=%s", continuationID),
	}
	body, err := utils.Request(req)
	if err != nil {
		err := fmt.Errorf("error sending continuation request: %v", err)
		return external.ItemsPagination{}, err
	}

	//Decode the response body into a item pagination struct
	var itemsPagination external.ItemsPagination
	err = json.Unmarshal(body, &itemsPagination)
	if err != nil {
		err := fmt.Errorf("error unmarshalling response body: %v", err)
		return external.ItemsPagination{}, err
	}

	if _error := itemsPagination.Error; _error != nil {
		err := fmt.Errorf("error getting continuation: %v", _error.Status)
		return external.ItemsPagination{}, err
	}

	return itemsPagination, nil
}
