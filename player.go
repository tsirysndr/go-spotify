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

func (s *PlayerService) GetCurrentlyPlaying() (*PlayingTrack, error) {
	var err error
	res := new(PlayingTrack)
	s.client.base.Path("me/player/").Get("currently-playing").Receive(res, err)
	return res, err
}
