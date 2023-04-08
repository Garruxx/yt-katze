package browse

import (
	"encoding/json"
	"fmt"
	"katze/src/models/external"
	"katze/src/services/models"
	"katze/src/services/utils"
)

func ArtistAlbums(browseID string, params string, visitorID string) (
	external.ArtistAlbums, error,
) {

	if visitorID == "" {
		err := fmt.Errorf("VisitorID is required")
		return external.ArtistAlbums{}, err
	}

	req := models.Request{
		Params:        params,
		GoogVisitorID: &visitorID,
		UrlPath:       "/youtubei/v1/browse",
		BrowseID:      browseID,
	}
	body, err := utils.Request(req)
	if err != nil {
		err := fmt.Errorf("Error sending request: %v", err)
		return external.ArtistAlbums{}, err
	}

	//Decode the response body into a artistAlbums struct
	var artistAlbums external.ArtistAlbums
	err = json.Unmarshal(body, &artistAlbums)
	if err != nil {
		err := fmt.Errorf("Error unmarshalling response body: %v", err)
		return external.ArtistAlbums{}, err
	}
	if artistAlbums.Error != nil {
		err := fmt.Errorf(
			"Error getting artistAlbums: %v", artistAlbums.Error.Status,
		)
		return external.ArtistAlbums{}, err
	}
	return artistAlbums, nil
}
