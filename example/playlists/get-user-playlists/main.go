package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-spotify"
)

func main() {
	// Replace with your ACCESS TOKEN
	token := "BQDXXiEU1Gs7J3oR1VkoLYiwbATiKj3L8Fif7hDJadN1uhJ64UC4U6PbUX74jEG0EzUNtanVDDNyDqY75N0RxbU0BFKHYrN87TeBiIM2UQyoY1_KsuESYznoc-aeGjsBB09_TKvfbq5zwlmZYIGFXzYaZN95-rm3z6N7KSIBDoG-hZ-1z60nfEPF9Kx_VreUCeFJ626fIZfMyTUukvfKpMsG_czFQH70Lp3tzzmjFh0r4k2uYUVl6kEFOGS4dhXS30ZyCAEjwgQLq2HJ6WEJ-VPDdWjh4w"
	client := spotify.NewClient(token)
	result, _ := client.Playlist.GetUserPlaylists("smedja", 10, 0)
	playlists, _ := json.Marshal(result)
	fmt.Println(string(playlists))
}
