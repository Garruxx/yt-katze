package lists

import (
	"katze/src/models"
	"katze/src/models/music"
)

type Artists struct {
	Artists      []music.Artist      `json:"artists"`
	Continuation models.Continuation `json:"continuation"`
	
}
