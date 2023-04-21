# Katze-YT

![YT-KATZE logo](./assets/logo.png)

Una API de [GraphQL](https://graphql.org) hecha en [Go](https://go.dev) para acceder de manera limpia a la API de [YouTube Music](https://music.youtube.com)

**Tecnologías usadas**

- [![Golang logo](./assets/go.png)Golang](https://go.dev)
- [![GraphQL logo](./assets/graphql.png)GraphQL](https://graphql.org)

**_No requiere autenticación_**

## Uso [YouTube Music](https://music.youtube.com)

- [Resultado de busqueda general](#busqueda-general)
- [Busqueda de canciones](#busqueda-de-canciones)
- [Paginación de canciones](#paginación-de-canciones)
- [Busqueda de albumes](#busqueda-de-albumes)
- [Pagination de albumes](#paginación-de-albumes)
- [Busqueda de artistas](#busqueda-de-artistas)
- [Paginación de artistas](#paginación-de-artistas)
- [Perfil de artista](#perfil-de-artista)
- [Canciones del artita](#canciones-del-artista)
- [Paginación de canciones del artista](#paginación-de-canciones-del-artista)
- [Albumes del artista](#albumes-del-artista)
- [Sencillos del artista](#sencillos-del-artista)
- [Contenido del album](#contenido-del-album)
- [Contenido del sencillo](#contenido-del-sencillo)

## Conjuntos de tipos:

- [common](#common-types)
  Son tipos compartidos, estos los utilizan otros tipos más complejos como los anteriormente mencionados.
  Contienen tipos primitivos como `album` o `song`

  - [thumbnails](#thumbnails)
  - [song](#song)
  - [single](#single)
  - [bestMatch](#bestmatch)
  - [artist](#artist)
  - [album](#album)
  - [continuation](#continuation)
  - [cardItem](#carditem)

- [Artist](#artist-types)
  Son los tipos relacionados a un artista, como su lista de canciones, de albumes y sencillos, estos contienen: un array de datos
  de algún otro tipo, params y continuationId

  - [artistAlbumsList](#artistalbumslist)
  - [artistSinglesList](#artistsingleslist)
  - [artistMusicsList](#artistmusicslist)

- [Lists](#lists-types)
  Son listas de otros tipos, pero estás incluyen continuación, por lo tanto tienen un array de datos de algún otro tipo
  acompañadas de un campo [continuation](#continuation), continuationId y un visitorId

  - [musicsList](#musicslist)
  - [albumsList](#albumslist)
  - [artistsList](#artistslist)

- [Shelves](#shelves-types)
  Son estanterías, estás incluyen tipos complejos, por ejemplo: Un album con sus canciones, imagenes y descripción,
  un artista con sus canciones, albumes, sencillos, y descripción, un sencillo con su canción, imagenes y descipción,
  la busqueda general con sus canciones, albumes y artistas.

  - [searchShelf](#searchshelf)
  - [artistShelf](#artistshelf)
  - [albumShelf](#albumshelf)

## Uso

### Busqueda general

En la query establesca la consulta y en visitorId el id del visitante.

> **Si no se establece el id del visitante, se generará uno automáticamente.**  
> !Recuerde almacenarlo para futuras consultas.  
> no podrá obtener los albumes de un artista, ni los sencillos de un artista sin un id de visitante valido
> que se haya usado previamente para solicitar el perfil del artista.

> El bestMatch es el mejor resultado de la busqueda, puede ser `song`, `album` o `artist`.
>
> - si es una `canción` **`No contendrá`** `albumType`
> - Si es `album` o `single` **`No contendrá`** `album` ni `albumId` el titulo del album será `title` y se especificará el tipo en `albumType`
> - si es un video devolvera song en type, pero **`No contendrá`** `albumType`, `album` ni `albumId`
> - si es un artista **`No contendrá`** `albumType`, `album` ni `albumId`

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

### Busqueda de canciones

> **Si no se establece el id del visitante, se generará uno automáticamente.**
> !Recuerde almacenarlo para futuras consultas.
> `params` es muy importante, este es quién realmente indica lo que buscamos.
> se obtiene en la query `general` en el campo `tracks.continuation.params` > `query` puede cambiarse u obtenerse del campo `tracks.continuation.query`

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

### Paginación de canciones

> **Si no se establece el id del visitante, se generará uno automáticamente.**  
> !Recuerde almacenarlo para futuras consultas.  
> `continuationId` es muy importante, este es quién realmente indica lo que buscamos.  
> se obtiene en la query `musicsList` en el campo `continuationId`

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

### Busqueda de albumes

> **Si no se establece el id del visitante, se generará uno automáticamente.**  
> !Recuerde almacenarlo para futuras consultas.  
> `params` es muy importante, este es quién realmente indica lo que buscamos.
> se obtiene en la query `general` en el campo `albums.continuation.params` > `query` puede cambiarse u obtenerse del campo `albums.continuation.query`

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

### Paginación de albumes

> **Si no se establece el id del visitante, se generará uno automáticamente.**
> !Recuerde almacenarlo para futuras consultas.  
> `continuationId` es muy importante, este es quién realmente indica lo que buscamos.
> se obtiene en la query `albumsList` en el campo `continuationId`

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

### Busqueda de artistas

> **Si no se establece el id del visitante, se generará uno automáticamente.**
> !Recuerde almacenarlo para futuras consultas.  
> `params` es muy importante, este es quién realmente indica lo que buscamos.
> `params` se obtiene en la query `general` en el campo `artists.continuation.params`
> y la query puede cambiarse u obtenerla de allí mismo.

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

### Paginación de artistas

> **Si no se establece el id del visitante, se generará uno automáticamente.**  
> !Recuerde almacenarlo para futuras consultas.  
> `continuationId` es muy importante, este es quién realmente indica lo que buscamos.
> `continuationId` se obtiene en la query `artistsList` en el campo `continuationId`

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

### Perfil de artista

> **Si no se establece el id del visitante, se generará uno automáticamente.**  
> !Recuerde almacenarlo para futuras consultas.
> `artistId` se encuentra en todos los resultados que contengan una tarjeta de artista.
> por ejemplo en la query `musicsList` en el campo `songs.artists.id`

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

### Canciones del artista

> El visitorId es opcional, permite personalizar los resultados de la búsqueda.  
> En este caso **No se genera visitorId**.
> `params` y continuation id se encuentran en el campo musicsList del [Perfil de artista](#perfil-de-artista)

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

### Paginación de canciones del artista

> El visitorId es opcional, permite personalizar los resultados de la búsqueda.  
> En este caso **No se genera visitorId**.
> `continuationId` se encuentra en [Canciones del artista](#canciones-del-artista)

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

### Albumes del artista

> El visitorId es obligatorio, permite personalizar los resultados de la búsqueda.  
> **NO HABRÁN RESULTADOS SI NO SE ENVÍA EL visitorId** > `artistId` y `continuationId`
> se encuentran en el campo albumsList de [artistProfile](#perfil-de-artista).

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

### Sencillos del artista

> El visitorId es obligatorio, permite personalizar los resultados de la búsqueda.  
> **NO HABRÁN RESULTADOS SI NO SE ENVÍA EL visitorId** > `artistId` y `continuationId`  
> se encuentran en el campo siglesList de [artistProfile](#perfil-de-artista).

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

### Contenido del album

> **Si no se envía el visitorId, se generará uno automáticamente.**
> !Recuerde almacenarlo para futuras consultas.  
> `albumId` Se encuentra en todos los albumes, es el id del album.

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

### Contenido del sencillo

> **Si no se envía el visitorId, se generará uno automáticamente.**
> !Recuerde almacenarlo para futuras consultas.  
> `singleId` Se encuentra en todos los sencillos, es el id del sencillo.

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

## Tipos

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
