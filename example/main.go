package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-spotify"
)

func main() {
	// Replace with your ACCESS TOKEN
	token := "BQD75iJHGK7OKnzYSs0Tw5MwGqwzofvx8eLm_QtURyK-WeyZZtnrIEOX7DdZwgVky4ItcDL7mvyuKmNOdptJ2x2rPMpckE85A6MfqVw9qcEGTTNkf5zC5mWKbN31P7klH7C1JlAbJCPH5i0gnlqdASJscpk_yes5nA"
	client := spotify.NewClient(token)
	result, _ := client.Track.ListTracks("6rqhFgbbKwnb9MLmUQDhG6")
	tracks, _ := json.Marshal(result)
	fmt.Println(string(tracks))
}
