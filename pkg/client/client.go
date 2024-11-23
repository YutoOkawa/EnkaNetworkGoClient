package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/YutoOkawa/EnkaNetworkGoClient/v0/pkg/model"
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

func (c *Client) GetCharacterData() (map[string]model.CharacterData, error) {
	url := "https://raw.githubusercontent.com/EnkaNetwork/API-docs/refs/heads/master/store/characters.json"

	body, err := c.fetchData(url)
	if err != nil {
		return nil, err
	}

	var characterData map[string]model.CharacterData
	err = json.Unmarshal(body, &characterData)
	if err != nil {
		return nil, err
	}

	return characterData, nil
}

func (c *Client) GetLocalizationData() (map[string]map[string]string, error) {
	url := "https://raw.githubusercontent.com/EnkaNetwork/API-docs/refs/heads/master/store/loc.json"

	body, err := c.fetchData(url)
	if err != nil {
		return nil, err
	}

	var localizationData map[string]map[string]string
	err = json.Unmarshal(body, &localizationData)
	if err != nil {
		return nil, err
	}

	return localizationData, nil
}

func (c *Client) GetAllData(playerId string) (*model.EnkaNetworkResponse, error) {
	// TODO: Optional endpoint
	url := fmt.Sprintf("https://enka.network/api/uid/%s", playerId)

	body, err := c.fetchData(url)
	if err != nil {
		return nil, err
	}

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
	if err != nil {
		return nil, err
	}

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

	if res.StatusCode >= 400 && res.StatusCode < 500 {
		return nil, fmt.Errorf("Client error: %d", res.StatusCode)
	}

	if res.StatusCode >= 500 {
		return nil, fmt.Errorf("Server error: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
