package artist

type AlbumsList struct {
	Albums             []CardItem `json:"albums,omitempty"`
	ContinuationParams string     `json:"params,omitempty"`
	ContinuationID     string     `json:"continuationId,omitempty"`
}
