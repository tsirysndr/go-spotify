package spotify

import (
	"github.com/dghubble/sling"
)

type Client struct {
	base     *sling.Sling
	common   service
	Album    *AlbumService
	Artist   *ArtistService
	Category *CategoryService
	Library  *LibraryService
	Me       *MeService
	Playlist *PlaylistService
	Search   *SearchService
	Track    *TrackService
	User     *UserService
}

type service struct {
	client *Client
}

func NewClient() *Client {
	base := sling.New().Base("https://api.spotify.com/v1/")
	c := &Client{}
	c.base = base
	c.common.client = c
	c.Album = (*AlbumService)(&c.common)
	c.Artist = (*ArtistService)(&c.common)
	c.Category = (*CategoryService)(&c.common)
	c.Library = (*LibraryService)(&c.common)
	c.Me = (*MeService)(&c.common)
	c.Playlist = (*PlaylistService)(&c.common)
	c.Search = (*SearchService)(&c.common)
	c.Track = (*TrackService)(&c.common)
	c.User = (*UserService)(&c.common)
	return c
}
