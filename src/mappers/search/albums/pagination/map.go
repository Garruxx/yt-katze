package pagination

import (
	"katze/src/mappers/utils/mappers"
	"katze/src/mappers/utils/simplifier"
	"katze/src/models/external"
	"katze/src/models/lists"
	"katze/src/utils"
)

func Map(paginationResult external.AlbumsPagination) (
	lists.Albums, error,
) {
	pagination, err := simplifier.ItemsPagination(
		external.ItemsPagination(paginationResult),
	)
	if err != nil {
		return lists.Albums{}, err
	}

	albums, err := utils.ArrayMap(pagination.Contents, mappers.Album)
	if err != nil {
		return lists.Albums{}, err
	}
	return lists.Albums{
		Albums:         albums,
		ContinuationID: pagination.NextPage,
		VisitorID:      pagination.VisitorID,
	}, nil
}
