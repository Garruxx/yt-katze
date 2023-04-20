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

// ArtistsPagination is a resolver for the graphql query
// Need continuationID, visitorID
func ArtistsPagination(gqlParams graphql.ResolveParams) (
	any, error,
) {

	continuationID, ok := gqlParams.Args["continuationId"].(string)
	if !ok {
		return lists.Artists{}, fmt.Errorf(
			"continuationId is not a string",
		)
	}
	visitorID, ok := gqlParams.Args["visitorId"].(string)
	if !ok {
		return lists.Artists{}, fmt.Errorf("visitorId is not a string")
	}

	resultData, err := services.SearchArtistPagination(
		continuationID, visitorID,
	)
	if err != nil {
		return lists.Artists{}, logger.Errorf("error %v", err)
	}

	result, err := mappers.SearchArtistPagination(
		external.ArtistPagination(resultData),
	)
	if err != nil {
		return lists.Artists{}, logger.Errorf("error %v", err)
	}

	return result, nil
}
