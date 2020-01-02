package spotify

type ArtistService service

type Artist struct {
	Name         string       `json:"name,omitempty"`
	ID           ID           `json:"id,omitempty"`
	URI          URI          `json:"uri,omitempty"`
	Href         string       `json:"href,omitempty"`
	ExternalURLs ExternalURLs `json:"external_urls,omitempty"`
	Type         string       `json:"type,omitempty"`
	Followers    *Followers   `json:"followers,omitempty"`
	Genres       []string     `json:"genres,omitempty"`
	Images       []Image      `json:"images,omitempty"`
	Popularity   int          `json:"popularity,omitempty"`
}

type ArtistParams struct {
	IDs string `url:"ids,omitempty"`
}

type ArtistsResult struct {
	Artists []Artist `json:"artists,omitempty"`
}

func (s *ArtistService) List(IDs string) (*ArtistsResult, error) {
	params := &ArtistParams{IDs: IDs}
	var err error
	res := &ArtistsResult{}
	s.client.base.Get("artists").QueryStruct(params).Receive(res, err)
	return res, err
}
