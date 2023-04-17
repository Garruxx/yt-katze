package list

import (
	"katze/src/mappers/utils/mappers"
	"katze/src/mappers/utils/simplifier"
	"katze/src/models/external"
	"katze/src/models/lists"
	"katze/src/utils"
)

// Map maps an external.AlbumList to a lists.Albums
func Map(albumsResults external.AlbumList) (
	lists.Albums, error,
) {
	albumList, err := simplifier.ItemsList(external.ItemsList(albumsResults))
	if err != nil {
		return lists.Albums{}, err
	}

	albums, err := utils.ArrayMap(albumList.Contents, mappers.Album)
	if err != nil {
		return lists.Albums{}, err
	}

	return lists.Albums{
		Albums:         albums,
		ContinuationID: albumList.NextPage,
		VisitorID:      albumList.VisitorID,
	}, nil
}
