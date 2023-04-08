package shelves

import "katze/src/models/music"

type Music struct {
	Tracks         []music.Song `json:"tracks"`
	ContinuationID string       `json:"continuation_id,omitempty"`
	VisitorID      string       `json:"visitor_id,omitempty"`
}
