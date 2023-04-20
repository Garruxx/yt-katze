package search

import (
	"katze/src/mappers/search/albums"
	"katze/src/mappers/search/artists"
	"katze/src/mappers/search/general"
	"katze/src/mappers/search/music"
)

var General = general.Result
var MusicList = music.List
var MusicPagination = music.Pagination
var ArtistList = artists.List
var ArtistPagination = artists.Pagination
var AlbumList = albums.List
var AlbumPagination = albums.Pagination

