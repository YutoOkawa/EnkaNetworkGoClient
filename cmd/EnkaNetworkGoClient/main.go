package main

import (
	"EnkaNetworkGoClient/pkg/client"
	"EnkaNetworkGoClient/pkg/model"
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
	fmt.Println("level:", enkaNetworkRes.AvatarInfoList[0].PropMap[model.PROP_TYPE_LEVEL].Val)

	playerData, err := client.GetPlayerInfo(playerID)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(*playerData)

	characterData, err := client.GetCharacterData()
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(characterData)

	localizationData, err := client.GetLocalizationData()
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(localizationData)
}
