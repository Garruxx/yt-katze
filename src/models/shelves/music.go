package shelves

import "katze/src/models/music"

type Music struct {
	Songs          []music.Song `json:"songs"`
	ContinuationID string       `json:"continuationID,omitempty"`
	VisitorID      string       `json:"visitorID,omitempty"`
}
