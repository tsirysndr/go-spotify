package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-spotify"
)

func main() {
	// Replace with your ACCESS TOKEN
	token := "BQB6N8XUQkf8vXEGfHO44HmIfqN_DvtnBF_3ISNZpGewv09SQcPPrGAJ2Mp2tldtShPT4U2QQXorBkV8tk4rKShPcNuup2q5INA0jeLpMUpjyBVbxMSBUCEyKTEVCnCbWayAVRV-ZYAqSB6wm0PfhZkJCwIB8BK2I3Ve7skrEQHurw_ypsbrtO3--ZN0EI9iDq0Lvsw1ap6P2K9zAuQgAOdLShlD4Z_YqmUAvP8tnMjTzCXYtlgx2iNsvxxrLW7eDoosscA_dsxeecyllnsBL4Bk4ThBsA"
	client := spotify.NewClient(token)
	result, _ := client.Playlist.Get("37i9dQZF1DXcBWIGoYBM5M")
	playlist, _ := json.Marshal(result)
	fmt.Println(string(playlist))
}
