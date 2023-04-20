package browse

import (
	"fmt"
	"katze/src/mappers"
	"katze/src/models/shelves"
	"katze/src/services"

	"github.com/graphql-go/graphql"
)

// Single is a resolver for the graphql query.
// need singleId and visitorId params
func Single(gqlParams graphql.ResolveParams) (
	any, error,
) {

	singleID, ok := gqlParams.Args["singleId"].(string)
	if !ok {
		return shelves.Single{}, fmt.Errorf("singleId is not a string")
	}
	visitorID, ok := gqlParams.Args["visitorId"].(string)
	if !ok {
		return shelves.Single{}, fmt.Errorf("visitorId is not a string")
	}

	resultData, err := services.BrowseSingle(singleID, &visitorID)
	if err != nil {
		return shelves.Single{}, err
	}

	result, err := mappers.BrowseSingle(resultData)
	if err != nil {
		return shelves.Single{}, err
	}

	return shelves.Single(result), nil
}
