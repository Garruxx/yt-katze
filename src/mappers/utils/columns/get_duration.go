package columns

import (
	"fmt"
	"katze/src/models"
	"regexp"
)

// GetDuration gets the duration from the flexColumns
func GetDuration(flexColumns []models.FlexColumn) (string, error) {
	if len(flexColumns) <= 0 {
		err := fmt.Errorf("error: flexcolumns is empty")
		return "", err
	}

	// year regexp
	durationRegexp := regexp.MustCompile(`^(\d*:)?([0-5]?\d:)([0-5]\d)$`)
	duration := flexColumns[len(flexColumns)-1].Text
	if !durationRegexp.MatchString(duration) {
		err := fmt.Errorf("error: duration is not valid")
		return "", err
	}

	return duration, nil
}
