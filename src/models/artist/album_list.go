package artist

type AlbumList struct {
	Albums             []CardItem `json:"albums,omitempty"`
	ContinuationParams string     `json:"params,omitempty"`
	ContinuationID     string     `json:"browseId,omitempty"`
}
