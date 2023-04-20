package browse

import (
	"fmt"
	"katze/logger"
	"katze/src/mappers"
	"katze/src/models/lists"
	"katze/src/services"

	"github.com/graphql-go/graphql"
)

// ArtistMusicList the ArtistMusicList query
// need continuationId and visitorId params
func ArtistMusicsListPagination(
	gqlParams graphql.ResolveParams,
) (
	any, error,
) {

	continuationID, ok := gqlParams.Args["continuationId"].(string)
	if !ok {
		return lists.Music{}, fmt.Errorf("continuationId is not a string")
	}
	visitorID, ok := gqlParams.Args["visitorId"].(string)
	if !ok {
		return lists.Music{}, fmt.Errorf("visitorId is not a string")
	}

	resultData, err := services.BrowseArtistMusicPagination(
		continuationID, &visitorID,
	)
	if err != nil {
		return lists.Music{}, logger.Errorf("error %v", err)
	}

	result, err := mappers.BrowseArtistMusicPagination(resultData)
	if err != nil {
		return lists.Music{}, logger.Errorf("error %v", err)
	}

	return result, nil
}
