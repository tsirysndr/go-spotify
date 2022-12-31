package spotify

import (
	"fmt"
	"net/url"
)

type PlaylistService service

type Playlist struct {
	Collaborative bool            `json:"collaborative,omitempty"`
	Description   string          `json:"description,omitempty"`
	ExternalURLs  *ExternalURLs   `json:"external_urls,omitempty"`
	Followers     *Followers      `json:"followers,omitempty"`
	Href          string          `json:"href,omitempty"`
	ID            ID              `json:"id,omitempty"`
	Images        []Image         `json:"images,omitempty"`
	Name          string          `json:"name,omitempty"`
	Owner         *Profile        `json:"profile,omitempty"`
	PrimaryColor  string          `json:"primary_color,omitempty"`
	Public        bool            `json:"public,omitempty"`
	SnapshotID    string          `json:"snapshot_id,omitempty"`
	Tracks        *PlaylistTracks `json:"tracks,omitempty"`
}

type PlaylistTracks struct {
	Href  string `json:"href,omitempty"`
	Items []struct {
		AddedAt        string   `json:"added_at,omitempty"`
		AddedBy        *Profile `json:"added_by,omitempty"`
		IsLocal        bool     `json:"is_local,omitempty"`
		PrimaryColor   string   `json:"primary_color,omitempty"`
		Track          *Track   `json:"track,omitempty"`
		VideoThumbnail *struct {
			URL string `json:"url,omitempty"`
		} `json:"video_thumbnail,omitempty"`
	} `json:"items,omitempty"`
}

type Playlists struct {
	Href     string     `json:"href,omitempty"`
	Items    []Playlist `json:"items,omitempty"`
	Limit    int        `json:"limit,omitempty"`
	Offset   int        `json:"offset,omitempty"`
	Previous string     `json:"previous,omitempty"`
	Total    int        `json:"total,omitempty"`
}

type PlaylistAddTracks struct {
	SnapshotID string `json:"snapshot_id"`
}

type CreatePlaylist struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Public      bool   `json:"public"`
}

func (s *PlaylistService) Get(ID string) (*Playlist, error) {
	var err error
	res := new(Playlist)
	s.client.base.Path("playlists/").Get(ID).Receive(res, err)
	return res, err
}

func (s *PlaylistService) GetCoverImages(ID string) (*[]Image, error) {
	var err error
	res := new([]Image)
	playlists := fmt.Sprintf("playlists/%s/", ID)
	s.client.base.Path(playlists).Get("images").Receive(res, err)
	return res, err
}

func (s *PlaylistService) GetTracks(ID string, limit, offset int) (*PlaylistTracks, error) {
	var err error
	params := &PaginationParams{limit, offset}
	res := new(PlaylistTracks)
	playlists := fmt.Sprintf("playlists/%s/", ID)
	s.client.base.Path(playlists).Get("tracks").QueryStruct(params).Receive(res, err)
	return res, err
}

func (s *PlaylistService) GetUserPlaylists(userID string, limit, offset int) (*Playlists, error) {
	var err error
	params := &PaginationParams{limit, offset}
	res := new(Playlists)
	users := fmt.Sprintf("users/%s/", userID)
	s.client.base.Path(users).Get("playlists").QueryStruct(params).Receive(res, err)
	return res, err
}

func (s *PlaylistService) GetCurrentUserPlaylists(limit, offset int) (*Playlists, error) {
	var err error
	params := &PaginationParams{limit, offset}
	res := new(Playlists)
	s.client.base.Path("me/").Get("playlists").QueryStruct(params).Receive(res, err)
	return res, err
}

func (s *PlaylistService) AddTracksToPlaylists(playlistID string, tracksIDS string) (*PlaylistAddTracks, error) {
	var err error
	res := new(PlaylistAddTracks)
	tracks := fmt.Sprintf("playlists/%s/tracks?uris=%s", playlistID, url.QueryEscape(tracksIDS))
	s.client.base.Post(tracks).Receive(res, err)

	return res, err
}

func (s *PlaylistService) CreatePlaylist(userID string, name string, description string, public bool) (*CreatePlaylist, error) {
	var err error
	res := new(CreatePlaylist)
	tracks := fmt.Sprintf("users/%s/playlists", userID)

	newPlaylist := CreatePlaylist{
		Name:        name,
		Description: description,
		Public:      public,
	}

	s.client.base.Post(tracks).BodyJSON(newPlaylist).Receive(res, err)

	return res, err
}
