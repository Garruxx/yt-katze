package search

import (
	"encoding/json"
	"fmt"
	"katze/src/models/external"
	"katze/src/services/models"
	"katze/src/services/utils"
)

func MusicPagination(continuationID string, visitorID *string) (
	external.MusicPagination, error,
) {

	req := models.Request{
		GoogVisitorID: visitorID,
		UrlPath:       "/youtubei/v1/search",
		UrlQueries:    fmt.Sprintf("continuation=%s", continuationID),
	}
	body, err := utils.Request(req)
	if err != nil {
		err := fmt.Errorf("error sending continuation request: %v", err)
		return external.MusicPagination{}, err
	}

	//Decode the response body into a trackpagination struct
	var trackPagination external.MusicPagination
	err = json.Unmarshal(body, &trackPagination)
	if err != nil {
		err := fmt.Errorf("error unmarshalling response body: %v", err)
		return external.MusicPagination{}, err
	}

	if _error := trackPagination.Error; _error != nil {
		err := fmt.Errorf("error getting continuation: %v", _error.Status)
		return external.MusicPagination{}, err
	}

	return trackPagination, nil
}
