package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-spotify"
)

func main() {
	// Replace with your ACCESS TOKEN
	token := "BQBPksHCt622hhSLEKPej4QEds6tYaSOXg5yAC5z58wo3J_M138yQEfOeAhETPNOB_dCwMACzGqmazEXjNS4E-vKnvMeqWEhZ5Vt0RLmmjoJiBms6qa5I6HMQ_YW-pp7HpBbjNdT7THIpfZExYb4is5bNK-mRNzBY6F2Fhwgq0ScWgK0dDeMUmGfEOHUadtfpu7WG6QvSIZ9HifEEioR7Lsi9FybMAUsmyUktyQkDyP1lUy9QbgKzjCg3cOYj98_"
	client := spotify.NewClient(token)
	result, _ := client.Playlist.AddTracksToPlaylists("1LoTXrTzdTPb12uhmF03S8", "spotify:track:6AXpjInvhGMq7wbDk0iHGP")
	snapshot, _ := json.Marshal(result)
	fmt.Println(string(snapshot))

}
