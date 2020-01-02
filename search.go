package spotify

type SearchService service

type SearchResult struct {
	Artists *struct {
		Href     string   `json:"href,omitempty"`
		Items    []Artist `json:"items,omitempty"`
		Limit    int      `json:"limit,omitempty"`
		Next     string   `json:"string,omitempty"`
		Offset   int      `json:"offset,omitempty"`
		Previous string   `json:"previous,omitempty"`
		Total    int      `json:"total,omitempty"`
	} `json:"artists,omitempty"`
	Tracks *struct {
		Href     string  `json:"href,omitempty"`
		Items    []Track `json:"items,omitempty"`
		Limit    int     `json:"limit,omitempty"`
		Next     string  `json:"string,omitempty"`
		Offset   int     `json:"offset,omitempty"`
		Previous string  `json:"previous,omitempty"`
		Total    int     `json:"total,omitempty"`
	} `json:"tracks,omitempty"`
	Albums *struct {
		Href     string  `json:"href,omitempty"`
		Items    []Album `json:"items,omitempty"`
		Limit    int     `json:"limit,omitempty"`
		Next     string  `json:"string,omitempty"`
		Offset   int     `json:"offset,omitempty"`
		Previous string  `json:"previous,omitempty"`
		Total    int     `json:"total,omitempty"`
	} `json:"albums,omitempty"`
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
