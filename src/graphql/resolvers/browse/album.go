package browse

import (
	"fmt"
	"katze/logger"
	"katze/src/mappers"
	"katze/src/models/shelves"
	"katze/src/services"

	"github.com/graphql-go/graphql"
)

// Album resolves the album query.
func Album(
	gqlParams graphql.ResolveParams) (any, error) {

	albumID, ok := gqlParams.Args["albumId"].(string)
	if !ok {
		return shelves.Album{}, fmt.Errorf("albumId is not a string")
	}
	visitorID, ok := gqlParams.Args["visitorId"].(string)
	if !ok {
		return shelves.Album{}, fmt.Errorf("visitorId is not a string")
	}

	resultData, err := services.BrowseAlbum(albumID, &visitorID)
	if err != nil {
		return shelves.Album{}, logger.Errorf("error %v", err)
	}

	result, err := mappers.BrowseAlbum(resultData)
	if err != nil {
		return shelves.Album{}, logger.Errorf("error %v", err)
	}

	return result, nil
}
