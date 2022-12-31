package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-spotify"
)

func main() {
	// Replace with your ACCESS TOKEN
	token := "BQDHY1TIdHbAev8CC5NIH0oExrDAq6Eijl4T91dLYmm84v9-XX_i9NPipGZadwlD8WaSsrFeXgvTmZY8VetPXTtQMIx_9t5FGjKBlha7mHTRx3wmWdHAM48A2IZJHoXj3yFueyBmubhpjhqMoUhZcfObHC2sxHTy0VJgbbr79gLGxpCUhZUD3j6AZ7ammxG0_Nsfcaf8u3fPRPICennDng3Qeh3ZHI3OWpy2zkJZecFUJsOfCZMabDRf-vXp1kyL"
	client := spotify.NewClient(token)
	result, _ := client.Playlist.CreatePlaylist("am14idaj3jbs5u8b5tskn286i", "My Playlist", "Another awesome playlist", true)
	snapshot, _ := json.Marshal(result)
	fmt.Println(string(snapshot))

}
