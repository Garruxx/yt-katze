package models

import "katze/src/models/artist"

type CarouseList struct {
	Items          []artist.CardItem `json:"albums,omitempty"`
	ContinuationID string            `json:"continuationId,omitempty"`
	ArtistID       string            `json:"artistId,omitempty"`
}
