package artist

type SinglesList struct {
	Singles        []CardItem `json:"singles,omitempty"`
	ContinuationID string     `json:"continuationId,omitempty"`
	ArtistID       string     `json:"artistId,omitempty"`
}
