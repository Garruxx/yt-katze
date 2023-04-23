package mappers

import (
	"katze/src/mappers/browse"
	"katze/src/mappers/player"
	"katze/src/mappers/search"
)

// Browse
// Artist
var BrowseArtistProfile = browse.ArtistProfile
var BrowseArtistMusicList = browse.ArtistMusicList
var BrowseArtistMusicPagination = browse.ArtistMusicPagination
var BrowseArtistAlbums = browse.ArtistAlbums
var BrowseArtistSingles = browse.ArtistSingles
var BeowseRecomendations = browse.Recomendations
// Album and Single
var BrowseAlbum = browse.Album
var BrowseSingle = browse.Single

// Search
// General
var SearchGeneral = search.General

// Music
var SearchMusicList = search.MusicList
var SearchMusicPagination = search.MusicPagination

// Artist
var SearchArtistList = search.ArtistList
var SearchArtistPagination = search.ArtistPagination

// Album
var SearchAlbumList = search.AlbumList
var SearchAlbumPagination = search.AlbumPagination

// Player
var PlayerInformation = player.Information
