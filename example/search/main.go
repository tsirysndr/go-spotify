package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-spotify"
)

func main() {
	// Replace with your ACCESS TOKEN
	token := "BQBVlmqz3pvLjgIiimObr-N5lOAgzqFulsG1uYAPFoWV1lQn3T_61L3XSjDEgGup_aSE7ZhpktnFFjaZt2fNS1_YZDw0HOHuQcPFWEpIduHzanssGUsIlYi-dHPWASVk9-xTtUxyCIUfTHd-dLN9zVy6PeSFArAa9clB_Mo5TDGzu54-EK5NkUg5vIqFRFdoYJYQP7XGtfdLJdGkyab4qHgyezpsnD3x_tdmbBGaVXxDgS5-froHJgnabb5PyhcKIvvrJxMGreClCSRSdfb0JV9FjMr_9Q"
	client := spotify.NewClient(token)
	r, _ := client.Search.Get("Muse", "track,artist", 10, 0)
	result, _ := json.Marshal(r)
	fmt.Println(string(result))
}
