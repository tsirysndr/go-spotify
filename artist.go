package spotify

type ArtistService service

type Artist struct {
	Name string `json:"name,omitempty"`
	ID   ID     `json:"id,omitempty"`
	// The Spotify URI for the artist.
	URI URI `json:"uri,omitempty"`
	// A link to the Web API enpoint providing full details of the artist.
	Endpoint     string            `json:"href,omitempty"`
	ExternalURLs map[string]string `json:"external_urls,omitempty"`
}
