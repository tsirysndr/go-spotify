package spotify

type PlayerService service

type PlayingTrack struct {
	Timestamp            int      `json:"timestamp,omitempty"`
	Context              *Context `json:"context,omitempty"`
	ProgressMS           int      `json:"progress_ms,omitempty"`
	Item                 *Track   `json:"item,omitempty"`
	CurrentlyPlayingType string   `json:"currently_playing_type,omitempty"`
	Actions              *Actions `json:"actions,omitempty"`
	IsPlaying            bool     `json:"is_playing,omitempty"`
}

type RecentlyPlayedTracks struct {
	Items []struct {
		Track    *Track   `json:"track,omitempty"`
		PlayedAt string   `json:"played_at,omitempty"`
		Context  *Context `json:"context,omitempty"`
	} `json:"items,omitempty"`
	Next    string `json:"next,omitempty"`
	Cursors *struct {
		After  string `json:"after,omitempty"`
		Before string `json:"before,omitempty"`
	} `json:"cursors,omitempty"`
	Limit int    `json:"limit,omitempty"`
	Href  string `json:"href,omitmepty"`
}

type Context struct {
	ExternalURLs ExternalURLs `json:"external_urls,omitempty"`
	Href         string       `json:"href,omitempty"`
	Type         string       `json:"type,omitempty"`
	URI          URI          `json:"uri,omitempty"`
}

type Actions struct {
	Disallows *struct {
		Resuming bool `json:"resuming,omitempty"`
	} `json:"disallows,omitempty"`
}

type RecentlyPlayedTracksParams struct {
	Type  string `url:"type"`
	Limit int    `url:"type"`
	After int    `url:"after"`
}

func (s *PlayerService) GetCurrentlyPlaying() (*PlayingTrack, error) {
	var err error
	res := new(PlayingTrack)
	s.client.base.Path("me/player/").Get("currently-playing").Receive(res, err)
	return res, err
}

func (s *PlayerService) GetRecentlyPlayed(filter string, limit, after int) (*RecentlyPlayedTracks, error) {
	var err error
	params := &RecentlyPlayedTracksParams{
		filter,
		limit,
		after,
	}
	res := new(RecentlyPlayedTracks)
	s.client.base.Path("me/player/").Get("recently-played").QueryStruct(params).Receive(res, err)
	return res, err
}
