package subtitle

import (
	"fmt"
	"katze/src/models"
	"katze/src/models/external"
)

// MapSubtitle maps the subtitle of a []models.FlexColumn
func ToFlexColumns(subtitle external.Subtitle) (
	[]models.FlexColumn, error,
) {

	var flexColumns []models.FlexColumn
	if subtitle.Runs == nil || len(subtitle.Runs) == 0 {

		err := fmt.Errorf("Could not find runs in subtitle")
		return flexColumns, err
	}

	for _, run := range subtitle.Runs {

		flexColumn, err := RunToFlexColumn(run)
		if err != nil {
			return flexColumns, err
		}

		flexColumns = append(flexColumns, flexColumn)
	}

	if len(flexColumns) == 0 {
		err := fmt.Errorf("Could not find flexColumns in subtitle")
		return flexColumns, err
	}
	return flexColumns, nil
}
