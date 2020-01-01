package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-spotify"
)

func main() {
	// token
	token := "BQBiktRNvkinqA2DNtvoZmTRtZ9u3Vl6sDmSph8nrwTWcMyeom1gA6KL-hCAnEAswq4ZWIvf4Gr88CkJn6wcMmmHnZVKovM_VdDW8bk4LPOd5ayUkTsxzlVvaMvXy3YmFq1qToACa9s8VHJO3onAIvsBBTtwex1EUw"
	client := spotify.NewClient(token)
	// User id
	id := "lyp9kipowzwqehy4ppj5owl7p"
	res, _ := client.User.GetProfile(id)
	user, _ := json.Marshal(res)
	fmt.Println(string(user))
}
