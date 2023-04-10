package mappers

import (
	"katze/src/mappers/utils/mappers/artist"
	"katze/src/mappers/utils/mappers/artists"
	"katze/src/mappers/utils/mappers/song"
	"katze/src/mappers/utils/mappers/songs"
)

var Song = song.Mapper
var Songs = songs.Mapper
var Artist = artist.Mapper
var Artists = artists.Mapper
