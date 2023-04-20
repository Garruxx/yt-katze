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

func TestArtistAlbums(t *testing.T) {

	// test case 1 - valid artistId, params and visitorID
	visitorID := "CgtpR05yLVQyVVo2TSitpoGiBg%3D%3D"
	artistId := "UCPC0L1d253x-KuMNwa05TpA"
	params := "6gPjAUdxY0JXcGdCQ3BVQkNpUjVkRjl3WVdkbFgzTnVZWEJ6YUc5MFgyMTFjMmxqWDNCaFoyVmZjbVZuYVc5dVlXd1NIMUY2UW5SMVdUVnJhamRKWWpSRU9EaFFkVGR1YUZwMmVuQlZOakJsVW1jYVRBQUFaVzRBQVVOUEFBRkRUd0FCQUVaRmJYVnphV05mWkdWMFlXbHNYMkZ5ZEdsemRBQUJBVU1BQUFFQUFRQUFBUUVBVlVOUVF6Qk1NV1F5TlRONExVdDFUVTUzWVRBMVZIQkJBQUh5MnJPcUNnWkFBVWdBVUNr"
	result, err := utils.GqlResolver[[]artist.CardItem](
		map[string]any{
			"artistId":  artistId,
			"params":    params,
			"visitorId": visitorID,
		},
		resolvers.ArtistAlbums,
	)
	if err != nil {
		t.Fatalf("Test case 1 failed: %v", err)
	}
	// verify that the result is not empty
	if len(result) == 0 {
		t.Errorf("Test case 1 failed: expected non-empty list of albums")
	}

	// test case 2 - invalid params
	visitorID = "CgtpR05yLVQyVVo2TSitpoGiBg%3D%3D"
	artistId = "UCPC0L1d253x-KuMNwa05TpA"
	params = "Tirilil"
	_, err = utils.GqlResolver[[]artist.CardItem](
		map[string]any{
			"artistId":  artistId,
			"params":    params,
			"visitorId": visitorID,
		},
		resolvers.ArtistAlbums,
	)
	if err == nil {
		t.Errorf("Test case 2 failed expected error but got nil: %v", err)
	}

	// test case 3 - invalid visitorID
	visitorID = "Tirilil"
	artistId = "UCPC0L1d253x-KuMNwa05TpA"
	params = "6gPjAUdxY0JXcGdCQ3BVQkNpUjVkRjl3WVdkbFgzTnVZWEJ6YUc5MFgyMTFjMmxqWDNCaFoyVmZjbVZuYVc5dVlXd1NIMUY2UW5SMVdUVnJhamRKWWpSRU9EaFFkVGR1YUZwMmVuQlZOakJsVW1jYVRBQUFaVzRBQVVOUEFBRkRUd0FCQUVaRmJYVnphV05mWkdWMFlXbHNYMkZ5ZEdsemRBQUJBVU1BQUFFQUFRQUFBUUVBVlVOUVF6Qk1NV1F5TlRONExVdDFUVTUzWVRBMVZIQkJBQUh5MnJPcUNnWkFBVWdBVUNr"
	_, err = utils.GqlResolver[[]artist.CardItem](
		map[string]any{
			"artistId":  artistId,
			"params":    params,
			"visitorId": visitorID,
		},
		resolvers.ArtistAlbums,
	)
	if err == nil {
		t.Fatalf("Test case 3 failed error: %v", err)
	}

	// test case 4 - invalid artistId type
	visitorID = "CgtpR05yLVQyVVo2TSitpoGiBg%3D%3D"
	params = ""
	_, err = utils.GqlResolver[[]artist.CardItem](
		map[string]any{
			"artistId":  4444,
			"params":    params,
			"visitorId": visitorID,
		},
		resolvers.ArtistAlbums,
	)
	if err == nil {
		t.Fatalf("Test case 4 failed error: %v", err)
	}

	// test case 5 - invalid params type
	visitorID = "CgtpR05yLVQyVVo2TSitpoGiBg%3D%3D"
	artistId = "UCPC0L1d253x-KuMNwa05TpA"
	_, err = utils.GqlResolver[[]artist.CardItem](
		map[string]any{
			"artistId":  artistId,
			"params":    4444,
			"visitorId": visitorID,
		},
		resolvers.ArtistAlbums,
	)
	if err == nil {
		t.Fatalf("Test case 5 failed error: %v", err)
	}

	// test case 6 - invalid visitorID type
	artistId = "UCPC0L1d253x-KuMNwa05TpA"
	_, err = utils.GqlResolver[[]artist.CardItem](
		map[string]any{
			"artistId":  artistId,
			"params":    params,
			"visitorId": 4444,
		},
		resolvers.ArtistAlbums,
	)
	if err == nil {
		t.Fatalf("Test case 6 failed error: %v", err)
	}
}
