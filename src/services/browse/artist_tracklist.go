package browse

import (
	"encoding/json"
	"fmt"
	"katze/src/models/external"
	"katze/src/services/models"
	"katze/src/services/utils"
)

func ArtistTracklist(browseID string, continuationID string, visitorID string) (
	external.ArtistTwoRowItem, error,
) {

	if visitorID == "" {
		err := fmt.Errorf("VisitorID is required")
		return external.ArtistTwoRowItem{}, err
	}

	req := models.Request{
		Params:        continuationID,
		GoogVisitorID: &visitorID,
		UrlPath:       "/youtubei/v1/browse",
		BrowseID:      browseID,
	}
	body, err := utils.Request(req)
	if err != nil {
		err := fmt.Errorf("error sending request: %v", err)
		return external.ArtistTwoRowItem{}, err
	}

	//Decode the response body into a artistTracklist struct
	var artistTracklist external.ArtistTwoRowItem
	err = json.Unmarshal(body, &artistTracklist)
	if err != nil {
		err := fmt.Errorf("error unmarshalling response body: %v", err)
		return external.ArtistTwoRowItem{}, err
	}
	if artistTracklist.Error != nil {
		err := fmt.Errorf(
			"error getting artistTracklist: %v", artistTracklist.Error.Status,
		)
		return external.ArtistTwoRowItem{}, err
	}
	return artistTracklist, nil
}
