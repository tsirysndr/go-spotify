package spotify

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
