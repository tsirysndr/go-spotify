package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-spotify"
)

func main() {
	// Replace with your ACCESS TOKEN
	token := "BQBVXTjKQWizqv_gapNizMTwYr3UsjqFErufF1UrhR4OCl5y-ygkhWrR6s4dKOhFEqw9YruOhDfK6GXldbJun-bkSZZrDkn-Gyl3mmeFeK1-mVaHb8SewS0a-X5S5BwZnr8KTUVz3EfF1aX7ET_9-faBQJSzS9PMaRxUNcRCEQdBvK50DjbVxQw4QY8nfq5z99MOT-zfxEZ8DqhL4MTPVGVW9FRp4CWGNDho96v4dQb2rEQ-HBs5PAUzodMuWXBO"
	client := spotify.NewClient(token)
	result, _ := client.Playlist.AddTracksToPlaylists("1LoTXrTzdTPb12uhmF03S8", "spotify:track:6AXpjInvhGMq7wbDk0iHGP")
	snapshot, _ := json.Marshal(result)
	fmt.Println(string(snapshot))

}
