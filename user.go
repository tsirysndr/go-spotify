package spotify

type UserService service

type SpfError struct {
	Error `json:"error,omitempty"`
}

type Error struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type Profile struct {
	DisplayName  string       `json:"display_name,omitempty"`
	ExternalURLs ExternalURLs `json:"external_urls,omitempty"`
	Followers    Followers    `json:"followers,omitempty"`
	Href         string       `json:"href,omitempty"`
	Id           string       `json:"id,omitempty"`
	Images       []Img        `json:"images,omitempty"`
	Type         string       `json:"type,omitempty"`
	Uri          string       `json:"uri,omitempty"`
}

type ExternalURLs struct {
	Spotify string `json:"spotify,omitempty"`
}

type Followers struct {
	Href  string `json:"href,omitempty"`
	Total int    `json:"total,omitempty"`
}

type Img struct {
	Height string `json:"height,omitempty"`
	Url    string `json:"url,omitempty"`
	Width  string `json:"width,omitempty"`
}

func (s *UserService) GetProfile(ID string) (*Profile, error) {
	var err error
	resp := new(Profile)
	s.client.base.Path("users/").Get(ID).Receive(resp, err)
	return resp, err
}
