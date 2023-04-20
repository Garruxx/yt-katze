package search

import (
	"fmt"
	"katze/logger"
	"katze/src/mappers"
	"katze/src/models/shelves"
	"katze/src/services"

	"github.com/graphql-go/graphql"
)

// General is a resolver for the general search
// need search and visitorId params
func General(params graphql.ResolveParams) (any, error) {

	query, ok := params.Args["query"].(string)
	if !ok {
		return shelves.General{}, fmt.Errorf("query is not a string")
	}
	visitorID, ok := params.Args["visitorId"].(string)
	if !ok {
		return shelves.General{}, fmt.Errorf("visitorId is not a string")
	}

	resultData, err := services.SearchGeneral(query, visitorID)
	if err != nil {
		return shelves.General{}, logger.Errorf("error %v", err)
	}

	result, err := mappers.SearchGeneral(resultData)
	if err != nil {
		return shelves.General{}, logger.Errorf("error %v", err)
	}

	return result, nil
}
