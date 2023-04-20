package tests

import (
	"katze/src/graphql/resolvers"
	"katze/src/graphql/resolvers/tests/utils"
	"katze/src/models/artist"
	"testing"
)

/*
* The params and visitorID must be obtained at the time of the artist
* profile, as they expire after some time.
**/

func TestArtistSingles(t *testing.T) {

	// test case 1 - valid artistId, params and visitorID
	visitorID := "CgtpR05yLVQyVVo2TSitpoGiBg%3D%3D"
	artistId := "UCPC0L1d253x-KuMNwa05TpA"
	params := "6gPjAUdxY0JXcGdCQ3BVQkNpUjVkRjl3WVdkbFgzTnVZWEJ6YUc5MFgyMTFjMmxqWDNCaFoyVmZjbVZuYVc5dVlXd1NIMUY2UW5SMVdUVnJhamRKWWpSRU9EaFFkVGR1YUZwMmVuQlZOakJsVW1jYVRBQUFaVzRBQVVOUEFBRkRUd0FCQUVaRmJYVnphV05mWkdWMFlXbHNYMkZ5ZEdsemRBQUJBVU1BQUFFQUFRQUFBUUVBVlVOUVF6Qk1NV1F5TlRONExVdDFUVTUzWVRBMVZIQkJBQUh5MnJPcUNnWkFBa2dBVURr"
	result, err := utils.GqlResolver[[]artist.CardItem](
		map[string]any{
			"artistId":  artistId,
			"params":    params,
			"visitorId": visitorID,
		},
		resolvers.ArtistSingles,
	)
	if err != nil {
		t.Errorf("Test case 1 failed: %v", err)
	}
	// verify that the result is not empty
	if len(result) == 0 {
		t.Errorf("Test case 1 failed: expected non-empty list of songs")
	}

	// test case 2 - invalid params
	visitorID = "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D"
	artistId = "UCPC0L1d253x-KuMNwa05TpA"
	params = "Tirilil"
	_, err = utils.GqlResolver[[]artist.CardItem](
		map[string]any{
			"artistId":  artistId,
			"params":    params,
			"visitorId": visitorID,
		},
		resolvers.ArtistSingles,
	)
	if err == nil {
		t.Errorf("Test case 2 failed expected error but got nil: %v", err)
	}

	// test case 3 - invalid visitorID
	visitorID = "Tirilil"
	artistId = "UCPC0L1d253x-KuMNwa05TpA"
	params = "6gPjAUdxY0JXcGdCQ3BVQkNpUjVkRjl3WVdkbFgzTnVZWEJ6YUc5MFgyMTFjMmxqWDNCaFoyVmZjbVZuYVc5dVlXd1NIMUY1WjFNNWFubHFNMVpKWlRSRU9EaFFkVGR1YUZwemVqVm5SMEZsVW1jYVRBQUFaVzRBQVVOUEFBRkRUd0FCQUVaRmJYVnphV05mWkdWMFlXbHNYMkZ5ZEdsemRBQUJBVU1BQUFFQUFRQUFBUUVBVlVOUVF6Qk1NV1F5TlRONExVdDFUVTUzWVRBMVZIQkJBQUh5MnJPcUNnWkFBa2dBVURn"
	_, err = utils.GqlResolver[[]artist.CardItem](
		map[string]any{
			"artistId":  artistId,
			"params":    params,
			"visitorId": visitorID,
		},
		resolvers.ArtistSingles,
	)
	if err == nil {
		t.Errorf("Test case 3 failed expected error but got nil: %v", err)
	}

	// test case 4 - invalid artistId type
	visitorID = "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D"
	_, err = utils.GqlResolver[[]artist.CardItem](
		map[string]any{
			"artistId":  123,
			"params":    params,
			"visitorId": visitorID,
		},
		resolvers.ArtistSingles,
	)
	if err == nil {
		t.Errorf("Test case 4 failed expected error but got nil: %v", err)
	}

	// test case 5 - invalid params type
	visitorID = "CgtLT0txVmNPcFlnOCjMgvehBg%3D%3D"
	_, err = utils.GqlResolver[[]artist.CardItem](
		map[string]any{
			"artistId":  "UCPC0L1d253x-KuMNwa05TpA",
			"params":    123,
			"visitorId": visitorID,
		},
		resolvers.ArtistSingles,
	)
	if err == nil {
		t.Errorf("Test case 5 failed expected error but got nil: %v", err)
	}

	// test case 6 - invalid visitorID type
	artistId = "UCPC0L1d253x-KuMNwa05TpA"
	_, err = utils.GqlResolver[[]artist.CardItem](
		map[string]any{
			"artistId":  artistId,
			"params":    params,
			"visitorId": 123,
		},
		resolvers.ArtistSingles,
	)
	if err == nil {
		t.Errorf("Test case 6 failed expected error but got nil: %v", err)
	}
}
