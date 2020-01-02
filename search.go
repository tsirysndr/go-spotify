package spotify

type SearchService service

type SearchResult struct {
	Artists *Artists `json:"artists,omitempty"`
	Tracks  *Tracks  `json:"tracks,omitempty"`
	Albums  *Artists `json:"albums,omitempty"`
}

type SearchParams struct {
	Q      string `url:"q"`
	Type   string `url:"type"`
	Limit  int    `url:"limit"`
	Offset int    `url:"offset"`
}

func (s *SearchService) Get(query, filter string, limit, offset int) (*SearchResult, error) {
	var err error
	params := &SearchParams{query, filter, limit, offset}
	res := new(SearchResult)
	s.client.base.Get("search").QueryStruct(params).Receive(res, err)
	return res, err
}
