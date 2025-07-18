package lists

import (
	"katze/src/models"
	"katze/src/models/music"
)

type Artists struct {
	Artists        []music.Artist      `json:"artists,omitempty"`
	Continuation   models.Continuation `json:"continuation,omitempty"`
	ContinuationID string              `json:"continuationId,omitempty"`
	VisitorID      string              `json:"visitorId,omitempty"`
}
