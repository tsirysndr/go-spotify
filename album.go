package spotify

type AlbumService service

// album object (full)

type Album struct {
	AlbumType            string      `json:"album_type"`
	Artists              []Artist    `json:"artists"`
	AvailableMarkets     []string    `json:"available_markets"`
	Copyrights           []Copyright `json:"copyrights"`
	ExternalIds          External    `json:"external_ids"`
	ExternalUrls         External    `json:"external_urls"`
	Genres               []string    `json:"genres"`
	Href                 string      `json:"href"`
	Id                   string      `json:"id"`
	Images               Image       `json:"images"`
	Label                string      `json:"label"`
	Name                 string      `json:"name"`
	Popularity           int         `json:"popularity"`
	ReleaseDate          string      `json:"release_date"`
	ReleaseDatePrecision string      `json:"release_date_precision"`
	Tracks               []Track     `json:"tracks"`
	Type                 string      `json:"type"`
	Uri                  string      `json:"uri"`
}

type Copyright struct {
	Text string `json:"album_type"`
	Type string `json:"type"`
}

type External struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Image struct {
	Height int    `json:"height"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
}

func (c *AlbumService) ListAlbum(ID string) (*Album, error) {
	var err error
	res := new(Album)
	c.client.base.Path("albums/").Get(ID).Receive(res, err)

	return res, err
}
