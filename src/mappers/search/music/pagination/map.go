package pagination

import (
	"katze/src/mappers/utils/mappers"
	"katze/src/mappers/utils/simplifier"
	"katze/src/models/external"
	"katze/src/models/lists"
	"katze/src/utils"
)

// Map maps an external.MusicPagination to a lists.Music
func Map(paginationResult external.MusicPagination) (lists.Music, error) {

	pagination, err := simplifier.ItemsPagination(
		external.ItemsPagination(paginationResult),
	)
	if err != nil {
		return lists.Music{}, err
	}

	musicItems, err := utils.ArrayMap(pagination.Contents, mappers.Song)
	if err != nil {
		return lists.Music{}, err
	}
	return lists.Music{
		Songs:          musicItems,
		ContinuationID: pagination.NextPage,
		VisitorID:      pagination.VisitorID,
	}, nil
}
