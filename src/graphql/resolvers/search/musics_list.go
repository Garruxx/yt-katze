package search

import (
	"fmt"
	"katze/src/mappers"
	"katze/src/models/external"
	"katze/src/models/lists"
	"katze/src/services"

	"github.com/graphql-go/graphql"
)

// MusicsList is a resolver for the graphql query
// need query, visitorId and params
func MusicsList(gqlParams graphql.ResolveParams) (
	any, error,
) {

	query, ok := gqlParams.Args["query"].(string)
	if !ok {
		return lists.Music{}, fmt.Errorf("query is not a string")
	}
	visitorID, ok := gqlParams.Args["visitorId"].(string)
	if !ok {
		return lists.Music{}, fmt.Errorf("visitorId is not a string")
	}
	params, ok := gqlParams.Args["params"].(string)
	if !ok {
		return lists.Music{}, fmt.Errorf("params is not a string")
	}

	resultData, err := services.SearchMusicList(query, params, visitorID)
	if err != nil {
		return lists.Music{}, err
	}

	result, err := mappers.SearchMusicList(external.MusicList(resultData))
	if err != nil {
		return lists.Music{}, err
	}

	return result, nil
}
