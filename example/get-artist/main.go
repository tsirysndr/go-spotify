package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-spotify"
)

func main() {
	// Replace with your ACCESS TOKEN
	token := "BQAI89IusXScO6IIKM274REYEcprAeviM0DaNx1pMbJKKFa7rZ3riplCHMU48dIhYWJP-lJXvchmQ_gRMFsms4-gjhRqOtMGm8KUoUmMlwOY4hDBqS3mc6ioKHRi5xGoWPhYDaPhkQ76NNV5WnyySAh-QWnO4mfb7K-3-BpL5cBhj8cPfwmthEtyyuZj7LXCUnMvRJcC-sK9zzVyPmQHyqDVQFq1HdTiMHdZ_PqqVqM4abyaImDR-r2vIWtw_n38dtRiU0ostu-IGy6vKR2djQsZGHIQVQ"
	client := spotify.NewClient(token)
	result, _ := client.Artist.Get("4tZwfgrHOc3mvqYlEYSvVi")
	artist, _ := json.Marshal(result)
	fmt.Println(string(artist))
}
