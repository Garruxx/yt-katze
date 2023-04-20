package browse

import (
	"fmt"
	"katze/logger"
	"katze/src/mappers"
	"katze/src/models/lists"
	"katze/src/services"

	"github.com/graphql-go/graphql"
)

// ArtistMusicList the ArtistMusicList query
// need playlistId, visitorId and params params
func ArtistMusicsList(
	gqlParams graphql.ResolveParams,
) (
	any, error,
) {

	playlistID, ok := gqlParams.Args["playlistId"].(string)
	if !ok {
		return lists.Music{}, fmt.Errorf("playlistId is not a string")
	}
	visitorID, ok := gqlParams.Args["visitorId"].(string)
	if !ok {
		return lists.Music{}, fmt.Errorf("visitorId is not a string")
	}
	params, ok := gqlParams.Args["params"].(string)
	if !ok {
		return lists.Music{}, fmt.Errorf("params is not a string")
	}

	resultData, err := services.BrowseArtistMusicList(
		playlistID, params, &visitorID,
	)
	if err != nil {
		return lists.Music{}, logger.Errorf("error %v", err)
	}

	result, err := mappers.BrowseArtistMusicList(resultData)
	if err != nil {
		return lists.Music{}, logger.Errorf("error %v", err)
	}

	return result, nil
}
