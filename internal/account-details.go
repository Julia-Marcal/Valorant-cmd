package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func AccountInformation(name string, tag string) (*AccountInfo, error) {
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/account/%s/%s", name, tag)
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var result AccountResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &AccountInfo{
		Puuid:        result.Data.Puuid,
		AccountLevel: result.Data.AccountLevel,
		Region:       result.Data.Region,
		Name:         result.Data.Name,
		Tag:          result.Data.Tag,
		Large:        result.Data.Card.Large,
	}, nil
}
