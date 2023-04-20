package search

import (
	"fmt"
	"katze/logger"
	"katze/src/mappers"
	"katze/src/models/external"
	"katze/src/models/lists"
	"katze/src/services"

	"github.com/graphql-go/graphql"
)

// AlbumsPagination is a resolver for the graphql query
// Need continuationID, visitorID
func AlbumsPagination(gqlParams graphql.ResolveParams) (
	any, error,
) {

	continuationID, ok := gqlParams.Args["continuationId"].(string)
	if !ok {
		return lists.Albums{}, fmt.Errorf("continuationId is not a string")
	}
	visitorID, ok := gqlParams.Args["visitorId"].(string)
	if !ok {
		return lists.Albums{}, fmt.Errorf("visitorId is not a string")
	}

	resultData, err := services.SearchAlbumPagination(
		continuationID, visitorID,
	)
	if err != nil {
		return lists.Albums{}, logger.Errorf("error %v", err)
	}

	result, err := mappers.SearchAlbumPagination(
		external.AlbumsPagination(resultData),
	)
	if err != nil {
		return lists.Albums{}, logger.Errorf("error %v", err)
	}

	return result, nil
}
