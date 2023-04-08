package browse

import (
	"encoding/json"
	"fmt"
	"katze/src/models/external"
	"katze/src/services/models"
	"katze/src/services/utils"
)

func Album(BrowseID string, visitorID *string) (external.Album, error) {

	req := models.Request{
		GoogVisitorID: visitorID,
		UrlPath:       "/youtubei/v1/browse",
		BrowseID:      BrowseID,
	}
	body, err := utils.Request(req)
	if err != nil {
		err := fmt.Errorf("Error sending request: %v", err)
		return external.Album{}, err
	}

	//Decode the response body into a trackpagination struct
	var album external.Album
	err = json.Unmarshal(body, &album)
	if err != nil {
		err := fmt.Errorf("Error unmarshalling response body: %v", err)
		return external.Album{}, err
	}
	if album.Error != nil {
		err := fmt.Errorf("Error getting album: %v", album.Error.Status)
		return external.Album{}, err
	}
	return album, nil
}
