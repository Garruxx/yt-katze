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
	albums := artistModels.AlbumsList{
		Albums:             albumsData.Items,
		ArtistID:           albumsData.ArtistID,
		ContinuationID: albumsData.ContinuationID,
	}

	singles := artistModels.SinglesList{
		Singles:            singlesData.Items,
		ArtistID:           singlesData.ArtistID,
		ContinuationID: singlesData.ContinuationID,
	}

	//artist music list
	musicsList, err := artistMappers.MusicsList(
		artistSimplified.MusicShelfRenderer,
	)
	if err != nil && err.Error() != "no music found" {
		return shelf, err
	}
	shelf = shelves.Artist{
		Name:        artistSimplified.Name,
		Thumbnails:  artistSimplified.Thumbnails,
		MusicsList:  musicsList,
		AlbumsList:  albums,
		SinglesList: singles,
		VisitorID:   artistSimplified.VisitorID,
	}
	return shelf, nil
}
