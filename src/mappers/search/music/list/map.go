package list

import (
	"katze/src/mappers/utils/mappers"
	"katze/src/mappers/utils/simplifier"
	"katze/src/models/external"
	"katze/src/models/lists"
	"katze/src/utils"
)

// Map maps an external.MusicList to a lists.Music
func Map(musicResults external.MusicList) (lists.Music, error) {

	itemsList, err := simplifier.ItemsList(external.ItemsList(musicResults))
	if err != nil {
		return lists.Music{}, err
	}

	musicItems, err := utils.ArrayMap(itemsList.Contents, mappers.Song)
	if err != nil {
		return lists.Music{}, err
	}

	return lists.Music{
		Songs:          musicItems,
		ContinuationID: itemsList.NextPage,
		VisitorID:      itemsList.VisitorID,
	}, nil
}
