package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-spotify"
)

func main() {
	// Replace with your ACCESS TOKEN
	token := "BQCBHqupQy8WQDKB6FYyJf_c3ff6PgxTC-HRofjyeJPxcUO_BF_9-BHxuAdPpKzn_Jop-zs2oT_u1Q4t2Cxi4HlD0D5e5G1wwUc2tE67dfoT8FBYGiXygC6wOmf6dDRs0qRDm2qCsntoNLM9lnqe5XofaY2OCVAN0Q"
	client := spotify.NewClient(token)
	result, _ := client.Album.List("382ObEPsp2rxGrnsizN5TX,1A2GTWGtFfWp7KSQTwWOyo,2noRn2Aes5aoNVsU6iWThc")
	albums, _ := json.Marshal(result)
	fmt.Println(string(albums))
}
