package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-spotify"
)

func main() {
	// Replace with your ACCESS TOKEN
	token := "BQCW2TeNbSRCXirfyaMDZJ9ijzPXUZh1MXRLLbk3WZCZN1oEf2LiFfjgw0GdOAtzF2BIBemgnqJ-s-6P_OsVNsBMOmwGrgLkjhsc9i8ekoLhiZjNF-CoKDfMTQgGO1ghwM505XVSDWdkA4C72Ja-vf5CGUWhWwo2REuSJVUJdFkS600wSc5ia9Z69LlrvAxk5Tc5sQHAkItTeixu5R8_Tqr5EMnfpxeFxP6bQRZYoJ6cI2h3l5Ei_aq1O_LTOoFklb-Qk7QXsLCrpOGS3P07w0vT-1aRjA"
	client := spotify.NewClient(token)
	result, _ := client.Player.GetRecentlyPlayed("track", 10, 1484811043508)
	recently_played, _ := json.Marshal(result)
	fmt.Println(string(recently_played))
}
