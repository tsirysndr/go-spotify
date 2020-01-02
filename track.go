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
	Href  string  `json:"href,omitempty"`
	Items []Track `json:"items,omitempty"`
}

type TracksResult struct {
	Tracks []Track `json:"tracks,omitempty"`
}

type AudioFeatures struct {
	Danceability     float32 `json:"danceablity,omitempty"`
	Energy           float32 `json:"energy,omitempty"`
	Key              int     `json:"key,omitempty"`
	Loudness         float32 `json:"loudness,omitempty"`
	Mode             int     `json:"mode,omitempty"`
	Speechiness      float32 `json:"speechiness,omitempty"`
	Acousticness     float32 `json:"acousticness,omitempty"`
	Instrumentalness float32 `json:"instrumentalness,omitempty"`
	Liveness         float32 `json:"liveness,omitempty"`
	Valence          float32 `json:"valence,omitempty"`
	Tempo            float32 `json:"tempo,omitempty"`
	Type             string  `json:"type,omitempty"`
	ID               ID      `json:"id,omitempty"`
	URI              URI     `json:"uri,omitempty"`
	TrackHref        string  `json:"track_href,omitempty"`
	AnalysisURL      string  `json:"analysis_url,omitempty"`
	DurationMS       int     `json:"duration_ms,omitempty"`
	TimeSignature    int     `json:"time_signature,omitempty"`
}

type AudioFeaturesResult struct {
	AudioFeatures []AudioFeatures `json:"audio_features,omitempty"`
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
