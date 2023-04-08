package lists

import (
	"katze/src/models"
	"katze/src/models/music"
)

type Music struct {
	Songs        []music.Song        `json:"songs"`
	Continuation models.Continuation `json:"continuation"`
}
