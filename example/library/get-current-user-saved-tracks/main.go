package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-spotify"
)

func main() {
	// Replace with your ACCESS TOKEN
	token := "BQB0-DfMFJWuIc3ZvH_P8Rj2bGC1HhIclKh43qf7YU4_GYbqQsBnL4ryEOcacOmpueYezCndmyx_4eSn6Pb6dNyL7JKa9JC5M75psW-lw_Um-Jsp7Hp92tEhPYiuDI57o7ysIWvXOkBWksNSlUdggynuNrxYBzAWNFTb4FD3pG-suwf5GDm1OYh8vjs700wRzjEVsGHAyncnyvthXKjMja2gq-_KgRwWAsIugjpW9ZdK-Y_jxzs-RW8ErpfjvMomB-2B19Dn88b9Qqd5uB8-N7IdC6-r1g"
	client := spotify.NewClient(token)
	result, _ := client.Library.GetCurrentUserSavedTracks(10, 0)
	tracks, _ := json.Marshal(result)
	fmt.Println(string(tracks))
}
