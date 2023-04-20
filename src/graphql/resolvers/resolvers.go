package resolvers

import (
	"katze/src/graphql/resolvers/browse"
	"katze/src/graphql/resolvers/search"
)

// Search - search resolver
var General = search.General
var MusicsList = search.MusicsList
var MusicsPagination = search.MusicsPagination
var AlbumsList = search.AlbumsList
var AlbumsPagination = search.AlbumsPagination
var ArtistList = search.ArtistsList
var ArtistsPagination = search.ArtistsPagination

// Browse - browse resolver
var Album = browse.Album
var Single = browse.Single
var ArtistAlbums = browse.ArtistAlbums
var ArtistMusicsList = browse.ArtistMusicsList
var ArtistMusicsListPagination = browse.ArtistMusicsListPagination
var ArtistProfile = browse.ArtistProfile
var ArtistSingles = browse.ArtistSingles
