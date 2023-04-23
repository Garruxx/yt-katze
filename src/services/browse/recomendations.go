package browse

import (
	"encoding/json"
	"fmt"
	"katze/src/models/external"
	"katze/src/services/models"
	"katze/src/services/utils"
)

func Recomendations(browseID string, visitorID string) (
	external.MusicRecomendations, error,
) {

	if visitorID == "" {
		err := fmt.Errorf("VisitorID is required")
		return external.MusicRecomendations{}, err
	}

	req := models.Request{
		GoogVisitorID: &visitorID,
		UrlPath:       "/youtubei/v1/browse",
		BrowseID:      browseID,
	}
	body, err := utils.Request(req)
	if err != nil {
		err := fmt.Errorf("error sending request: %v", err)
		return external.MusicRecomendations{}, err
	}

	//Decode the response body into a recomendations struct
	var recomendations external.MusicRecomendations
	err = json.Unmarshal(body, &recomendations)
	if err != nil {
		err := fmt.Errorf("error unmarshalling response body: %v", err)
		return external.MusicRecomendations{}, err
	}
	if recomendations.Error != nil {
		err := fmt.Errorf(
			"error getting recomendations: %v", recomendations.Error.Message,
		)
		return external.MusicRecomendations{}, err
	}
	return recomendations, nil
}
