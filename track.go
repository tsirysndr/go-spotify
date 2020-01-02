package spotify

type TrackService service

type TrackParams struct {
	IDs string `url:"ids,omitempty"`
}

type Track struct {
	Artists     []Artist     `json:"artists,omitempty"`
	Album       *Album       `json:"album,omitempty"`
	ExternalIDs *ExternalIDs `json:"external_ids,omitempty"`
	Popularity  int          `json:"popularity,omitempty"`
	DiscNumber  int          `json:"disc_number,omitempty"`
	DurationMS  int          `json:"duration_ms,omitempty"`
	ID          string       `json:"id,omitempty"`
	Href        string       `json:"href,omitempty"`
	IsLocal     bool         `json:"is_local,omitempty"`
	IsPlayable  bool         `json:"is_playable,omitempty"`
	Name        string       `json:"name,omitempty"`
	PreviewURL  string       `json:"preview_url,omitempty"`
	TrackNumber int          `json:"track_number,omitempty"`
	Type        string       `json:"type,omitempty"`
	URI         string       `json:"uri,omitempty"`
}

type Tracks struct {
	Href     string  `json:"href,omitempty"`
	Items    []Track `json:"items,omitempty"`
	Limit    int     `json:"limit,omitempty"`
	Next     string  `json:"next,omitempty"`
	Offset   int     `json:"offset,omitempty"`
	Previous string  `json:"previous,omitempty"`
	Total    int     `json:"total,omitempty"`
}

type TracksResult struct {
	Tracks []Track `json:"tracks,omitempty"`
}

func (s *TrackService) List(IDs string) (*TracksResult, error) {
	params := &TrackParams{IDs: IDs}
	var err error
	res := &TracksResult{}
	s.client.base.Get("tracks").QueryStruct(params).Receive(res, err)
	return res, err
}

func (s *TrackService) Get(ID string) (*Track, error) {
	var err error
	res := new(Track)
	s.client.base.Path("tracks/").Get(ID).Receive(res, err)
	return res, err
}

func (s *TrackService) AudioFeatures(IDs string) (*AudioFeaturesResult, error) {
	var err error
	params := &TrackParams{IDs}
	res := new(AudioFeaturesResult)
	s.client.base.Get("audio-features/").QueryStruct(params).Receive(res, err)
	return res, err
}

func (s *TrackService) AudioFeaturesByTrackID(ID string) (*AudioFeatures, error) {
	var err error
	res := new(AudioFeatures)
	s.client.base.Path("audio-features/").Get(ID).Receive(res, err)
	return res, err
}

func (s *TrackService) AudioAnalysis(ID string) (*AudioAnalysis, error) {
	var err error
	res := new(AudioAnalysis)
	s.client.base.Path("audio-analysis/").Get(ID).Receive(res, err)
	return res, err
}
