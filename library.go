package spotify

type LibraryService service

type SavedTracks struct {
	Href  string `json:"href,omitempty"`
	Items []struct {
		AddedAt string `json:"added_at,omitempty"`
		Track   *Track `json:"track,omitempty"`
	} `json:"items,omitempty"`
	Limit    int    `json:"limit,omitempty"`
	Next     string `json:"next,omitempty"`
	Offset   int    `json:"offset,omitempty"`
	Previous string `json:"previous,omitempty"`
	Total    int    `json:"total,omitempty"`
}

type SavedAlbums struct {
	Href  string `json:"href,omitempty"`
	Items []struct {
		AddedAt string `json:"added_at,omitempty"`
		Album   *Album `json:"album,omitempty"`
	} `json:"items,omitempty"`
	Limit    int    `json:"limit,omitempty"`
	Next     string `json:"next,omitempty"`
	Offset   int    `json:"offset,omitempty"`
	Previous string `json:"previous,omitempty"`
	Total    int    `json:"total,omitempty"`
}

type LibraryParams struct {
	IDs string `url:"ids"`
}

func (s *LibraryService) GetCurrentUserSavedTracks(limit, offset int) (*SavedTracks, error) {
	var err error
	params := &PaginationParams{limit, offset}
	res := new(SavedTracks)
	s.client.base.Path("me/").Get("tracks").QueryStruct(params).Receive(res, err)
	return res, err
}

func (s *LibraryService) GetCurrentUserSavedAlbums(limit, offset int) (*SavedAlbums, error) {
	var err error
	params := &PaginationParams{limit, offset}
	res := new(SavedAlbums)
	s.client.base.Path("me/").Get("albums").QueryStruct(params).Receive(res, err)
	return res, err
}

func (s *LibraryService) CheckCurrentUserSavedTracks(IDs string) (*[]bool, error) {
	var err error
	res := new([]bool)
	params := LibraryParams{IDs}
	s.client.base.Path("me/tracks/").Get("contains").QueryStruct(params).Receive(res, err)
	return res, err
}

func (s *LibraryService) CheckCurrentUserSavedAlbums(IDs string) (*[]bool, error) {
	var err error
	res := new([]bool)
	params := LibraryParams{IDs}
	s.client.base.Path("me/albums/").Get("contains").QueryStruct(params).Receive(res, err)
	return res, err
}
