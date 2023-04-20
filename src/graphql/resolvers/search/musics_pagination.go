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

// MusicsPagination is a resolver for the graphql query
// need continuationId, visitorId
func MusicsPagination(gqlParams graphql.ResolveParams) (
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

	resultData, err := services.SearchMusicPagination(continuationID, visitorID)
	if err != nil {
		return lists.Music{}, logger.Errorf("error %v", err)
	}

	result, err := mappers.SearchMusicPagination(
		external.MusicPagination(resultData),
	)
	if err != nil {
		return lists.Music{}, logger.Errorf("error %v", err)
	}

	return result, nil
}
