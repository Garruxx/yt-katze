package artist

type AlbumsList struct {
	Albums         []CardItem `json:"albums,omitempty"`
	ContinuationID string     `json:"continuationId,omitempty"`
	ArtistID       string     `json:"ArtistId,omitempty"`
}
