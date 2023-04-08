package columns

import (
	"fmt"
	"katze/src/models"
	"regexp"
	"strconv"
)

func GetYear(flexColumns []models.FlexColumn) (int, error) {

	if len(flexColumns) == 0 {
		err := fmt.Errorf("error: flexcolumns is empty")
		return 0, err
	}

	// year regexp
	yearRegexp := regexp.MustCompile(`^\d{4}$`)
	year := flexColumns[len(flexColumns)-1].Text
	if !yearRegexp.MatchString(year) {
		err := fmt.Errorf("error: year is not valid")
		return 0, err
	}

	yearInt, err := strconv.Atoi(year)
	if err != nil {
		err := fmt.Errorf("error: year is not valid")
		return 0, err
	}
	return yearInt, nil
}
