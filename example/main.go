package main

import (
	"fmt"

	"github.com/tsirysndr/go-spotify"
)

func main() {
	// Replace with your ACCESS TOKEN
	token := "BQAnXnzb2stOGUJ4qdp3H800Dc-NJ_Gb-cXzCnR9BuyNQvT_Eyk2D2_otMKaEGeguiPDLQs1Z2Y8ULuGjZHGpbOXcsYtDgr8Jic_skIJbaYCkGuaZIcnU48dADcZ-R0TS5DrLertSu1HmNVsD5WOqyqbnqIl1Dc__A"
	client := spotify.NewClient(token)
	result, err := client.Track.ListTracks("6rqhFgbbKwnb9MLmUQDhG6")
	fmt.Println(result, err)
}
