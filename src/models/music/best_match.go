package music

import "katze/src/models"

type BestMatch struct {
	Type       string             `json:"type,omitempty"`
	AlbumType  string             `json:"albumType,omitempty"`
	Title      string             `json:"title,omitempty"` 
	ID         string             `json:"id,omitempty"`
	Thumbnails []models.Thumbnail `json:"thumbnail,omitempty"`
	WatchID    string             `json:"watchId,omitempty"`
	Artists    []Artist           `json:"artist,omitempty"`
	AlbumTitle string             `json:"album,omitempty"`
	AlbumID    string             `json:"album_id,omitempty"`
	Duration   string             `json:"duration,omitempty"`
	Explicit   bool               `json:"explicit,omitempty"`
}
