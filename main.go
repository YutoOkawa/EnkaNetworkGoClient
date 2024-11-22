package main

import (
	"EnkaNetworkGoClient/pkg"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	playerID := os.Getenv("PLAYER_ID")
	url := fmt.Sprintf("https://enka.network/api/uid/%s", playerID)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	req.Header.Add("User-Agent", "EnkaNetworkGoClient")

	res, err := client.Do(req)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	var enkaNetworkRes pkg.EnkaNetworkResponse
	err = json.Unmarshal(body, &enkaNetworkRes)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(enkaNetworkRes.PlayerInfo.NickName)
}
