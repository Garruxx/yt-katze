package browse

import (
	"fmt"
	"katze/logger"
	"katze/src/mappers"
	"katze/src/models/shelves"
	"katze/src/services"

	"github.com/graphql-go/graphql"
)

// ArtistProfile resolves the artist profile query.
// need artistId and visitorId params
func ArtistProfile(agqlParams graphql.ResolveParams) (
	any, error,
) {

	artistID, ok := agqlParams.Args["artistId"].(string)
	if !ok {
		return shelves.Artist{}, fmt.Errorf("artistId is not a string")
	}
	visitorID, ok := agqlParams.Args["visitorId"].(string)
	if !ok {
		return shelves.Artist{}, fmt.Errorf("visitorId is not a string")
	}

	resultData, err := services.BrowseArtistProfile(artistID, &visitorID)
	if err != nil {
		return shelves.Artist{}, logger.Errorf("error %v", err)
	}

	result, err := mappers.BrowseArtistProfile(resultData)
	if err != nil {
		return shelves.Artist{}, logger.Errorf("error %v", err)
	}

	return result, nil
}
