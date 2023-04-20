package browse

import (
	"katze/src/mappers/browse/artist"
	"katze/src/mappers/browse/tracklist"
)

// Artist is a mapper for the artist page
var ArtistProfile = artist.Profile
var ArtistMusicList = artist.MusicList
var ArtistMusicPagination = artist.MusicPagination
var ArtistAlbums = artist.Albums
var ArtistSingles = artist.Singles

var Album = tracklist.Map
var Single = tracklist.Map
