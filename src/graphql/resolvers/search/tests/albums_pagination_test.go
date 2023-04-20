package tests

import (
	"katze/src/graphql/resolvers"
	"katze/src/graphql/resolvers/tests/utils"
	"katze/src/models/lists"
	"testing"
)

func TestAlbumsPagination(t *testing.T) {
	// Test case 1 - valid continuationID and visitorID
	continuationID := "EvUFEgRhbW9yGuwFRWdXS0FRSVlBVWdVYWdvUUF4QUVFQWtRQ2hBRmdnRVlWVU10Y1ROcU56aEdRMFF4VGpVeVF6SjZhR1ozU0VOQmdnRVlWVU5NWWpCQ1NYaFpVMHBLUjNkQlRFSnJRakJxVmpGUmdnRVlWVU5RVWtONmQxcDJVVkkyUkZSM016WnllWEZPVTBKM2dnRVlWVU4zTVd4VVpHeHhiMkV5T1dwMU9VSjVUR0kzZDFkQmdnRVlWVU5tVWpReE4zSlpWMGN4Y21kUGIzUk5kaTFyV1RCM2dnRVlWVU01ZFhRM01IZzJVVlpVWmpWaGNtZHRZMFZxTVd4M2dnRVlWVU01VEZVeFVtSkhjR0U0WWtOc2VUTlhaMkU1VTBWUmdnRVlWVU4zVlZaVGN6WnJTM0kwUzNKdGVWWkZiRE5qVVZWbmdnRVlWVU10ZUY5eGVtWkdSRkJSYUVWV05XTkhSRTlhWVd0QmdnRVlWVU5WUTJObU1GRktSMmd6VkcxRlZsZGhURFZNT1Y5bmdnRVlWVU5QY2pCMFQxWXlhRWxLWDA5TVh6VTViMFZNYTA1M2dnRVlWVU5yTmxkdmJtaEpiRFpEYnpZd1ltUklPRUZUZFdSM2dnRVlWVU4zU1RSV1pFVXplbkpXTnpCUlp5MUNjbk5ZTFRWbmdnRVlWVU5mV25wVmJXcFVRblpxZG1kS1IxQjJkMUF6YkhsQmdnRVlWVU53VjFGeVNWQjVSMmN0WWpjMWEwWnNlazlNVFRoM2dnRVlWVU51VDNsdFVtMXRkRTQ1UVZORmVVbEhjRkpmTjFsM2dnRVlWVU51ZVhGME9YaDVTa2RmU1U1VWJESkhRamhtWkhCM2dnRVlWVU5mV1V4VmRHaG1NMlZtZGtSU01HWlBVbm8yVmw5M2dnRVlWVU5sVVVNMlRUSjBZMDl0VGpWMk5FMXpXVmxRUjFSQmdnRVlWVU5JVjNKb1Mxa3lWbEZHU210bE9HcFliblZaTjBsUhjx6tAu"
	visitorID := "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D"
	result, err := utils.GqlResolver[lists.Albums](
		map[string]any{
			"continuationId": continuationID,
			"visitorId":      visitorID,
		},
		resolvers.AlbumsPagination,
	)
	if err != nil {
		t.Fatalf("Test case 1 failed: %v", err)
	}
	// Test case 1.1 verify that the result is not nil
	if len(result.Albums) == 0 {
		t.Errorf("Test case 1.1 failed: expected non-nil album list")
	}

	// Test case 2 - invalid continuationID
	continuationID = "Tirilil"
	visitorID = "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D"
	result, err = utils.GqlResolver[lists.Albums](
		map[string]any{
			"continuationId": continuationID,
			"visitorId":      visitorID,
		},
		resolvers.AlbumsPagination,
	)
	if err == nil {
		t.Fatalf("Test case 2 failed expected error but got nil: %v", err)
	}

	// Test case 3 - invalid visitorID
	continuationID = "EvUFEgRhbW9yGuwFRWdXS0FRSVlBVWdVYWdvUUF4QUVFQWtRQ2hBRmdnRVlWVU10Y1ROcU56aEdRMFF4VGpVeVF6SjZhR1ozU0VOQmdnRVlWVU5NWWpCQ1NYaFpVMHBLUjNkQlRFSnJRakJxVmpGUmdnRVlWVU5RVWtONmQxcDJVVkkyUkZSM016WnllWEZPVTBKM2dnRVlWVU4zTVd4VVpHeHhiMkV5T1dwMU9VSjVUR0kzZDFkQmdnRVlWVU5tVWpReE4zSlpWMGN4Y21kUGIzUk5kaTFyV1RCM2dnRVlWVU01ZFhRM01IZzJVVlpVWmpWaGNtZHRZMFZxTVd4M2dnRVlWVU01VEZVeFVtSkhjR0U0WWtOc2VUTlhaMkU1VTBWUmdnRVlWVU4zVlZaVGN6WnJTM0kwUzNKdGVWWkZiRE5qVVZWbmdnRVlWVU10ZUY5eGVtWkdSRkJSYUVWV05XTkhSRTlhWVd0QmdnRVlWVU5WUTJObU1GRktSMmd6VkcxRlZsZGhURFZNT1Y5bmdnRVlWVU5QY2pCMFQxWXlhRWxLWDA5TVh6VTViMFZNYTA1M2dnRVlWVU5yTmxkdmJtaEpiRFpEYnpZd1ltUklPRUZUZFdSM2dnRVlWVU4zU1RSV1pFVXplbkpXTnpCUlp5MUNjbk5ZTFRWbmdnRVlWVU5mV25wVmJXcFVRblpxZG1kS1IxQjJkMUF6YkhsQmdnRVlWVU53VjFGeVNWQjVSMmN0WWpjMWEwWnNlazlNVFRoM2dnRVlWVU51VDNsdFVtMXRkRTQ1UVZORmVVbEhjRkpmTjFsM2dnRVlWVU51ZVhGME9YaDVTa2RmU1U1VWJESkhRamhtWkhCM2dnRVlWVU5mV1V4VmRHaG1NMlZtZGtSU01HWlBVbm8yVmw5M2dnRVlWVU5sVVVNMlRUSjBZMDl0VGpWMk5FMXpXVmxRUjFSQmdnRVlWVU5JVjNKb1Mxa3lWbEZHU210bE9HcFliblZaTjBsUhjx6tAu"
	visitorID = "Tirilil"
	result, err = utils.GqlResolver[lists.Albums](
		map[string]any{
			"continuationId": continuationID,
			"visitorId":      visitorID,
		},
		resolvers.AlbumsPagination,
	)
	if err != nil {
		t.Fatalf("Test case 3 failed expected error but got nil: %v", err)
	}
	// Test case 3.1 verify that the result is no nil
	if len(result.Albums) == 0 {
		t.Errorf("Test case 3.1 failed: expected album list")
	}

	// Test case 4 - invalid visitorID type
	_, err = utils.GqlResolver[lists.Albums](
		map[string]any{
			"continuationId": continuationID,
			"visitorId":      123,
		},
		resolvers.AlbumsPagination,
	)
	if err == nil {
		t.Errorf("Test case 4 failed expected error but got nil: %v", err)
	}

	// Test case 5 - invalid continuationID type
	_, err = utils.GqlResolver[lists.Albums](
		map[string]any{
			"continuationId": 123,
			"visitorId":      visitorID,
		},
		resolvers.AlbumsPagination,
	)
	if err == nil {
		t.Errorf("Test case 5 failed expected error but got nil: %v", err)
	}
}
