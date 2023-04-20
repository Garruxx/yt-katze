package artist

type SinglesList struct {
	Singles            []CardItem `json:"singles,omitempty"`
	ContinuationParams string     `json:"params,omitempty"`
	ContinuationID     string     `json:"browseId,omitempty"`
}
