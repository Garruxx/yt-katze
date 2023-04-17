package utils

// ArrayMap applies a function to each element
//of an array and returns the result
func ArrayMap[T, U any](arr []T, f func(T) (U, error)) ([]U, error) {

	if len(arr) == 0 {
		return []U{}, nil
	}
	result := make([]U, len(arr))
	for i, x := range arr {
		item, err := f(x)
		if err != nil {
			return nil, err
		}
		result[i] = item
	}
	return result, nil
}
