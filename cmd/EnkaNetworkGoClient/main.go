package main

import (
	"context"
	"fmt"
	"os"

	"github.com/YutoOkawa/EnkaNetworkGoClient/v0/pkg/client"
	"github.com/YutoOkawa/EnkaNetworkGoClient/v0/pkg/model"
)

func main() {
	playerID := os.Getenv("PLAYER_ID")

	client := client.NewClient()

	enkaNetworkRes, err := client.GetAllData(context.Background(), playerID)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println("level:", enkaNetworkRes.AvatarInfoList[0].PropMap[model.PROP_TYPE_LEVEL].Val)

	playerData, err := client.GetPlayerInfo(context.Background(), playerID)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(*playerData)

	characterData, err := client.GetCharacterData(context.Background())
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(characterData)

	localizationData, err := client.GetLocalizationData(context.Background())
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(localizationData)
}
