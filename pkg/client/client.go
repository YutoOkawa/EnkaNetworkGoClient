package client

import (
	"EnkaNetworkGoClient/pkg/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// TODO: cache
// TODO: retry (only for 5xx)
type Client struct {
	client *http.Client
}

func NewClient() *Client {
	return &Client{
		client: &http.Client{
			Timeout: 10 * time.Second, // TODO: Optional timeout
		},
	}
}

func (c *Client) GetAllData(playerId string) (*model.EnkaNetworkResponse, error) {
	// TODO: Optional endpoint
	url := fmt.Sprintf("https://enka.network/api/uid/%s", playerId)

	body, err := c.fetchData(url)

	var enkaNetworkRes model.EnkaNetworkResponse
	err = json.Unmarshal(body, &enkaNetworkRes)
	if err != nil {
		return nil, err
	}

	return &enkaNetworkRes, nil
}

func (c *Client) GetPlayerInfo(playerId string) (*model.PlayerInfo, error) {
	// TODO: Optional endpoint
	url := fmt.Sprintf("https://enka.network/api/uid/%s?info", playerId)

	body, err := c.fetchData(url)

	var playerInfo struct {
		PlayerInfo model.PlayerInfo `json:"playerInfo"`
	}
	err = json.Unmarshal(body, &playerInfo)
	if err != nil {
		return nil, err
	}

	return &playerInfo.PlayerInfo, nil
}

func (c *Client) fetchData(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "EnkaNetworkGoClient")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// TODO: status code handling
	// refer: https://github.com/EnkaNetwork/API-docs/blob/master/api_ja.md#httpレスポンスコード

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
