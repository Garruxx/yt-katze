package artist

import "katze/src/models/music"

type MusicList struct {
	Songs    []music.Song `json:"songs,omitempty"`
	ContinuationParams string     `json:"params,omitempty"`
	ContinuationID     string     `json:"browseId,omitempty"`
}
