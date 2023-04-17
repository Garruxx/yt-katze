package mappers

import (
	"katze/src/mappers/browse/models"
	"katze/src/utils"
)

// CarouselList maps a MusicCarouselShelfRenderer to a CarouselList
func CarouselList(carouselShelf models.CarouselShelfRenderer) (
	models.CarouseList, error,
) {

	// Map items to cardItmes arr
	items, err := utils.ArrayMap(carouselShelf.Contents, TwoRowItemRenderer)
	if err != nil {
		return models.CarouseList{}, err
	}

	return models.CarouseList{
		Items:              items,
		ContinuationParams: carouselShelf.Params,
		ContinuationID:     carouselShelf.BrowseID,
	}, nil
}
