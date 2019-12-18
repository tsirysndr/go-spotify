package spotify

import (
	"fmt"
)

type UserService service

type SpfError struct {
	Error `json:"error"`
}

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResultProfile struct {
	DisplayName  string      `json:"display_name"`
	ExternalUrls spotifyUri  `json:"external_urls"`
	Followers    dataFlowers `json:"followers"`
	Href         string      `json:"href"`
	Id           string      `json:"id"`
	Images       []Img       `json:"images"`
	Type         string      `json:"type"`
	Uri          string      `json:"uri"`
}

type spotifyUri struct {
	Spotify string `json:"spotify"`
}

type dataFlowers struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type Img struct {
	Height string `json:"height"`
	Url    string `json:"url"`
	Width  string `json:"width"`
}

func (c *UserService) UserProfile(id string) (interface{}, error) {
	var err error
	res := new(ResultProfile)
	spfErr := new(SpfError)
	path := fmt.Sprintf("%s/%s", "users", id)
	resp, err := c.client.base.New().Get(path).Receive(res, spfErr)

	if resp.StatusCode != 200 {
		return spfErr, err
	}

	return res, err

}
