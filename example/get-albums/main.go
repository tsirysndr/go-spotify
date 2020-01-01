package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-spotify"
)

func main() {
	// Replace with your ACCESS TOKEN
	token := "BQBiktRNvkinqA2DNtvoZmTRtZ9u3Vl6sDmSph8nrwTWcMyeom1gA6KL-hCAnEAswq4ZWIvf4Gr88CkJn6wcMmmHnZVKovM_VdDW8bk4LPOd5ayUkTsxzlVvaMvXy3YmFq1qToACa9s8VHJO3onAIvsBBTtwex1EUw"
	client := spotify.NewClient(token)
	result, _ := client.Album.List("382ObEPsp2rxGrnsizN5TX,1A2GTWGtFfWp7KSQTwWOyo,2noRn2Aes5aoNVsU6iWThc")
	albums, _ := json.Marshal(result)
	fmt.Println(string(albums))
}
