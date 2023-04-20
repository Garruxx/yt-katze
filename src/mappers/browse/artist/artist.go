package artist

import (
	"katze/src/mappers/browse/artist/card"
	"katze/src/mappers/browse/artist/profile"
	"katze/src/mappers/browse/artist/music"
)

var Profile = profile.Map
var Albums = card.TwoRowItem
var Singles = card.TwoRowItem
var MusicList = music.List
var MusicPagination = music.Pagination
