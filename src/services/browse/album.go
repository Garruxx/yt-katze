package browse

import (
	"encoding/json"
	"fmt"
	"katze/src/models/external"
	"katze/src/services/models"
	"katze/src/services/utils"
)

func Album(BrowseID string, visitorID *string) (external.Tracklist, error) {

	req := models.Request{
		GoogVisitorID: visitorID,
		UrlPath:       "/youtubei/v1/browse",
		BrowseID:      BrowseID,
	}
	body, err := utils.Request(req)
	if err != nil {
		err := fmt.Errorf("error sending request: %v", err)
		return external.Tracklist{}, err
	}

	//Decode the response body into a trackpagination struct
	var album external.Tracklist
	err = json.Unmarshal(body, &album)
	if err != nil {
		err := fmt.Errorf("error unmarshalling response body: %v", err)
		return external.Tracklist{}, err
	}
	if album.Error != nil {
		err := fmt.Errorf("error getting album: %v", album.Error.Status)
		return external.Tracklist{}, err
	}
	return album, nil
}
