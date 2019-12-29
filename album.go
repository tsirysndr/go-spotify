package spotify

type AlbumService service

// album object (full)

type Album struct {
	AlbumType            string      `json:"album_type,omitempty"`
	Artists              []Artist    `json:"artists,omitempty"`
	AvailableMarkets     []string    `json:"available_markets,omitempty"`
	Copyrights           []Copyright `json:"copyrights,omitempty"`
	ExternalIds          External    `json:"external_ids,omitempty"`
	ExternalUrls         External    `json:"external_urls,omitempty"`
	Genres               []string    `json:"genres,omitempty"`
	Href                 string      `json:"href,omitempty"`
	Id                   string      `json:"id,omitempty"`
	Images               Image       `json:"images,omitempty"`
	Label                string      `json:"label,omitempty"`
	Name                 string      `json:"name,omitempty"`
	Popularity           int         `json:"popularity,omitempty"`
	ReleaseDate          string      `json:"release_date,omitempty"`
	ReleaseDatePrecision string      `json:"release_date_precision,omitempty"`
	Tracks               []Track     `json:"tracks,omitempty"`
	Type                 string      `json:"type,omitempty"`
	Uri                  string      `json:"uri,omitempty"`
}

type Copyright struct {
	Text string `json:"album_type,omitempty"`
	Type string `json:"type,omitempty"`
}

type External struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type Image struct {
	Height int    `json:"height,omitempty"`
	Url    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
}

func (c *AlbumService) ListAlbum(ID string) (*Album, error) {
	var err error
	res := new(Album)
	c.client.base.Path("albums/").Get(ID).Receive(res, err)

	return res, err
}
