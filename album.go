package spotify

type AlbumService service

// album object (full)

type Album struct {
	AlbumType            string        `json:"album_type,omitempty"`
	Artists              []Artist      `json:"artists,omitempty"`
	AvailableMarkets     []string      `json:"available_markets,omitempty"`
	Copyrights           []Copyright   `json:"copyrights,omitempty"`
	ExternalIDs          *ExternalIDs  `json:"external_ids,omitempty"`
	ExternalURLs         *ExternalURLs `json:"external_urls,omitempty"`
	Genres               []string      `json:"genres,omitempty"`
	Href                 string        `json:"href,omitempty"`
	ID                   string        `json:"id,omitempty"`
	Images               []Image       `json:"images,omitempty"`
	Label                string        `json:"label,omitempty"`
	Name                 string        `json:"name,omitempty"`
	Popularity           int           `json:"popularity,omitempty"`
	ReleaseDate          string        `json:"release_date,omitempty"`
	ReleaseDatePrecision string        `json:"release_date_precision,omitempty"`
	TotalTracks          int           `json:"total_tracks,omitempty"`
	Tracks               *struct {
		Href  string  `json:"href,omitempty"`
		Items []Track `json:"items,omitempty"`
	} `json:"tracks,omitempty"`
	Type string `json:"type,omitempty"`
	URI  string `json:"uri,omitempty"`
}

type Copyright struct {
	Text string `json:"album_type,omitempty"`
	Type string `json:"type,omitempty"`
}

type ExternalIDs struct {
	UPC string `json:"upc,omitempty"`
}

type Image struct {
	Height int    `json:"height,omitempty"`
	Url    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
}

type AlbumParams struct {
	IDs string `url:"ids,omitempty"`
}

type AlbumsResult struct {
	Albums []Album `json:"albums,omitempty"`
}

func (s *AlbumService) Get(ID string) (*Album, error) {
	var err error
	res := new(Album)
	s.client.base.Path("albums/").Get(ID).Receive(res, err)
	return res, err
}

func (s *AlbumService) List(IDs string) (*AlbumsResult, error) {
	params := &AlbumParams{IDs: IDs}
	var err error
	res := &AlbumsResult{}
	s.client.base.Get("albums").QueryStruct(params).Receive(res, err)
	return res, err
}
