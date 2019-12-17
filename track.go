package spotify

type TrackService service

type TrackParams struct {
	Ids string `url:"ids,omitempty"`
}

type Track struct {
	Artists []Artist `json:"artists"`
	Album   *struct {
		AlbumType string `json:"album_type"`
	} `json:"album,omitempty"`
	ExternalIDs map[string]string `json:"external_ids"`
	Popularity  int               `json:"popularity"`
}

type TracksResult struct {
	Tracks []Track `json:"tracks"`
}

func (s *TrackService) ListTracks(ids string) (*TracksResult, error) {
	params := &TrackParams{Ids: ids}

	var err error
	res := &TracksResult{}
	s.client.base.Get("tracks").QueryStruct(params).Receive(res, err)
	return res, err
}
