package search

import (
	"fmt"
	"katze/src/mappers"
	"katze/src/models/external"
	"katze/src/models/lists"
	"katze/src/services"

	"github.com/graphql-go/graphql"
)

// ArtistsList is a resolver for the graphql query
// need query, visitorId and params params
func ArtistsList(gqlParams graphql.ResolveParams) (
	any, error,
) {

	query, ok := gqlParams.Args["query"].(string)
	if !ok {
		return lists.Artists{}, fmt.Errorf("query is not a string")
	}
	visitorID, ok := gqlParams.Args["visitorId"].(string)
	if !ok {
		return lists.Artists{}, fmt.Errorf("visitorId is not a string")
	}
	params, ok := gqlParams.Args["params"].(string)
	if !ok {
		return lists.Artists{}, fmt.Errorf("params is not a string")
	}
	resultData, err := services.SearchArtistList(query, params, visitorID)
	if err != nil {
		return lists.Artists{}, err
	}

	result, err := mappers.SearchArtistList(external.ArtistList(resultData))
	if err != nil {
		return lists.Artists{}, err
	}

	return result, nil
}
