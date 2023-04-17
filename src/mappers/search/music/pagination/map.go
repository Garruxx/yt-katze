package pagination

import (
	"katze/src/mappers/utils/mappers"
	"katze/src/mappers/utils/simplifier"
	"katze/src/models/external"
	"katze/src/models/shelves"
	"katze/src/utils"
)

// Map maps an external.MusicPagination to a shelves.Music
func Map(paginationResult external.MusicPagination) (shelves.Music, error) {

	pagination, err := simplifier.ItemsPagination(
		external.ItemsPagination(paginationResult),
	)
	if err != nil {
		return shelves.Music{}, err
	}

	musicItems, err := utils.ArrayMap(pagination.Contents, mappers.Song)
	if err != nil {
		return shelves.Music{}, err
	}
	return shelves.Music{
		Songs:          musicItems,
		ContinuationID: pagination.NextPage,
		VisitorID:      pagination.VisitorID,
	}, nil
}
