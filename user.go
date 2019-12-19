package spotify

type UserService service

type SpfError struct {
	Error `json:"error"`
}

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Profile struct {
	DisplayName  string      `json:"display_name"`
	SpotifyUri   spotifyUri  `json:"external_urls"`
	Followers    Followers  `json:"followers"`
	Href         string      `json:"href"`
	Id           string      `json:"id"`
	Images       []Img       `json:"images"`
	Type         string      `json:"type"`
	Uri          string      `json:"uri"`
}

type spotifyUri struct {
	Spotify string `json:"spotify"`
}

type Followers  struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type Img struct {
	Height string `json:"height"`
	Url    string `json:"url"`
	Width  string `json:"width"`
}

func (c *UserService) UserProfile(Id string) (interface{}, error) {
	var err error
	res := new(Profile)
	spfErr := new(SpfError)
	resp, err := c.client.base.Path("users/").Get(ID).Receive(res, spfErr)

	if resp.StatusCode != 200 {
		return spfErr, err
	}

	return res, err
}
