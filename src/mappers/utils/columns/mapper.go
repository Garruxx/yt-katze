package columns

import (
	"fmt"
	"katze/src/mappers/models"
	"katze/src/models/external"
)

// Mapper gets an external column.FluffyFlexColumn
// and returns the data mapped in models.ColumnsItems
func Mapper(flexColumns []external.FluffyFlexColumn) (
	[]models.ColumnsItems, error,
) {

	var result []models.ColumnsItems
	if len(flexColumns) < 1 {
		err := fmt.Errorf("error: flexcolumns is empty")
		return []models.ColumnsItems{}, err
	}

	for _, flexItem := range flexColumns {

		// Get the column and declare a variable to store the items
		column := flexItem.MusicResponsiveListItemFlexColumnRenderer.Text.Runs
		items := models.MultiColumn{}

		if len(column) < 1 {
			continue
		}

		for _, item := range column {
			flexColumn, err := ColumnMapper(item)
			if err != nil {
				return []models.ColumnsItems{}, err
			}
			items = append(items, flexColumn)
		}

		// Create a new ColumnsItems and append it to the result
		result = append(result, models.ColumnsItems{
			Items: items,
		})
	}

	return result, nil
}
