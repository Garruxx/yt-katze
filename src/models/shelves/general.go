package shelves

import (
	"katze/src/models/lists"
	"katze/src/models/music"
)

type General struct {
	BestMatch music.BestMatch `json:"best_match,omitempty"`
	Tracks    lists.Music     `json:"tracks,omitempty"`
	Albums    lists.Albums    `json:"albums,omitempty"`
	Artists   lists.Artists   `json:"artists,omitempty"`
	VisitorID string          `json:"visitor_id,omitempty"`
}
