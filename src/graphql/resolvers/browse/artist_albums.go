package browse

import (
	"fmt"
	"katze/logger"
	"katze/src/mappers"
	"katze/src/models/artist"
	"katze/src/services"

	"github.com/graphql-go/graphql"
)

// ArtistAlbums resolves the albumArtist query
func ArtistAlbums(hqlParams graphql.ResolveParams) (any, error) {

	artistID, ok := hqlParams.Args["artistId"].(string)
	if !ok {
		return []artist.CardItem{}, fmt.Errorf("artistId is not a string")
	}
	visitorID, ok := hqlParams.Args["visitorId"].(string)
	if !ok {
		return []artist.CardItem{}, fmt.Errorf("visitorId is not a string")
	}
	params, ok := hqlParams.Args["params"].(string)
	if !ok {
		return []artist.CardItem{}, fmt.Errorf("params is not a string")
	}

	resultData, err := services.BrowseArtistAlbums(artistID, params, visitorID)
	if err != nil {
		return []artist.CardItem{}, logger.Errorf("error %v", err)
	}
	result, err := mappers.BrowseArtistAlbums(resultData)
	if err != nil {
		return []artist.CardItem{}, logger.Errorf("error %v", err)
	}

	return result, nil
}
