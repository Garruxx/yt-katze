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

// AlbumsList is a resolver for the graphql query
// need query, params, visitorID
func AlbumsList(gqlParams graphql.ResolveParams) (
	any, error,
) {

	query, ok := gqlParams.Args["query"].(string)
	if !ok {
		return lists.Albums{}, fmt.Errorf("query is not a string")
	}
	visitorID, ok := gqlParams.Args["visitorId"].(string)
	if !ok {
		return lists.Albums{}, fmt.Errorf("visitorId is not a string")
	}
	params, ok := gqlParams.Args["params"].(string)
	if !ok {
		return lists.Albums{}, fmt.Errorf("params is not a string")
	}
	resultData, err := services.SearchAlbumList(query, params, visitorID)
	if err != nil {
		return lists.Albums{}, logger.Errorf("error %v", err)
	}

	result, err := mappers.SearchAlbumList(external.AlbumList(resultData))
	if err != nil {
		return lists.Albums{}, logger.Errorf("error %v", err)
	}

	return result, nil
}
