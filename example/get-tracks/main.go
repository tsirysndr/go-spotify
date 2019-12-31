package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-spotify"
)

func main() {
	// Replace with your ACCESS TOKEN
	token := "BQDid21lB1igrlfPLB5pmWWr31BSeeKgJU8YJf1lQju6MhJLlzn097bgULQA3ATzEzhHlgzHjGCHxn9lKpIuBsO907agI9r8Bzka0OQtn1DlMN5Y-cuZPEuNQwNgxE1rEjLomYVw6fE5NmTBq4bVegeFnM3pGVqcyQ"
	client := spotify.NewClient(token)
	result, _ := client.Track.ListTracks("6rqhFgbbKwnb9MLmUQDhG6")
	tracks, _ := json.Marshal(result)
	fmt.Println(string(tracks))
}
