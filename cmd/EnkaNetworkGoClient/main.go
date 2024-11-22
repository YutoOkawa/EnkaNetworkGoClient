package main

import (
	"EnkaNetworkGoClient/pkg/client"
	"fmt"
	"os"
)

func main() {
	playerID := os.Getenv("PLAYER_ID")

	client := client.NewClient()

	enkaNetworkRes, err := client.GetAllData(playerID)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(*enkaNetworkRes)

	playerData, err := client.GetPlayerInfo(playerID)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(*playerData)
}
