package utils

import (
	"fmt"
	"katze/src/mappers/browse/models"
)

func FindCarouselShelf(
	carousels []models.CarouselShelfRenderer, query string,
) (
	models.CarouselShelfRenderer, error,
) {
	if len(carousels) == 0 {
		err := fmt.Errorf("carousels is empty")
		return models.CarouselShelfRenderer{}, err
	}
	for _, carousel := range carousels {
		if carousel.Title == query {
			return carousel, nil
		}
	}

	return models.CarouselShelfRenderer{}, fmt.Errorf("carousel not found")
}
