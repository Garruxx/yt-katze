package browse

import (
	"fmt"
	"katze/logger"
	"katze/src/mappers"
	"katze/src/models/music"
	"katze/src/services"

	"github.com/graphql-go/graphql"
)

// SongsRecomendation is a resolver for the songs recomendation
func SongsRecomendation(params graphql.ResolveParams) (any, error) {

	//Get the songID
	songID, ok := params.Args["songId"].(string)
	if !ok {
		return []music.Song{}, fmt.Errorf("songId is not a string")
	}
	//Get the visitorID
	visitorID, ok := params.Args["visitorId"].(string)
	if !ok {
		return []music.Song{}, fmt.Errorf("visitorId is not a string")
	}

	//Get the browseID for the recomendations
	player, err := services.PlayerInformation(songID, &visitorID)
	if err != nil {
		return []music.Song{}, err
	}
	playerMapped, err := mappers.PlayerInformation(player)
	if err != nil {
		return []music.Song{}, err
	}

	//Get the recomendations
	recomendations, err := services.BrowseRecomendations(
		playerMapped.Related.BrowseID, visitorID,
	)
	if err != nil {
		return []music.Song{}, logger.Errorf("error %v",err)
	}
	recomendationsMapped, err := mappers.BeowseRecomendations(
		recomendations,
	)
	if err != nil {
		return []music.Song{}, logger.Errorf("error %v %s  %s",err, 
		playerMapped.Related.BrowseID, visitorID)
	}

	return recomendationsMapped, nil
}
