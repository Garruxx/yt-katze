package services

import (
	"katze/src/services/browse"
	"katze/src/services/search"
)

// Search

// General
var SearchGeneral = search.General

// Music
var SearchMusicList = search.ItemsList
var SearchMusicPagination = search.ItemsPagination

// Artist
var SearchArtistList = search.ItemsList
var SearchArtistPagination = search.ItemsPagination

// Album
var SearchAlbumList = search.ItemsList
var SearchAlbumPagination = search.ItemsPagination

// Browse

// Artist
var BrowseArtistProfile = browse.ArtistProfile
var BrowseArtistMusicList = browse.ArtistMusicList
var BrowseArtistMusicPagination = browse.ArtistMusicListPagination
var BrowseArtistAlbums = browse.ArtistTracklist
var BrowseArtistSingles = browse.ArtistTracklist

// Album and Single
var BrowseAlbum = browse.Tracklist
var BrowseSingle = browse.Tracklist
