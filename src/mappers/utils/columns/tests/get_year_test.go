package tests

import (
	"katze/src/mappers/utils/columns"
	"katze/src/mappers/utils/columns/tests/data"
	"testing"
)

func TestGetYear(t *testing.T) {

	// Test case 1 Valid year
	flexColumn := data.FLEX_COLUMN_VALID_YEAR
	result, err := columns.GetYear(flexColumn)
	if err != nil {
		t.Errorf("failed test 1 error: %v", err)
	}
	// Test case 1.1: result year is equal to expected "2022"
	if result != 2022 {
		t.Errorf(
			"failed test 1.1 error: invalid year expected: %v, got: %v",
			"2022", result,
		)
	}

	// Test case 2 Invalid year
	flexColumn = data.FLEX_COLUMN_INVALID_YEAR
	result, err = columns.GetYear(flexColumn)
	if err == nil {
		t.Errorf("failed test 2 error: expected error, got: %v", err)
	}
	// Test case 2.1: result year is equal to expected "0"
	if result != 0 {
		t.Errorf(
			"failed test 2.1 error: invalid year expected: %d, got: %v",
			0, result,
		)
	}

	// Test case 3 invalid year 2
	flexColumn = data.FLEX_COLUMN_INVALID_YEAR_2
	result, err = columns.GetYear(flexColumn)
	if err == nil {
		t.Errorf("failed test 3 error: expected error, got: %v", err)
	}
	// Test case 3.1: result year is equal to expected "0"
	if result != 0 {
		t.Errorf(
			"failed test 3.1 error: invalid year expected: %d, got: %v",
			0, result,
		)
	}

	// Test case 4 invalid flex column empty
	flexColumn = data.FLEX_COLUMN_EMPTY
	result, err = columns.GetYear(flexColumn)
	if err == nil {
		t.Errorf("failed test 4 error: expected error, got: %v", err)
	}
	// Test case 4.1: result year is equal to expected "0"
	if result != 0 {
		t.Errorf(
			"failed test 4.1 error: invalid year expected: %d got: %v",
			0, result,
		)
	}

}
