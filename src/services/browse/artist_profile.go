package browse

import (
	"encoding/json"
	"fmt"
	"katze/src/models/external"
	"katze/src/services/models"
	"katze/src/services/utils"
)

func ArtistProfile(browseID string, visitorID *string) (
	external.Artist, error,
) {

	req := models.Request{
		GoogVisitorID: visitorID,
		UrlPath:       "/youtubei/v1/browse",
		BrowseID:      browseID,
	}
	body, err := utils.Request(req)
	if err != nil {
		err := fmt.Errorf("Error sending request: %v", err)
		return external.Artist{}, err
	}

	//Decode the response body into a artistProfile struct
	var artistProfile external.Artist
	err = json.Unmarshal(body, &artistProfile)
	if err != nil {
		err := fmt.Errorf("Error unmarshalling response body: %v", err)
		return external.Artist{}, err
	}
	if artistProfile.Error != nil {
		err := fmt.Errorf(
			"Error getting artistProfile: %v",
			artistProfile.Error.Status,
		)
		return external.Artist{}, err
	}
	return artistProfile, nil

}
