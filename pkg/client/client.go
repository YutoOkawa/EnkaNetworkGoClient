package client

import (
	"EnkaNetworkGoClient/pkg"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

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

func (c *Client) Get(playerId string) (*pkg.EnkaNetworkResponse, error) {
	// TODO: Optional endpoint
	url := fmt.Sprintf("https://enka.network/api/uid/%s", playerId)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// TODO: Optional headers
	req.Header.Add("User-Agent", "EnkaNetworkGoClient")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var enkaNetworkRes pkg.EnkaNetworkResponse
	err = json.Unmarshal(body, &enkaNetworkRes)
	if err != nil {
		return nil, err
	}

	return &enkaNetworkRes, nil
}
