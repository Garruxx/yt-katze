package player

import (
	"encoding/json"
	"fmt"
	"katze/src/models/external"
	"katze/src/services/models"
	"katze/src/services/utils"
)

func Information(videoID string, visitorID *string) (
	external.PlayerInformation, error,
) {

	req := models.Request{
		GoogVisitorID: visitorID,
		UrlPath:       "/youtubei/v1/next",
		VideoID:       videoID,
	}
	body, err := utils.Request(req)
	if err != nil {
		err := fmt.Errorf("error sending request: %v", err)
		return external.PlayerInformation{}, err
	}

	//Decode the response body into a trackpagination struct
	var album external.PlayerInformation
	err = json.Unmarshal(body, &album)
	if err != nil {
		err := fmt.Errorf("error unmarshalling response body: %v", err)
		return external.PlayerInformation{}, err
	}
	if album.Error != nil {
		err := fmt.Errorf("error getting album: %v", album.Error.Message)
		return external.PlayerInformation{}, err
	}
	return album, nil
}
