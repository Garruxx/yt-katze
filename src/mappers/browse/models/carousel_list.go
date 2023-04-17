package models

import "katze/src/models/artist"

type CarouseList struct {
	Items              []artist.CardItem `json:"albums,omitempty"`
	ContinuationParams string            `json:"params,omitempty"`
	ContinuationID     string            `json:"browseId,omitempty"`
}
