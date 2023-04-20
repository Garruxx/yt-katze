package schema

import (
	"katze/src/graphql/schema/artist"
	"katze/src/graphql/schema/lists"
	"katze/src/graphql/schema/shelves"
)

// ArtistType is the GraphQL types for an artist.
var ArtistAlbumsList = artist.AlbumsListType
var ArtistSinglesList = artist.SinglesListType
var ArtistMusicList = artist.MusicsListType

// ShelvesType is the GraphQL types for a shelf.
var ShelfArtist = shelves.ArtistType
var ShelfAlbum = shelves.AlbumType
var ShelfGeneral = shelves.GeneralType
var ShelfSingle = shelves.SingleType

// ListsType is the GraphQL types for a list.
var AlbumsList = lists.AlbumType
var ArtistsList = lists.ArtistType
var MusicsList = lists.MusicType
