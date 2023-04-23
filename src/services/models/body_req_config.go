package models

type BodyReqConfig struct {
	Params   string `json:"params,omitempty"`
	Query    string `json:"query,omitempty"`
	BrowseID string `json:"browseId,omitempty"`
	VideoID  string `json:"videoId,omitempty"`
}
