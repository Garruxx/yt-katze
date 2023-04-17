package list

import (
	"katze/src/mappers/utils/mappers"
	"katze/src/mappers/utils/simplifier"
	"katze/src/models/external"
	"katze/src/models/shelves"
	"katze/src/utils"
)

// Map maps an external.MusicList to a shelves.Music
func Map(musicResults external.MusicList) (shelves.Music, error) {

	itemsList, err := simplifier.ItemsList(external.ItemsList(musicResults))
	if err != nil {
		return shelves.Music{}, err
	}

	musicItems, err := utils.ArrayMap(itemsList.Contents, mappers.Song)
	if err != nil {
		return shelves.Music{}, err
	}

	return shelves.Music{
		Songs:          musicItems,
		ContinuationID: itemsList.NextPage,
		VisitorID:      itemsList.VisitorID,
	}, nil
}
