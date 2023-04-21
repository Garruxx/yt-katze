package browse

import (
	"fmt"
	"katze/logger"
	"katze/src/mappers"
	"katze/src/models/artist"
	"katze/src/services"

	"github.com/graphql-go/graphql"
)

// ArtistSingles resolves the artistSingles query
func ArtistSingles(
	gqlParams graphql.ResolveParams) (any, error) {

	artistId, ok := gqlParams.Args["artistId"].(string)
	if !ok {
		return []artist.CardItem{}, fmt.Errorf("artistId is not a string")
	}
	visitorID, ok := gqlParams.Args["visitorId"].(string)
	if !ok {
		return []artist.CardItem{}, fmt.Errorf("visitorId is not a string")
	}
	continuationID, ok := gqlParams.Args["continuationId"].(string)
	if !ok {
		return []artist.CardItem{}, fmt.Errorf("continuationId is not a string")
	}

	resultData, err := services.BrowseArtistSingles(
		artistId, continuationID, visitorID,
	)
	if err != nil {
		return []artist.CardItem{}, logger.Errorf("error %v", err)
	}

	result, err := mappers.BrowseArtistSingles(resultData)
	if err != nil {
		return []artist.CardItem{}, logger.Errorf("error %v", err)
	}

	return result, nil
}
