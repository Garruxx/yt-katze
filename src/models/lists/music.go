package lists

import (
	"katze/src/models"
	"katze/src/models/music"
)

type Music struct {
	Songs          []music.Song        `json:"songs,omitempty"`
	Continuation   models.Continuation `json:"continuation,omitempty"`
	ContinuationID string              `json:"continuationId,omitempty"`
	VisitorID      string              `json:"visitorId,omitempty"`
}
