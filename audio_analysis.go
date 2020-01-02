package spotify

type Key int
type Mode int

type AudioAnalysis struct {
	Bars     []Marker      `json:"bars,omitempty"`
	Beats    []Marker      `json:"beats,omitempty"`
	Meta     AnalysisMeta  `json:"meta,omitempty"`
	Sections []Section     `json:"sections,omitempty"`
	Segments []Segment     `json:"segments,omitempty"`
	Tatums   []Marker      `json:"tatums,omitempty"`
	Track    AnalysisTrack `json:"track,omitempty"`
}

// Marker represents beats, bars, tatums and are used in segment and section
// descriptions.
type Marker struct {
	Start      float64 `json:"start,omitempty"`
	Duration   float64 `json:"duration,omitempty"`
	Confidence float64 `json:"confidence,omitempty"`
}

// AnalysisMeta describes details about Spotify's audio analysis of the track
type AnalysisMeta struct {
	AnalyzerVersion string  `json:"analyzer_version,omitempty"`
	Platform        string  `json:"platform,omitempty"`
	DetailedStatus  string  `json:"detailed_status,omitempty"`
	StatusCode      int     `json:"status,omitempty"`
	Timestamp       int64   `json:"timestamp,omitempty"`
	AnalysisTime    float64 `json:"analysis_time,omitempty"`
	InputProcess    string  `json:"input_process,omitempty"`
}

// Section represents a large variation in rhythm or timbre, e.g. chorus, verse,
// bridge, guitar solo, etc. Each section contains its own descriptions of
// tempo, key, mode, time_signature, and loudness.
type Section struct {
	Marker
	Loudness                float64 `json:"loudness,omitempty"`
	Tempo                   float64 `json:"tempo,omitempty"`
	TempoConfidence         float64 `json:"tempo_confidence,omitempty"`
	Key                     Key     `json:"key,omitempty"`
	KeyConfidence           float64 `json:"key_confidence,omitempty"`
	Mode                    Mode    `json:"mode,omitempty"`
	ModeConfidence          float64 `json:"mode_confidence,omitempty"`
	TimeSignature           int     `json:"time_signature,omitempty"`
	TimeSignatureConfidence float64 `json:"time_signature_confidence,omitempty"`
}

// Segment is characterized by it's perceptual onset and duration in seconds,
// loudness (dB), pitch and timbral content.
type Segment struct {
	Marker
	LoudnessStart   float64   `json:"loudness_start,omitempty"`
	LoudnessMaxTime float64   `json:"loudness_max_time,omitempty"`
	LoudnessMax     float64   `json:"loudness_max,omitempty"`
	LoudnessEnd     float64   `json:"loudness_end,omitempty"`
	Pitches         []float64 `json:"pitches,omitempty"`
	Timbre          []float64 `json:"timbre,omitempty"`
}

// AnalysisTrack contains audio analysis data about the track as a whole
type AnalysisTrack struct {
	NumSamples              int64   `json:"num_samples,omitempty"`
	Duration                float64 `json:"duration,omitempty"`
	SampleMD5               string  `json:"sample_md5,omitempty"`
	OffsetSeconds           int     `json:"offset_seconds,omitempty"`
	WindowSeconds           int     `json:"window_seconds,omitempty"`
	AnalysisSampleRate      int64   `json:"analysis_sample_rate,omitempty"`
	AnalysisChannels        int     `json:"analysis_channels,omitempty"`
	EndOfFadeIn             float64 `json:"end_of_fade_in,omitempty"`
	StartOfFadeOut          float64 `json:"start_of_fade_out,omitempty"`
	Loudness                float64 `json:"loudness,omitempty"`
	Tempo                   float64 `json:"tempo,omitempty"`
	TempoConfidence         float64 `json:"tempo_confidence,omitempty"`
	TimeSignature           int     `json:"time_signature,omitempty"`
	TimeSignatureConfidence float64 `json:"time_signature_confidence,omitempty"`
	Key                     Key     `json:"key,omitempty"`
	KeyConfidence           float64 `json:"key_confidence,omitempty"`
	Mode                    Mode    `json:"mode,omitempty"`
	ModeConfidence          float64 `json:"mode_confidence,omitempty"`
	CodeString              string  `json:"codestring,omitempty"`
	CodeVersion             float64 `json:"code_version,omitempty"`
	EchoprintString         string  `json:"echoprintstring,omitempty"`
	EchoprintVersion        float64 `json:"echoprint_version,omitempty"`
	SynchString             string  `json:"synchstring,omitempty"`
	SynchVersion            float64 `json:"synch_version,omitempty"`
	RhythmString            string  `json:"rhythmstring,omitempty"`
	RhythmVersion           float64 `json:"rhythm_version,omitempty"`
}
