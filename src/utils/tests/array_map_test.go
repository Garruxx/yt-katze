package tests

import (
	"katze/src/utils"
	"testing"
)

func TestArrayMap(t *testing.T) {

	// Test case 1 valid case
	arr := []int{1, 2, 3}
	f := func(x int) (int, error) {
		return x * 2, nil
	}
	expected := []int{2, 4, 6}
	result, err := utils.ArrayMap(arr, f)
	if err != nil {
		t.Errorf("test 1 failed: %v", err)
	}
	// Test case 1.1 check if the array has the correct values
	for i, x := range result {
		if x != expected[i] {
			t.Errorf(
				"test 1.1 failed, error expected %d, got: %d", expected[i], x,
			)
		}
	}

	// Test case 2 invalid case
	arr = []int{}
	result, err = utils.ArrayMap(arr, f)
	if err != nil {
		t.Errorf("test 2 failed: %v", err)
	}
	// Test case 2.1 check if the array is empty
	if len(result) != 0 {
		t.Errorf("test 2.1 failed error expected empty array, got: %v", result)
	}

}
