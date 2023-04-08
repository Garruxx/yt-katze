package tests
import (
	"katze/src/mappers/utils/columns"
	"katze/src/mappers/utils/columns/tests/data"
	"testing"
)

func TestGetDuration(t *testing.T) {

	// Test case 1 Valid duration
	flexColumn := data.FLEX_COLUMN_VALID_DURATION
	result, err := columns.GetDuration(flexColumn)
	if err != nil {
		t.Errorf("failed test 1 error: %v", err)
	}
	// Test case 1.1: result duration is equal to expected "20:15:01"
	if result != "20:15:01" {
		t.Errorf(
			"failed test 1.1 error: invalid duration expected: %v, got: %v",
			"20:15:01", result,
		)
	}

	// Test case 2 Invalid duration
	flexColumn = data.FLEX_COLUMN_INVALID_DURATION
	result, err = columns.GetDuration(flexColumn)
	if err == nil {
		t.Errorf("failed test 2 error: expected error, got: %v", err)
	}
	// Test case 2.1: result duration is equal to expected ""
	if result != "" {
		t.Errorf(
			"failed test 2.1 error: invalid duration expected: %s, got: %v",
			"empty", result,
		)
	}

	// Test case 3 invalid duration 2
	flexColumn = data.FLEX_COLUMN_INVALID_DURATION_2
	result, err = columns.GetDuration(flexColumn)
	if err == nil {
		t.Errorf("failed test 3 error: expected error, got: %v", err)
	}
	// Test case 3.1: result duration is equal to expected ""
	if result != "" {
		t.Errorf(
			"failed test 3.1 error: invalid duration expected: %s, got: %v",
			"empty", result,
		)
	}

	// Test case 4 invalid flex column empty
	flexColumn = data.FLEX_COLUMN_DURATION_EMPTY
	result, err = columns.GetDuration(flexColumn)
	if err == nil {
		t.Errorf("failed test 4 error: expected error, got: %v", err)
	}
	// Test case 4.1: result duration is equal to expected ""
	if result != "" {
		t.Errorf(
			"failed test 4.1 error: invalid duration expected: %d got: %v",
			0, result,
		)
	}

}
