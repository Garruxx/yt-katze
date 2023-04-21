package browse

import (
	"encoding/json"
	"fmt"
	"katze/src/models/external"
	"katze/src/services/models"
	"katze/src/services/utils"
)

func ArtistMusicList(continuationID string, params string, visitorID *string) (
	external.ArtistTracklist, error,
) {

	req := models.Request{
		Params:        params,
		GoogVisitorID: visitorID,
		UrlPath:       "/youtubei/v1/browse",
		BrowseID:      continuationID,
	}
	body, err := utils.Request(req)

	if err != nil {
		err := fmt.Errorf("error sending request: %v", err)
		return external.ArtistTracklist{}, err
	}

	//Decode the response body into a artistTracklist struct
	var artistTracklist external.ArtistTracklist
	err = json.Unmarshal(body, &artistTracklist)
	if err != nil {
		err := fmt.Errorf("error unmarshalling response body: %v", err)
		return external.ArtistTracklist{}, err
	}
	if artistTracklist.Error != nil {
		err := fmt.Errorf(
			"error getting artistTracklist: %v",
			artistTracklist.Error.Status,
		)
		return external.ArtistTracklist{}, err
	}
	return artistTracklist, nil
}
