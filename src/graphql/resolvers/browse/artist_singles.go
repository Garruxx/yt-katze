package browse

import (
	"fmt"
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
	params, ok := gqlParams.Args["params"].(string)
	if !ok {
		return []artist.CardItem{}, fmt.Errorf("params is not a string")
	}

	resultData, err := services.BrowseArtistSingles(artistId, params, visitorID)
	if err != nil {
		return []artist.CardItem{}, err
	}

	result, err := mappers.BrowseArtistSingles(resultData)
	if err != nil {
		return []artist.CardItem{}, err
	}

	return result, nil
}
