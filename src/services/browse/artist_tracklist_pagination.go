package browse

import (
	"encoding/json"
	"fmt"
	"katze/src/models/external"
	"katze/src/services/models"
	"katze/src/services/utils"
)

func ArtistTracklistPagination(continuationId string, visitorId *string) (
	external.ArtistTracklistPagination, error,
) {

	req := models.Request{
		UrlPath:       "/youtubei/v1/browse",
		UrlQueries:    fmt.Sprintf("continuation=%s", continuationId),
		GoogVisitorID: visitorId,
	}
	body, err := utils.Request(req)
	if err != nil {
		err := fmt.Errorf("Error sending request: %v", err)
		return external.ArtistTracklistPagination{}, err
	}

	//Decode the response body into a artistTracklistPagination struct
	var artistTracklistPagination external.ArtistTracklistPagination
	err = json.Unmarshal(body, &artistTracklistPagination)
	if err != nil {
		err := fmt.Errorf("Error unmarshalling response body: %v", err)
		return external.ArtistTracklistPagination{}, err
	}
	if artistTracklistPagination.Error != nil {
		err := fmt.Errorf(
			"Error getting artistTracklistPagination: %v",
			artistTracklistPagination.Error.Status,
		)
		return external.ArtistTracklistPagination{}, err
	}
	return artistTracklistPagination, nil
}
