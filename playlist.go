package spotify

import "fmt"

type PlaylistService service

type Playlist struct {
	Collaborative bool          `json:"collaborative,omitempty"`
	Description   string        `json:"description,omitempty"`
	ExternalURLs  *ExternalURLs `json:"external_urls,omitempty"`
	Followers     *Followers    `json:"followers,omitempty"`
	Href          string        `json:"href,omitempty"`
	ID            ID            `json:"id,omitempty"`
	Images        []Image       `json:"images,omitempty"`
	Name          string        `json:"name,omitempty"`
	Owner         *Profile      `json:"profile,omitempty"`
	PrimaryColor  string        `json:"primary_color,omitempty"`
	Public        bool          `json:"public,omitempty"`
	SnapshotID    string        `json:"snapshot_id,omitempty"`
	Tracks        *struct {
		Href  string `json:"href,omitempty"`
		Items []struct {
			AddedAt        string   `json:"added_at,omitempty"`
			AddedBy        *Profile `json:"added_by,omitempty"`
			IsLocal        bool     `json:"is_local,omitempty"`
			PrimaryColor   string   `json:"primary_color,omitempty"`
			Track          *Track   `json:"track,omitempty"`
			VideoThumbnail *struct {
				URL string `json:"url.omitempty"`
			} `json:"video_thumbnail,omitempty"`
		} `json:"items,omitempty"`
	} `json:"tracks,omitempty"`
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
