package models

type FlexColumn struct {
	Text     string `json:"text"`
	BrowseID string `json:"browseId,omitempty"`
	WatchID  string `json:"watchId,omitempty"`
	PageType string `json:"pageType,omitempty"`
}
