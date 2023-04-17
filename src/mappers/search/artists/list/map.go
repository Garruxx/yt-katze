package list

import (
	"katze/src/mappers/utils/mappers"
	"katze/src/mappers/utils/simplifier"
	"katze/src/models/external"
	"katze/src/models/lists"
	"katze/src/utils"
)

// Map maps an external.ArtistList to a lists.Artists
func Map(artistListResult external.ArtistList) (lists.Artists, error) {

	artistsList, err := simplifier.ItemsList(
		external.ItemsList(artistListResult),
	)
	if err != nil {
		return lists.Artists{}, err
	}

	artists, err := utils.ArrayMap(artistsList.Contents, mappers.Artist)
	if err != nil {
		return lists.Artists{}, err
	}

	return lists.Artists{
		Artists:        artists,
		ContinuationID: artistsList.NextPage,
		VisitorID:      artistsList.VisitorID,
	}, nil
}
