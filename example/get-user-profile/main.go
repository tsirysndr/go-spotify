package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-spotify"
)

func main() {
	// token
	token := "BQDid21lB1igrlfPLB5pmWWr31BSeeKgJU8YJf1lQju6MhJLlzn097bgULQA3ATzEzhHlgzHjGCHxn9lKpIuBsO907agI9r8Bzka0OQtn1DlMN5Y-cuZPEuNQwNgxE1rEjLomYVw6fE5NmTBq4bVegeFnM3pGVqcyQ"
	client := spotify.NewClient(token)
	// User id
	id := "lyp9kipowzwqehy4ppj5owl7p"
	res, _ := client.User.GetProfile(id)
	user, _ := json.Marshal(res)
	fmt.Println(string(user))
}
