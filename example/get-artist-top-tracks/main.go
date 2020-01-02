package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-spotify"
)

func main() {
	// Replace with your ACCESS TOKEN
	token := "BQD0XIGP5Y5W3xBlef-l45aqhtHevYC2zr6GcinH8hVsmcytK5KQ7ESqqKuNBao5d7m7XzK0PG9DoRxi0JpgHeKfx_oCz8BbEBjOqPn3VyKgfxGXGTn-1jM0zx-j0tYySr-cmqC-m_0CkHmO8kPsvmn6MvuBKdMNsNSvU6XIl8tN_0xO6bEA-gTTqQLZqXOri69yFz8hOaB8fgvMsK_tSeh8P2z_lPHImZ0TlBdpBh2NtydEoTL3u5dTGU1h4aKTJR-dRwvQUQTeZZAqb8ay8IjJBVroPQ"
	client := spotify.NewClient(token)
	result, _ := client.Artist.GetTopTracks("4tZwfgrHOc3mvqYlEYSvVi", "US")
	tracks, _ := json.Marshal(result)
	fmt.Println(string(tracks))
}
