# Katze-YT

![YT-KATZE logo](./assets/logo.png)

A [GraphQL](https://graphql.org) API made in [Go](https://go.dev) for clean access to the [YouTube Music](https://music.youtube.com) API.
**Technologies used**

- [![Golang logo](./assets/go.png)Golang](https://go.dev)
- [![GraphQL logo](./assets/graphql.png)GraphQL](https://graphql.org)

**_No authentication required_**

<<<<<<< HEAD
## Usage [YouTube Music](https://music.youtube.com)
=======
## Uso
>>>>>>> aec7c35fa43a1eda5954b57f4e55425e38be48cb

- [General search result](#general-search)
- [Songs search](#search-for-songs)
- [Songs-pagination](#songs-pagination)
- [Albums-search](#search-for-albums)
- [Pagination of albums](#albums-pagination)
- [Search for artists](#artist-search)
- [Pagination of artists](#artist-pagination)
- [Artist-profile](#artist-singles)
- [Artist's songs](#songs-of-the-artist)
- [Artist's song pagination](#artist-song-pagination).
- [Artist albums](#artists-albums)
- [Artist's singles](#artist-singles)
- [Album-content](#album-content)
- [Single-content](#contents-of-the-single)

## Type sets:

- [common](#common-types)
  They are shared types, these are used by other more complex types such as those mentioned above.
  They contain primitive types such as `album` or `song`.

  - [thumbnails](#thumbnails)
  - [song](#song)
  - [single](#single)
  - [bestMatch](#bestmatch)
  - [artist](#artist)
  - [album](#album)
  - [continuation](#continuation)
  - [cardItem](#carditem)

- [Artist](#artist-types)
  These are the types related to an artist, such as their track list, albums and singles, they contain:
  an array of data of some other type, params and continuationId, params and continuationId.
  of some other type, params and continuationId

  - [artistAlbumsList](#artistalbumslist)
  - [artistSinglesList](#artistsingleslist)
  - [artistMusicsList](#artistmusicslist)

- [Lists](#lists-types)
  They are lists of other types, but these include continuation, therefore they have an array of data of some other type
  accompanied by a [continuation](#continuation) field, continuationId and a visitorId

  - [musicsList](#musicslist)
  - [albumsList](#albumslist)
  - [artistsList](#artistslist)

- [Shelves](#shelves-types)
  These are shelves, they include complex types, for example: An album with its songs, images and description,
  an artist with its songs, albums, singles, and description, a single with its song, images and description,
  the general search with its songs, albums and artists.

  - [searchShelf](#searchshelf)
  - [artistShelf](#artistshelf)
  - [albumShelf](#albumshelf)

## Usage

### General search

In the query set the query and in visitorId the visitor id.

> **If the visitor id is not set, one will be generated automatically**.  
> Remember to store it for future reference.  
> You will not be able to get the albums of an artist, nor the singles of an artist without a valid visitor id that has been previously used to request the artist's profile.
> that has been previously used to request the artist's profile.

> The bestMatch is the best result of the search, it can be `song`, `album` or `artist`.
>
> - if it is a `song` **`Will not contain`** `albumType`.
> - if `album` or `single` **`Will not contain`** `album` or `albumId` the album title will be `title` and the type will be specified in `albumType`.
> - if it is a video it will return song in type, but **`will not contain`** `albumType`, `album` or `albumId`.
> - if it is an artist **`Will not contain`** `albumType`, `album` or `albumId`.

```graphql
{
  general(visitorId: "", query: "") {
    bestMatch {
      title
      id
      type
      album
      albumId
      albumType
      artists {
        id
        name
        thumbnails {
          height
          width
          url
        }
      }

      duration
      explicit
      thumbnails {
        width
        height
        url
      }
      watchId
    }
    tracks {
      continuation {
        params
        query
      }
      continuationId
      songs {
        album
        albumId
        thumbnails {
          width
          height
          url
        }
        artists {
          id
          name
        }
        duration
        explicit
        title
        id
      }
    }
    albums {
      albums {
        artists {
          id
          name
        }
        ep
        explicit
        id
        single
        thumbnails {
          height
          height
          width
          url
        }
        year
        title
      }
      continuation {
        query
        params
      }
      continuationId
    }
    artists {
      artists {
        id
        name
        thumbnails {
          width
          height
          url
        }
      }
      continuation {
        params
        query
      }
    }
    visitorId
  }
}
```

### Search for songs

> **If the visitor id is not set, one will be generated automatically.** ** **Remember to store it for future reference.
> Remember to store it for future queries.
> `params` is very important, this is who really indicates what we are looking for.
> it is obtained in the `general` query in the `tracks.continuation.params` field > `query` can be changed or obtained from the `tracks.continuation.query` field.

```graphql
{
  musicsList(visitorId: "", query: "", params: "") {
    continuationId
    visitorId
    songs {
      title
      id
      album
      albumId
      thumbnails {
        width
        height
        url
      }
      artists {
        id
        name
      }
      duration
      explicit
    }
  }
}
```

### Songs pagination

> **If the visitor id is not set, one will be generated automatically.** ** **If the visitor id is not set, one will be generated automatically.  
> Remember to store it for future queries.  
> `continuationId` is very important, this is who really indicates what we are looking for.  
> is obtained in the query `musicsList` in the `continuationId` field.

```graphql
{
  musicsPagination(visitorId: "", continuationId: "") {
    continuationId
    visitorId
    songs {
      title
      id
      album
      albumId
      thumbnails {
        width
        height
        url
      }
      artists {
        id
        name
      }
      duration
      explicit
    }
  }
}
```

### Search for albums

> **If the visitor id is not set, one will be generated automatically**.  
> Remember to store it for future queries.  
> `params` is very important, this is who really indicates what we are looking for.
> it is obtained in the `general` query in the `albums.continuation.params` field > `query` can be changed or obtained from the `albums.continuation.query` field.

```graphql
{
  albumsList(visitorId: "", query: "", params: "") {
    continuationId
    visitorId
    albums {
      artists {
        id
        name
      }
      ep
      explicit
      id
      single
      thumbnails {
        height
        width
        url
      }
      year
      title
    }
  }
}
```

### Albums pagination

> **If the visitor id is not set, one will be generated automatically.** ** **If the visitor id is not set, one will be generated automatically.
> Remember to store it for future queries.  
> `continuationId` is very important, this is who really indicates what we are looking for.
> is obtained in the query `albumsList` in the `continuationId` field.

```graphql
{
  albumsPagination(visitorId: "", continuationId: "") {
    continuationId
    visitorId
    albums {
      artists {
        id
        name
      }
      ep
      explicit
      id
      single
      thumbnails {
        height
        width
        url
      }
      year
      title
    }
  }
}
```

### Artist search

> **If the visitor id is not set, one will be generated automatically.** ** **If the visitor id is not set, one will be generated automatically.
> Remember to store it for future queries.  
> `params` is very important, this is who really indicates what we are looking for.
> `params` is obtained in the `general` query in the `artists.continuation.params` field.
> and the query can be changed or obtained from there.

```graphql
{
  artistsList(visitorId: "", query: "", params: "") {
    continuationId
    visitorId
    artists {
      id
      name
      thumbnails {
        width
        height
        url
      }
    }
  }
}
```

### Artist pagination

> **If the visitor id is not set, one will be generated automatically.** ** **If the visitor id is not set, one will be generated automatically.  
> Remember to store it for future queries.  
> `continuationId` is very important, this is who really indicates what we are looking for.
> `continuationId` is obtained in the query `artistsList` in the `continuationId` field.

```graphql
{
  artistsPagination(visitorId: "", continuationId: "") {
    continuationId
    visitorId
    artists {
      id
      name
      thumbnails {
        width
        height
        url
      }
    }
  }
}
```

### Artist profile

> **If the visitor id is not set, one will be generated automatically**.  
> Remember to store it for future queries.
> `artistId` is found in all results containing an artist card.
> for example in the query `musicsList` in the `songs.artists.id` field.

```graphql
{
  artistProfile(visitorId: "", artistId: "") {
    name
    visitorId
    thumbnails {
      width
      height
      url
    }
    albumsList {
      albums {
        explicit
        id
        title
        year
        thumbnails {
          height
          width
          url
        }
      }
      artistId
      continuationId
    }
    musicsList {
      songs {
        title
        id
        duration
        album
        albumId
        explicit
        thumbnails {
          width
          height
          url
        }
        artists {
          id
          name
        }
        duration
        explicit
      }
      params
      continuationId
    }
    singlesList {
      singles {
        id
        explicit
        title
        year
        thumbnails {
          width
          height
          url
        }
      }
      artistId
      continuationId
    }
  }
}
```

### Songs of the artist

> The visitorId is optional, it allows you to customize the search results.  
> In this case **no visitorId** is generated.
> `params` and continuation id are found in the musicsList field of the [Artist-profile](#artist-profile)

```graphql
    artistMusicsList(visitorId: "", continuationId: "", params: "") {
      songs {
        title
        id
        duration
        album
        albumId
        explicit
        thumbnails {
          width
          height
          url
        }
        artists {
          id
          name
        }
        duration
        explicit
      }
      continuationId
    }
```

### Artist song pagination

> The visitorId is optional, it allows you to customize the search results.  
> In this case **no visitorId** is generated.
> `continuationId` is found in [Artist's songs](#artist's-songs)

```graphql
    artistMusicsPagination(visitorId: "", continuationId: "") {
      songs {
        title
        id
        duration
        album
        albumId
        explicit
        duration
        explicit
        thumbnails {
          width
          height
          url
        }
        artists {
          id
          name
        }
      }
      params
      continuationId
    }
```

### Artist's albums

> The visitorId is mandatory, it allows you to customize the search results.  
> **There will be NO RESULTS IF THE visitorId IS NOT SUBMITTED**     
> `artistId` and `continuationId`
> are in the albumsList field of [artistProfile](#artist-profile).

```graphql
    artistAlbumsList(visitorId: "", continuationId: "", params: "") {
      albums {
        explicit
        id
        params
        title
        year
        thumbnails {
          height
          width
          url
        }
      }
    }
```

### Artist singles

> The visitorId is mandatory, it allows you to customize the search results.  
> **There will be NO RESULTS IF THE visitorId IS NOT SUBMITTED** > `artistId` and `continuationId`  
> are in the siglesList field of [artistProfile](#artist-profile).

```graphql
    artistSinglesList(visitorId: "", artistId: "", continuationId: "") {
      singles {
        id
        explicit
        params
        title
        year
        thumbnails {
          width
          height
          url
        }
      }
    }
```

### Album content

> **If the visitorId is not sent, one will be generated automatically**.
> Remember to store it for future reference.  
> `albumId` Found in all albums, it is the album id.

```graphql
    album(visitorId: "", albumId: "") {
        title
        duration
        artists {
          id
          name
        }
        thumbnails {
          width
          height
          url
        }
        tracks {
            title
            id
            duration
            trackNumber
            explicit
            artists {
              id
              name
            }
            thumbnails {
              width
              height
              url
            }
        }
    }
```

### Contents of the single

> **If the visitorId is not sent, one will be generated automatically**.
> Remember to store it for future queries.  
> `singleId` Found in all singles, it is the id of the single.

```graphql
    single(visitorId: "", singleId: "") {
        title
        duration
        artists {
          id
          name
        }
        thumbnails {
          width
          height
          url
        }
        tracks {
            title
            id
            duration
            trackNumber
            explicit
            artists {
              id
              name
            }
            thumbnails {
              width
              height
              url
            }
        }
    }
```

## Types

### Common types

- #### _thumbnails_

  ```go
  [{
    width  int
    height int
    url    string
  }]
  ```

- ### _song_

  [artists](#artist) [thumbnails](#thumbnails)

  ```go
  {
    title        string
    id           string
    artists      []artist
    albumTitle?  string
    albumID?     string
    duration?    string
    explicit?    bool
    thumbnails   thumbnails
    trackNumber?  int
  }
  ```

- ### _single_

  [artists](#artist) [thumbnails](#thumbnails)

  ```go
  {
    title      string
    artists    []artist
    id         string
    year       int
    thumbnails thumbnails
    explicit   bool
  }
  ```

- ### _bestMatch_

  [artists](#artist) [thumbnails](#thumbnails)

  ```c
  {
    type         string
    albumType?   string
    title        string
    id           string
    thumbnails   thumbnails
    artists?     []artist
    albumTitle?  string
    albumID?     string
    duration?    string
    explicit?    bool
  }
  ```

- ### _artist_

  [thumbnails](#thumbnails)

  ```go
  {
    name       string
    id         string
    thumbnails thumbnails
  }
  ```

- ### _album_

  [artists](#artist) [thumbnails](#thumbnails)

  ```go
  {
    title       string
    id          string
    artists?    []artist
    thumbnails? thumbnails
    single?     bool
    ep?         bool
    year?       int
    explicit?   bool
  }
  ```

- ### _continuation_

  ```go
  {
    query   string
    params? string
  }
  ```

- ### _cardItem_

  [thumbnails](#thumbnails)

  ```go
  {
    thumbnails  thumbnails
    title       string
    year        int
    browseId    string
    params      string
    explicit    bool
  }
  ```

### Lists types

- ### _albumsList_

  [albums](#album) [continuation](#continuation)

  ```go
  {
      albums []album
      continuation continuation
      continuationId string
      visitorId      string
  }
  ```

- ### _musicsList_

  [songs](#song) [continuation](#continuation)

  ```go
  {
      songs []song
      continuation continuation
      continuationId string
      visitorId      string
  }
  ```

- ### _artistsList_
  [artists](#artist) [continuation](#continuation)
  ```go
  {
      artists []artist
      continuation continuation
      continuationId string
      visitorId      string
  }
  ```

### Artist types

- ### _artistAlbumsList_

  [cardItem](#carditem)

  ```go
  {
      albums             []cardItem
      continuationParams string
      continuationId     string
  }
  ```

- ### _artistSinglesList_

  [cardItem](#carditem)

  ```go
  {
      singles            []cardItem
      continuationParams string
      continuationId     string
  }
  ```

- ### _artistMusicsList_

  [songs](#song)

  ```go
  {
      songs              []song
      continuationParams string
      continuationId     string
  }
  ```

### Shelves types

- ### _searchShelf_

  [bestMatch](#bestmatch) [musicsList](#musicslist) [albumsList](#albumslist) [artistsList](#artistslist)

  ```go
  {
      bestMatch bestMatch
      tracks    musicsList
      albums    albumsList
      artists   artistsList
      visitorId string
  }
  ```

- ### _artistShelf_

  [thumbnails](#thumbnails) [artistMusicsList](#artistmusicslist) [artistAlbumsList](#artistalbumslist) [artistSinglesList](#artistsingleslist)

  ```go
    {
        name              string
        thumbnails        thumbnails
        musicsList        artistMusicsList
        albumsList        artistAlbumsList
        singlesList       artistSinglesList
        visitorId         string
    }
  ```

- ### _albumShelf_

  [artists](#artist) [thumbnails](#thumbnails) [songs](#song)

  ```go
    {
        title       string
        artists     []artist
        thumbnails  thumbnails
        tracks      []song
        duration    string
        year        int
        visitorId   string
    }
  ```

- ### _singleShelf_

  [artists](#artist) [thumbnails](#thumbnails) [songs](#song)

  ```go
    {
        title       string
        artists     []artist
        thumbnails  thumbnails
        tracks      []song
        duration    string
        year        int
        visitorId   string
    }
  ```
