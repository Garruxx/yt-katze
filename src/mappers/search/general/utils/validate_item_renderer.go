package utils

import (
	"fmt"
	"katze/src/mappers/search/models"
)

func ValidateItemRenderer(
	itemRenderer models.MusicResponsiveListItemRenderer,
	pageTypeExpected string,
	length int,
) error {
	if itemRenderer.PageType != pageTypeExpected {
		err := fmt.Errorf("error: pageType is not %s", pageTypeExpected)
		return err
	}
	if len(itemRenderer.FlexColumns) != length {
		err := fmt.Errorf("error: flexColumns length is not %v", length)
		return err
	}
	if itemRenderer.ID == "" {
		err := fmt.Errorf("error: id is empty")
		return err
	}
	return nil
}
