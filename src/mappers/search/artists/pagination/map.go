package pagination

import (
	"katze/src/mappers/utils/mappers"
	"katze/src/mappers/utils/simplifier"
	"katze/src/models/external"
	"katze/src/models/lists"
	"katze/src/utils"
)

// Map maps an external.ArtistPagination to a lists.Artists
func Map(paginationResult external.ArtistPagination) (lists.Artists, error) {

	pagination, err := simplifier.ItemsPagination(
		external.ItemsPagination(paginationResult),
	)
	if err != nil {
		return lists.Artists{}, err
	}

	artists, err := utils.ArrayMap(pagination.Contents, mappers.Artist)
	if err != nil {
		return lists.Artists{}, err
	}
	return lists.Artists{
		Artists:        artists,
		ContinuationID: pagination.NextPage,
		VisitorID:      pagination.VisitorID,
	}, nil
}
