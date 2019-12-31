package spotify

type TrackService service

type TrackParams struct {
	Ids string `url:"ids,omitempty"`
}

type Track struct {
	Artists []Artist `json:"artists,omitempty"`
	Album   *struct {
		AlbumType string `json:"album_type,omitempty"`
	} `json:"album,omitempty"`
	ExternalIDs map[string]string `json:"external_ids,omitempty"`
	Popularity  int               `json:"popularity,omitempty"`
}

type TracksResult struct {
	Tracks []Track `json:"tracks,omitempty"`
}

func (s *TrackService) List(ids string) (*TracksResult, error) {
	params := &TrackParams{Ids: ids}

	var err error
	res := &TracksResult{}
	s.client.base.Get("tracks").QueryStruct(params).Receive(res, err)
	return res, err
}
