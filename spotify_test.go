package spotify

import (
	"encoding/json"
	"testing"
)

func TestAlbumGet_JSON(t *testing.T) {
	data := `{
		"album_type": "album",
		"artists": [{"id": "0oOet2f43PA68X5RxKobEy", "name": "Radiohead", "uri": "spotify:artist:0oOet2f43PA68X5RxKobEy"}],
		"id": "6kCAaStByMSWbQwBGji5gp",
		"name": "OK Computer",
		"popularity": 78,
		"release_date": "1997-05-21",
		"total_tracks": 12,
		"type": "album",
		"uri": "spotify:album:6kCAaStByMSWbQwBGji5gp",
		"external_urls": {"spotify": "https://open.spotify.com/album/6kCAaStByMSWbQwBGji5gp"},
		"images": [{"height": 640, "url": "https://i.scdn.co/image/abc", "width": 640}],
		"label": "XL Recordings",
		"genres": ["alternative rock", "art rock"],
		"copyrights": [{"text": "2017 XL Recordings", "type": "C"}],
		"available_markets": ["US", "GB"]
	}`
	var album Album
	if err := json.Unmarshal([]byte(data), &album); err != nil {
		t.Fatalf("failed to unmarshal album: %v", err)
	}
	if album.ID != "6kCAaStByMSWbQwBGji5gp" {
		t.Errorf("expected ID 6kCAaStByMSWbQwBGji5gp, got %s", album.ID)
	}
	if album.Name != "OK Computer" {
		t.Errorf("expected OK Computer, got %s", album.Name)
	}
	if len(album.Artists) != 1 {
		t.Errorf("expected 1 artist, got %d", len(album.Artists))
	}
	if album.Artists[0].Name != "Radiohead" {
		t.Errorf("expected Radiohead, got %s", album.Artists[0].Name)
	}
	if album.Popularity != 78 {
		t.Errorf("expected popularity 78, got %d", album.Popularity)
	}
	if album.Tracks != nil {
		t.Errorf("expected nil tracks, got non-nil")
	}
}

func TestAlbumList_JSON(t *testing.T) {
	data := `{
		"albums": [{"id": "1", "name": "Album 1"}, {"id": "2", "name": "Album 2"}]
	}`
	var result AlbumsResult
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		t.Fatalf("failed to unmarshal albums result: %v", err)
	}
	if len(result.Albums) != 2 {
		t.Errorf("expected 2 albums, got %d", len(result.Albums))
	}
	if result.Albums[1].Name != "Album 2" {
		t.Errorf("expected Album 2, got %s", result.Albums[1].Name)
	}
}

func TestAlbumsPaginated_JSON(t *testing.T) {
	data := `{
		"href": "https://api.spotify.com/v1/albums?ids=1,2",
		"items": [{"id": "1", "name": "Album A"}, {"id": "2", "name": "Album B"}],
		"limit": 2,
		"offset": 0,
		"total": 2
	}`
	var albums Albums
	if err := json.Unmarshal([]byte(data), &albums); err != nil {
		t.Fatalf("failed to unmarshal albums list: %v", err)
	}
	if albums.Total != 2 {
		t.Errorf("expected total 2, got %d", albums.Total)
	}
	if len(albums.Items) != 2 {
		t.Errorf("expected 2 items, got %d", len(albums.Items))
	}
}

func TestArtist_JSON(t *testing.T) {
	data := `{
		"id": "0oOet2f43PA68X5RxKobEy",
		"name": "Radiohead",
		"popularity": 78,
		"genres": ["alternative rock", "art rock"],
		"type": "artist",
		"uri": "spotify:artist:0oOet2f43PA68X5RxKobEy",
		"followers": {"total": 1234567},
		"images": [{"height": 640, "url": "https://i.scdn.co/image/abc", "width": 640}],
		"external_urls": {"spotify": "https://open.spotify.com/artist/0oOet2f43PA68X5RxKobEy"}
	}`
	var artist Artist
	if err := json.Unmarshal([]byte(data), &artist); err != nil {
		t.Fatalf("failed to unmarshal artist: %v", err)
	}
	if artist.Name != "Radiohead" {
		t.Errorf("expected Radiohead, got %s", artist.Name)
	}
	if artist.Followers.Total != 1234567 {
		t.Errorf("expected 1234567 followers, got %d", artist.Followers.Total)
	}
}

func TestArtistsResult_JSON(t *testing.T) {
	data := `{
		"artists": [{"id": "1", "name": "Artist One"}, {"id": "2", "name": "Artist Two"}]
	}`
	var result ArtistsResult
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		t.Fatalf("failed to unmarshal artists result: %v", err)
	}
	if len(result.Artists) != 2 {
		t.Errorf("expected 2 artists, got %d", len(result.Artists))
	}
}

func TestTrack_JSON(t *testing.T) {
	data := `{
		"id": "6rqhFgbbKwnb9MLmUQDhG6",
		"name": "Paranoid Android",
		"duration_ms": 383000,
		"popularity": 65,
		"track_number": 2,
		"disc_number": 1,
		"type": "track",
		"uri": "spotify:track:6rqhFgbbKwnb9MLmUQDhG6",
		"is_local": false,
		"is_playable": true,
		"preview_url": "https://p.scdn.co/mp3-preview/abc",
		"artists": [{"id": "0oOet2f43PA68X5RxKobEy", "name": "Radiohead"}],
		"album": {"id": "6kCAaStByMSWbQwBGji5gp", "name": "OK Computer"}
	}`
	var track Track
	if err := json.Unmarshal([]byte(data), &track); err != nil {
		t.Fatalf("failed to unmarshal track: %v", err)
	}
	if track.Name != "Paranoid Android" {
		t.Errorf("expected Paranoid Android, got %s", track.Name)
	}
	if track.DurationMS != 383000 {
		t.Errorf("expected duration 383000, got %d", track.DurationMS)
	}
}

func TestTracks_JSON(t *testing.T) {
	data := `{
		"href": "https://api.spotify.com/v1/albums/abc/tracks",
		"items": [
			{"id": "1", "name": "Track 1"},
			{"id": "2", "name": "Track 2"}
		],
		"limit": 20,
		"offset": 0,
		"total": 2
	}`
	var tracks Tracks
	if err := json.Unmarshal([]byte(data), &tracks); err != nil {
		t.Fatalf("failed to unmarshal tracks: %v", err)
	}
	if tracks.Total != 2 {
		t.Errorf("expected 2 total tracks, got %d", tracks.Total)
	}
	if len(tracks.Items) != 2 {
		t.Errorf("expected 2 items, got %d", len(tracks.Items))
	}
}

func TestTracksResult_JSON(t *testing.T) {
	data := `{
		"tracks": [{"id": "1", "name": "Track 1"}]
	}`
	var result TracksResult
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		t.Fatalf("failed to unmarshal tracks result: %v", err)
	}
	if len(result.Tracks) != 1 {
		t.Errorf("expected 1 track, got %d", len(result.Tracks))
	}
}

func TestSearchResult_JSON(t *testing.T) {
	data := `{
		"artists": {
			"href": "https://api.spotify.com/v1/search?q=radiohead&type=artist",
			"items": [{"id": "0oOet2f43PA68X5RxKobEy", "name": "Radiohead"}],
			"limit": 20,
			"total": 1
		},
		"tracks": {
			"href": "https://api.spotify.com/v1/search?q=radiohead&type=track",
			"items": [{"id": "6rqhFgbbKwnb9MLmUQDhG6", "name": "Paranoid Android"}],
			"limit": 20,
			"total": 1
		}
	}`
	var result SearchResult
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		t.Fatalf("failed to unmarshal search result: %v", err)
	}
	if result.Artists == nil {
		t.Fatal("expected non-nil artists")
	}
	if len(result.Artists.Items) != 1 {
		t.Errorf("expected 1 artist item, got %d", len(result.Artists.Items))
	}
	if result.Artists.Items[0].Name != "Radiohead" {
		t.Errorf("expected Radiohead, got %s", result.Artists.Items[0].Name)
	}
	if len(result.Tracks.Items) != 1 {
		t.Errorf("expected 1 track item, got %d", len(result.Tracks.Items))
	}
}

func TestPlaylist_JSON(t *testing.T) {
	data := `{
		"id": "37i9dQZF1DXcBWIGoYBM5M",
		"name": "Today's Top Hits",
		"description": "The hottest tracks right now",
		"public": true,
		"collaborative": false,
		"type": "playlist",
		"uri": "spotify:playlist:37i9dQZF1DXcBWIGoYBM5M",
		"owner": {
			"display_name": "Spotify",
			"id": "spotify",
			"type": "user"
		},
		"followers": {"total": 12345678},
		"tracks": {
			"href": "https://api.spotify.com/v1/playlists/abc/tracks",
			"items": [{"track": {"id": "1", "name": "Song"}}],
			"total": 50
		},
		"images": [{"height": 640, "url": "https://i.scdn.co/image/abc", "width": 640}],
		"external_urls": {"spotify": "https://open.spotify.com/playlist/abc"}
	}`
	var playlist Playlist
	if err := json.Unmarshal([]byte(data), &playlist); err != nil {
		t.Fatalf("failed to unmarshal playlist: %v", err)
	}
	if playlist.Name != "Today's Top Hits" {
		t.Errorf("expected Today's Top Hits, got %s", playlist.Name)
	}
	if playlist.Owner.DisplayName != "Spotify" {
		t.Errorf("expected Spotify, got %s", playlist.Owner.DisplayName)
	}
	if len(playlist.Tracks.Items) != 1 {
		t.Errorf("expected 1 track item, got %d", len(playlist.Tracks.Items))
	}
}

func TestPlaylistTracks_JSON(t *testing.T) {
	data := `{
		"href": "https://api.spotify.com/v1/playlists/abc/tracks",
		"items": [
			{"track": {"id": "1", "name": "Song 1"}, "added_at": "2024-01-01T00:00:00Z"}
		],
		"total": 1
	}`
	var pt PlaylistTracks
	if err := json.Unmarshal([]byte(data), &pt); err != nil {
		t.Fatalf("failed to unmarshal playlist tracks: %v", err)
	}
	if len(pt.Items) != 1 {
		t.Errorf("expected 1 item, got %d", len(pt.Items))
	}
	if pt.Items[0].Track.Name != "Song 1" {
		t.Errorf("expected Song 1, got %s", pt.Items[0].Track.Name)
	}
}

func TestPlaylists_JSON(t *testing.T) {
	data := `{
		"href": "https://api.spotify.com/v1/me/playlists",
		"items": [{"id": "1", "name": "Playlist 1"}],
		"total": 1
	}`
	var playlists Playlists
	if err := json.Unmarshal([]byte(data), &playlists); err != nil {
		t.Fatalf("failed to unmarshal playlists: %v", err)
	}
	if len(playlists.Items) != 1 {
		t.Errorf("expected 1 item, got %d", len(playlists.Items))
	}
}

func TestProfile_JSON(t *testing.T) {
	data := `{
		"display_name": "John Doe",
		"id": "johnspotify",
		"type": "user",
		"uri": "spotify:user:johnspotify",
		"followers": {"total": 100},
		"images": [{"height": 64, "url": "https://i.scdn.co/image/abc", "width": 64}]
	}`
	var profile Profile
	if err := json.Unmarshal([]byte(data), &profile); err != nil {
		t.Fatalf("failed to unmarshal profile: %v", err)
	}
	if profile.DisplayName != "John Doe" {
		t.Errorf("expected John Doe, got %s", profile.DisplayName)
	}
	if profile.Followers.Total != 100 {
		t.Errorf("expected 100 followers, got %d", profile.Followers.Total)
	}
}

func TestAudioFeatures_JSON(t *testing.T) {
	data := `{
		"danceability": 0.5,
		"energy": 0.8,
		"key": 2,
		"loudness": -8.5,
		"mode": 1,
		"speechiness": 0.05,
		"acousticness": 0.1,
		"instrumentalness": 0.0,
		"liveness": 0.15,
		"valence": 0.3,
		"tempo": 120.0,
		"type": "audio_features",
		"id": "6rqhFgbbKwnb9MLmUQDhG6",
		"uri": "spotify:track:6rqhFgbbKwnb9MLmUQDhG6",
		"duration_ms": 383000,
		"time_signature": 4
	}`
	var af AudioFeatures
	if err := json.Unmarshal([]byte(data), &af); err != nil {
		t.Fatalf("failed to unmarshal audio features: %v", err)
	}
	if af.Danceability != 0.5 {
		t.Errorf("expected danceability 0.5, got %f", af.Danceability)
	}
	if af.Tempo != 120.0 {
		t.Errorf("expected tempo 120, got %f", af.Tempo)
	}
}

func TestAudioFeaturesResult_JSON(t *testing.T) {
	data := `{
		"audio_features": [
			{"id": "1", "danceability": 0.5},
			{"id": "2", "danceability": 0.8}
		]
	}`
	var afr AudioFeaturesResult
	if err := json.Unmarshal([]byte(data), &afr); err != nil {
		t.Fatalf("failed to unmarshal audio features result: %v", err)
	}
	if len(afr.AudioFeatures) != 2 {
		t.Errorf("expected 2 features, got %d", len(afr.AudioFeatures))
	}
}

func TestPlayingTrack_JSON(t *testing.T) {
	data := `{
		"is_playing": true,
		"progress_ms": 123000,
		"item": {"id": "6rqhFgbbKwnb9MLmUQDhG6", "name": "Paranoid Android"},
		"currently_playing_type": "track",
		"context": {"type": "playlist", "uri": "spotify:playlist:abc"}
	}`
	var pt PlayingTrack
	if err := json.Unmarshal([]byte(data), &pt); err != nil {
		t.Fatalf("failed to unmarshal playing track: %v", err)
	}
	if !pt.IsPlaying {
		t.Errorf("expected is_playing true")
	}
	if pt.Item.Name != "Paranoid Android" {
		t.Errorf("expected Paranoid Android, got %s", pt.Item.Name)
	}
}

func TestSavedTracks_JSON(t *testing.T) {
	data := `{
		"href": "https://api.spotify.com/v1/me/tracks",
		"items": [
			{"added_at": "2024-01-01T00:00:00Z", "track": {"id": "1", "name": "Saved Track"}}
		],
		"total": 1
	}`
	var st SavedTracks
	if err := json.Unmarshal([]byte(data), &st); err != nil {
		t.Fatalf("failed to unmarshal saved tracks: %v", err)
	}
	if len(st.Items) != 1 {
		t.Errorf("expected 1 item, got %d", len(st.Items))
	}
}

func TestSavedAlbums_JSON(t *testing.T) {
	data := `{
		"href": "https://api.spotify.com/v1/me/albums",
		"items": [
			{"added_at": "2024-01-01T00:00:00Z", "album": {"id": "1", "name": "Saved Album"}}
		],
		"total": 1
	}`
	var sa SavedAlbums
	if err := json.Unmarshal([]byte(data), &sa); err != nil {
		t.Fatalf("failed to unmarshal saved albums: %v", err)
	}
	if len(sa.Items) != 1 {
		t.Errorf("expected 1 item, got %d", len(sa.Items))
	}
}

func TestPlaylistAddTracks_JSON(t *testing.T) {
	data := `{"snapshot_id": "abc123"}`
	var pat PlaylistAddTracks
	if err := json.Unmarshal([]byte(data), &pat); err != nil {
		t.Fatalf("failed to unmarshal playlist add tracks: %v", err)
	}
	if pat.SnapshotID != "abc123" {
		t.Errorf("expected abc123, got %s", pat.SnapshotID)
	}
}

func TestRecentlyPlayedTracks_JSON(t *testing.T) {
	data := `{
		"href": "https://api.spotify.com/v1/me/player/recently-played",
		"items": [
			{"track": {"id": "1", "name": "Recent Track"}, "played_at": "2024-01-01T00:00:00Z"}
		],
		"limit": 1,
		"next": "https://api.spotify.com/v1/me/player/recently-played?after=123"
	}`
	var rpt RecentlyPlayedTracks
	if err := json.Unmarshal([]byte(data), &rpt); err != nil {
		t.Fatalf("failed to unmarshal recently played: %v", err)
	}
	if len(rpt.Items) != 1 {
		t.Errorf("expected 1 item, got %d", len(rpt.Items))
	}
}

func TestError_JSON(t *testing.T) {
	data := `{"error": {"status": 401, "message": "Invalid access token"}}`
	var spf SpfError
	if err := json.Unmarshal([]byte(data), &spf); err != nil {
		t.Fatalf("failed to unmarshal error: %v", err)
	}
	if spf.Error.Status != 401 {
		t.Errorf("expected 401, got %d", spf.Error.Status)
	}
	if spf.Error.Message != "Invalid access token" {
		t.Errorf("expected 'Invalid access token', got '%s'", spf.Error.Message)
	}
}

func TestAudioAnalysis_JSON(t *testing.T) {
	data := `{
		"meta": {"analyzer_version": "4.0.0", "status": 0},
		"track": {"duration": 240.0, "tempo": 120.0, "key": 2, "mode": 1, "time_signature": 4, "loudness": -8.5, "codestring": "eJw"},
		"bars": [{"start": 0.0, "duration": 2.0, "confidence": 1.0}],
		"beats": [{"start": 0.0, "duration": 1.0, "confidence": 1.0}],
		"sections": [{"start": 0.0, "duration": 30.0, "confidence": 1.0, "tempo": 120.0, "key": 2, "mode": 1, "time_signature": 4, "loudness": -8.5}],
		"segments": [{"start": 0.0, "duration": 0.5, "confidence": 1.0, "loudness_start": -60.0, "loudness_max": -10.0, "pitches": [1,0,0,0,0,0,0,0,0,0,0,0], "timbre": [0,0,0,0,0,0,0,0,0,0,0,0]}],
		"tatums": [{"start": 0.0, "duration": 0.25, "confidence": 1.0}]
	}`
	var aa AudioAnalysis
	if err := json.Unmarshal([]byte(data), &aa); err != nil {
		t.Fatalf("failed to unmarshal audio analysis: %v", err)
	}
	if aa.Track.Duration != 240.0 {
		t.Errorf("expected duration 240.0, got %f", aa.Track.Duration)
	}
	if len(aa.Bars) != 1 {
		t.Errorf("expected 1 bar, got %d", len(aa.Bars))
	}
}

func TestClientCreation(t *testing.T) {
	client := NewClient("test-token")
	if client == nil {
		t.Fatal("expected non-nil client")
	}
	if client.Album == nil {
		t.Error("expected non-nil Album service")
	}
	if client.Artist == nil {
		t.Error("expected non-nil Artist service")
	}
	if client.Track == nil {
		t.Error("expected non-nil Track service")
	}
	if client.Search == nil {
		t.Error("expected non-nil Search service")
	}
	if client.Playlist == nil {
		t.Error("expected non-nil Playlist service")
	}
	if client.Player == nil {
		t.Error("expected non-nil Player service")
	}
	if client.Library == nil {
		t.Error("expected non-nil Library service")
	}
	if client.User == nil {
		t.Error("expected non-nil User service")
	}
	if client.base == nil {
		t.Error("expected non-nil base sling")
	}
}

func TestRecentlyPlayedTracksParams_URLTags(t *testing.T) {
	p := RecentlyPlayedTracksParams{Type: "tracks", Limit: 20, After: 123456}
	if p.Limit != 20 {
		t.Errorf("expected Limit 20, got %d", p.Limit)
	}
	// The url tags should be: type, limit, after
	// Previously limit had url:"type" (copy-paste bug)
	if p.Type != "tracks" {
		t.Errorf("expected tracks, got %s", p.Type)
	}
}
