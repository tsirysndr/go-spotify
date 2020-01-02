package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-spotify"
)

func main() {
	// Replace with your ACCESS TOKEN
	token := "BQAmJGDyhpY8ucAaWV2EWg-plZFozoRwFtLgZV8iPIjFLpDSksCwUz6yyuCtCDTT1Ho6aV9PSutt9xT5Jwe8hU_R1jijIeZJ0yHlZGHbhqb_w32t5wpVxpWCzWO-KtyQkvoxK3BMo8jL7KyA9YgXsgcyJWtnye2TMTzHiQsfBB3Ef-IU42XolNTNz2YWspgdAkY0kaYAPqRPnEGk8rD8ZR-9IqfFSJlbET3WsDMCINEurugt9_Vy0qrsvt7bM2OaawvDKcsDvXCU-zXM5YQMxWd6XKIixg"
	client := spotify.NewClient(token)
	r, _ := client.Search.Get("Muse", "track,artist", 10, 0)
	result, _ := json.Marshal(r)
	fmt.Println(string(result))
}
