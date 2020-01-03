package spotify

import (
	"github.com/dghubble/sling"
	"golang.org/x/oauth2"
)

// URI identifies an artist, album, track, or category.  For example,
// spotify:track:6rqhFgbbKwnb9MLmUQDhG6
type URI string

// ID is a base-62 identifier for an artist, track, album, etc.
// It can be found at the end of a spotify.URI.
type ID string

type Client struct {
	base     *sling.Sling
	common   service
	Album    *AlbumService
	Artist   *ArtistService
	Category *CategoryService
	Library  *LibraryService
	Me       *MeService
	Playlist *PlaylistService
	Player   *PlayerService
	Search   *SearchService
	Track    *TrackService
	User     *UserService
}

type service struct {
	client *Client
}

func NewClient(accessToken string) *Client {
	config := &oauth2.Config{}
	token := &oauth2.Token{AccessToken: accessToken}
	httpClient := config.Client(oauth2.NoContext, token)
	base := sling.New().Base("https://api.spotify.com/v1/").Client(httpClient)
	c := &Client{}
	c.base = base
	c.common.client = c
	c.Album = (*AlbumService)(&c.common)
	c.Artist = (*ArtistService)(&c.common)
	c.Category = (*CategoryService)(&c.common)
	c.Library = (*LibraryService)(&c.common)
	c.Me = (*MeService)(&c.common)
	c.Playlist = (*PlaylistService)(&c.common)
	c.Player = (*PlayerService)(&c.common)
	c.Search = (*SearchService)(&c.common)
	c.Track = (*TrackService)(&c.common)
	c.User = (*UserService)(&c.common)
	return c
}
