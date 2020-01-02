package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-spotify"
)

func main() {
	// Replace with your ACCESS TOKEN
	token := "BQDfZknpyt7zzXmP7UYmlvL40RBKQZFywvCo71X1SxocQiEQlMXByzPdpqTiernA8QR5vC8jsv9v-AxyuV3ZvcP5ju5lbB51Hfppnanh2slRqXQGMsauHxOphKDe9uBMnNN5OF0WJsI-9ijD-rD5XvzwlA_UggbQ0XaYaLceXbUVuuHl1HwVeFPIvwEavnVU4EEyP112swyY2ngXxQDAX4kEFW5ArOT8SzP6gK37n7K58SPnn-tQYGeeybx3xtG038ZvELPIm1yEB76VBl6z1YtZI9xXCw"
	client := spotify.NewClient(token)
	result, _ := client.Playlist.GetTracks("37i9dQZF1DXcBWIGoYBM5M", 10, 0)
	tracks, _ := json.Marshal(result)
	fmt.Println(string(tracks))
}
