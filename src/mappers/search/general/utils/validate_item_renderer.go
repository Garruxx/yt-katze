package utils

import (
	"fmt"
	"katze/src/mappers/search/models"
)

func ValidateItemRenderer(
	itemRenderer models.MusicResponsiveListItemRenderer,
	pageTypeExpected string,
	flexColumnLength int,
) error {
	if itemRenderer.PageType != pageTypeExpected {
		err := fmt.Errorf("error: pageType is not %s", pageTypeExpected)
		return err
	}
	if len(itemRenderer.FlexColumns) != flexColumnLength {
		err := fmt.Errorf("error: flexColumns length is not %v", flexColumnLength)
		return err
	}
	if itemRenderer.ID == "" {
		err := fmt.Errorf("error: id is empty")
		return err
	}
	return nil
}
