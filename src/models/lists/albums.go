package lists

import (
	"katze/src/models"
	"katze/src/models/music"
)

type Albums struct {
	Albums       []music.Album       `json:"albums,omitempty"`
	Continuation models.Continuation `json:"continuation,omitempty"`
}
