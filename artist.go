package spotify

type ArtistService service

type Artist struct {
	Name         string       `json:"name,omitempty"`
	ID           ID           `json:"id,omitempty"`
	URI          URI          `json:"uri,omitempty"`
	Href         string       `json:"href,omitempty"`
	ExternalURLs ExternalURLs `json:"external_urls,omitempty"`
	Type         string       `json:"type,omitempty"`
}
