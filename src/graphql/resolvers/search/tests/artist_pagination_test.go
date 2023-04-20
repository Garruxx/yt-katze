package tests

import (
	"katze/src/graphql/resolvers"
	"katze/src/graphql/resolvers/tests/utils"
	"katze/src/models/lists"
	"testing"
)

func TestArtistPagination(t *testing.T) {

	// Test case 1 - valid continuationID and visitorID
	continuationID := "EvQFEgNzaGUa7AVFZ1dLQVFJZ0FVZ1VhZ29RQXhBRUVBa1FDaEFGZ2dFWVZVTkhTRFpzTFVodmVtcHpPSFpUVnpnelpUUjFMVWQzZ2dFWVZVTkZSRlpmTjI5cFVXOUlUSEZZZUdSTlpqbHlUbVJCZ2dFWVZVTm9TR2RsWkVSeWFtaFVUMHBMTldKdlYxUnphelYzZ2dFWVZVTjRSSG94WDBkaVdXNDNNMjVQZFZNd01WVlJMVUYzZ2dFWVZVTTRRM1puVFc5bFpsa3dZazU1TlRKeGN6QmpVRkJuZ2dFWVZVTm9YemRvUVVOVmIydEViR3BWT0Zoc2JqWXhOSGxuZ2dFWVZVTnJSRkV0VFhaVVdXcElhbVZHWVVsUlNuZzBaVXRSZ2dFWVZVTkpiRVV5ZFVsWFVXWXRWbGRTVjFaaU1TMTVkRkZuZ2dFWVZVTXRaWHA2TVd3MU5rZEZWbUZ3TFVWNlVYQlRVbkYzZ2dFWVZVTTNNMjVKYmtwV1JIQkJXSEZ5VUVGa05GbzRkblIzZ2dFWVZVTkNhV1o2V0VnMlkxSnRUazFKWkZWVVVrdDBjek5SZ2dFWVZVTnliVGR1TW01QlRtOTJkVmRYZFdwUGMxVnBVbXhuZ2dFWVZVTjRkVFl4YlVSa1dFZGxTWGxJVVV4cFJVaDNiRk4zZ2dFWVZVTlZWMVJaWnkxb2EyZEdRVXhzYlY5bVdtSmFaSE4zZ2dFWVZVTXdjbmhoUTNVeWJUUXlPVWRvWTB3M1RteERkbkIzZ2dFWVZVTmZVbU5wZUdkSWFFWk1ZbEkxU3pac1QxQlJla05uZ2dFWVZVTnJiR3BCWTBWdU1WWnFaRkZxTjJwVWEzaENWaTFuZ2dFWVZVTkplamhuWjBGd05HWkZaazl4VkdjeVJWTk5SbXRCZ2dFWVZVTjBMVkp4VlRod2JUVmphVXQ0ZGtGUVluSTRTa1puZ2dFWVZVTjNVbVI1Tkd4SE16UTBjblozUVZCMUxVWTFTMVYzGPHq0C4"
	visitorID := "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D"
	result, err := utils.GqlResolver[lists.Artists](
		map[string]any{
			"continuationId": continuationID,
			"visitorId":      visitorID,
		},
		resolvers.ArtistsPagination,
	)
	if err != nil {
		t.Fatalf("Test case 1 failed: %v", err)
	}
	// Test case 1.1 verify that the result is not nil
	if len(result.Artists) == 0 {
		t.Fatalf("Test case 1.1 failed: expected non-nil artist list")
	}

	// Test case 2 - invalid continuationID
	continuationID = "Tirilil"
	visitorID = "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D"
	result, err = utils.GqlResolver[lists.Artists](
		map[string]any{
			"continuationId": continuationID,
			"visitorId":      visitorID,
		},
		resolvers.ArtistsPagination,
	)
	if err == nil {
		t.Errorf("Test case 2 failed expected an error, got: %v", err)
	}
	// Test case 2.1 verify that the result is nil
	if len(result.Artists) != 0 {
		t.Fatalf("Test case 2.1 failed: expected nil artist list")
	}

	// Test case 3 - invalid visitorID
	visitorID = "Tirilil"
	continuationID = "EvQFEgNzaGUa7AVFZ1dLQVFJZ0FVZ1VhZ29RQXhBRUVBa1FDaEFGZ2dFWVZVTkhTRFpzTFVodmVtcHpPSFpUVnpnelpUUjFMVWQzZ2dFWVZVTkZSRlpmTjI5cFVXOUlUSEZZZUdSTlpqbHlUbVJCZ2dFWVZVTm9TR2RsWkVSeWFtaFVUMHBMTldKdlYxUnphelYzZ2dFWVZVTjRSSG94WDBkaVdXNDNNMjVQZFZNd01WVlJMVUYzZ2dFWVZVTTRRM1puVFc5bFpsa3dZazU1TlRKeGN6QmpVRkJuZ2dFWVZVTm9YemRvUVVOVmIydEViR3BWT0Zoc2JqWXhOSGxuZ2dFWVZVTnJSRkV0VFhaVVdXcElhbVZHWVVsUlNuZzBaVXRSZ2dFWVZVTkpiRVV5ZFVsWFVXWXRWbGRTVjFaaU1TMTVkRkZuZ2dFWVZVTXRaWHA2TVd3MU5rZEZWbUZ3TFVWNlVYQlRVbkYzZ2dFWVZVTTNNMjVKYmtwV1JIQkJXSEZ5VUVGa05GbzRkblIzZ2dFWVZVTkNhV1o2V0VnMlkxSnRUazFKWkZWVVVrdDBjek5SZ2dFWVZVTnliVGR1TW01QlRtOTJkVmRYZFdwUGMxVnBVbXhuZ2dFWVZVTjRkVFl4YlVSa1dFZGxTWGxJVVV4cFJVaDNiRk4zZ2dFWVZVTlZWMVJaWnkxb2EyZEdRVXhzYlY5bVdtSmFaSE4zZ2dFWVZVTXdjbmhoUTNVeWJUUXlPVWRvWTB3M1RteERkbkIzZ2dFWVZVTmZVbU5wZUdkSWFFWk1ZbEkxU3pac1QxQlJla05uZ2dFWVZVTnJiR3BCWTBWdU1WWnFaRkZxTjJwVWEzaENWaTFuZ2dFWVZVTkplamhuWjBGd05HWkZaazl4VkdjeVJWTk5SbXRCZ2dFWVZVTjBMVkp4VlRod2JUVmphVXQ0ZGtGUVluSTRTa1puZ2dFWVZVTjNVbVI1Tkd4SE16UTBjblozUVZCMUxVWTFTMVYzGPHq0C4"
	result, err = utils.GqlResolver[lists.Artists](
		map[string]any{
			"continuationId": continuationID,
			"visitorId":      visitorID,
		},
		resolvers.ArtistsPagination,
	)
	if err != nil {
		t.Fatalf("Test case 3 failed error: %v", err)
	}
	// Test case 3.1 verify that the result is not nil
	if len(result.Artists) == 0 {
		t.Fatalf("Test case 3.1 failed: expected non-nil artist list")
	}

	// Test case 4 - invalid visitorID type
	result, err = utils.GqlResolver[lists.Artists](
		map[string]any{
			"continuationId": continuationID,
			"visitorId":      123,
		},
		resolvers.ArtistsPagination,
	)
	if err == nil {
		t.Errorf("Test case 4 failed expected an error, got: %v", err)
	}

	// Test case 5 - invalid continuationID type
	result, err = utils.GqlResolver[lists.Artists](
		map[string]any{
			"continuationId": 123,
			"visitorId":      visitorID,
		},
		resolvers.ArtistsPagination,
	)
	if err == nil {
		t.Errorf("Test case 5 failed expected an error, got: %v", err)
	}
}
