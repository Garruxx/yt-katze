package profile

import (
	artistMappers "katze/src/mappers/browse/artist/profile/mappers"
	"katze/src/mappers/browse/artist/profile/simplifier"
	browseUtils "katze/src/mappers/browse/utils"
	"katze/src/mappers/browse/utils/mappers"
	browseSimplifier "katze/src/mappers/browse/utils/simplifier"
	artistModels "katze/src/models/artist"
	"katze/src/models/external"
	"katze/src/models/shelves"
	"katze/src/utils"
)

func Map(artist external.Artist) (
	shelf shelves.Artist, err error,
) {
	artistSimplified, err := simplifier.Artist(artist)
	if err != nil {
		return shelf, err
	}
	// Simplify carousels
	carouselsSimplified, err := utils.ArrayMap(
		artistSimplified.Carousels,
		browseSimplifier.CarouselShelfRenderer,
	)
	if err != nil {
		return shelf, err
	}

	// Find carousels
	albumsCarousel, err := browseUtils.FindCarouselShelf(
		carouselsSimplified, "Albums",
	)
	if err != nil && err.Error() != "carousel not found" {
		return shelf, err
	}

	singlesCarousel, err := browseUtils.FindCarouselShelf(
		carouselsSimplified, "Singles",
	)
	if err != nil && err.Error() != "carousel not found" {
		return shelf, err
	}

	// Map carousels
	albumsData, err := mappers.CarouselList(albumsCarousel)
	if err != nil {
		return shelf, err
	}

	singlesData, err := mappers.CarouselList(singlesCarousel)
	if err != nil {
		return shelf, err
	}

	// Map albums and singles
	albums := artistModels.AlbumList{
		Albums:             albumsData.Items,
		ContinuationID:     albumsData.ContinuationID,
		ContinuationParams: albumsData.ContinuationParams,
	}

	singles := artistModels.SingleList{
		Singles:            singlesData.Items,
		ContinuationID:     singlesData.ContinuationID,
		ContinuationParams: singlesData.ContinuationParams,
	}

	//artist music list
	musicList, err := artistMappers.MusicList(
		artistSimplified.MusicShelfRenderer,
	)
	if err != nil && err.Error() != "no music found" {
		return shelf, err
	}
	shelf = shelves.Artist{
		Name:        artistSimplified.Name,
		Thumbnails:  artistSimplified.Thumbnails,
		MusicList:   musicList,
		AlbumsList:  albums,
		SinglesList: singles,
		VisitorID:   artistSimplified.VisitorID,
	}
	return shelf, nil
}
