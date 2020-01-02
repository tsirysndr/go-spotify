package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-spotify"
)

func main() {
	// Replace with your ACCESS TOKEN
	token := "BQDpg2UEcYn1SH1V97tQLGCM2fPBQ5Bh9gyA-sJoaliV_kxF3PDZmnR_NvGkYj4RwANN6341IkrjS56FJx282liaZdkVV6D8L8oAZft2ApaETA32hZYCKmPGYw4flyR7ye57u-unJ_JPJJ0bL-rwgycwbgu55kiwBVHrWkUt8uGgF-X8uoqm4WTsugWy3RsDAGJ4-LEpGWmtw9RjCK4VtQfA8qjIYDl5JkTx9ClYgtN0n6XWaK43FmLb91pM_iL-7xMfKKLXQe_qFRv-0vviDADxFhDKZA"
	client := spotify.NewClient(token)
	result, _ := client.Artist.GetAlbums("4tZwfgrHOc3mvqYlEYSvVi", "single,appears_on", "US", 10, 0)
	albums, _ := json.Marshal(result)
	fmt.Println(string(albums))
}
