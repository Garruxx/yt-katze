package artist

type SingleList struct {
	Singles            []CardItem `json:"albums,omitempty"`
	ContinuationParams string     `json:"params,omitempty"`
	ContinuationID     string     `json:"browseId,omitempty"`
}
