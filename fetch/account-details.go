package fetch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AccountResponse struct {
	Status int `json:"status"`
	Data   struct {
		AccountLevel int    `json:"account_level"`
		Region       string `json:"region"`
		Name         string `json:"name"`
		Tag          string `json:"tag"`
		Card         struct {
			Large string `json:"large"`
		} `json:"card"`
	} `json:"data"`
}

type AccountInfo struct {
	AccountLevel int
	Region       string
	Name         string
	Tag          string
	Large        string
}

func AccountInformation(name string, tag string) (*AccountInfo, error) {
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/account/%s/%s", name, tag)
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response: %w", err)
	}

	var result AccountResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal response: %w", err)
	}

	return &AccountInfo{
		AccountLevel: result.Data.AccountLevel,
		Region:       result.Data.Region,
		Name:         result.Data.Name,
		Tag:          result.Data.Tag,
		Large:        result.Data.Card.Large,
	}, nil
}

func FetchAccount(Name string, Tag string) (string, int, string, error) {
	accountInfo, err := AccountInformation(Name, Tag)
	if err != nil {
		return accountInfo.Region, accountInfo.AccountLevel, accountInfo.Large, nil
	}
	return accountInfo.Region, accountInfo.AccountLevel, accountInfo.Large, nil
}
