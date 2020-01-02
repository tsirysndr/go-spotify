package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-spotify"
)

func main() {
	// Replace with your ACCESS TOKEN
	token := "BQBLuNIffs-u1YJz2wuco-J5AZKyXG8icxHtfQSn-KWvtH8_2TDKJacM8W2UXj3ZFeK7LSY0_pgZjiKtH71Jw20xvmgCIxdLDsEJuAsTBBbzwYAMxl6w_ucBFH8GHkmmxFaBDx2oxtkfIXADge5gLDj9fNREkJsuI1vdv3W1VeDhrZLAjRI3kGtdUC8Yyupet3RCmhKRnamSjMJA3fMddkqKJnnVA9UMhCnfqlSAstBraMahP575pQNdPvR_JHARjNlWz5XEUC063Zw2kq-T-RQOYLaYSw"
	client := spotify.NewClient(token)
	result, _ := client.Track.AudioAnalysis("11dFghVXANMlKmJXsNCbNl")
	audio_analysis, _ := json.Marshal(result)
	fmt.Println(string(audio_analysis))
}
