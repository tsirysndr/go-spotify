package main

import (
	"fmt"

	"github.com/tsirysndr/go-spotify"
)

func init() {
	// token
	token := "BQDCbs0OPy3VK24R8SM-FEjsD_G5GEBg__AR-EfFQC1C5NIrmXrs4IlQFzux-M_d5utUY6mh40K7GtdE"
	client := spotify.NewClient(token)
	// User id
	id := "6pcxz3cnzy501a95plbsc7f2l"
	res, err := client.User.UserProfile(id)

	fmt.Println(res, err)

}
