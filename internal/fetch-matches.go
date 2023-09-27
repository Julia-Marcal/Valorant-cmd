package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ApiResponse struct {
	Data []DataStruct `json:"data"`
}

var httpClient = &http.Client{Timeout: 10 * time.Second}

func FetchMatches(region string, puuid string) ([]int, map[string]string, []float64, string, error) {
	url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/by-puuid/lifetime/matches/%s/%s", region, puuid)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, nil, "", fmt.Errorf("creating request failed: %w", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, nil, nil, "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	var apiResponse ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, nil, nil, "", err
	}

	if len(apiResponse.Data) == 0 {
		return nil, nil, nil, "", fmt.Errorf("no data available")
	}

	if apiResponse.Data[0].Stats.Character.ID == "" || apiResponse.Data[0].Stats.Character.Name == "" {
		return nil, nil, nil, "", fmt.Errorf("Character data missing or incomplete")
	}

	return aggregateData(apiResponse)
}

func aggregateData(apiResponse ApiResponse) ([]int, map[string]string, []float64, string, error) {
	lastMatch := apiResponse.Data[0]
	totalKills, totalDeaths, totalAssists, numMatches := 0, 0, 0, len(apiResponse.Data)
	lookupCharacter := make(map[string]int)
	maxWins := 0
	var winningTeam, currentCharacter, bestCharacter string

	for _, match := range apiResponse.Data {
		character := match.Stats.Character.Name
		lookupCharacter[character] = 0
	}

	for i := 0; i < numMatches; i++ {
		totalKills += apiResponse.Data[i].Stats.Kills
		totalDeaths += apiResponse.Data[i].Stats.Deaths
		totalAssists += apiResponse.Data[i].Stats.Assists

		myTeam := apiResponse.Data[i].Stats.Team
		redScore := apiResponse.Data[i].Teams.Red
		blueScore := apiResponse.Data[i].Teams.Blue

		if redScore > blueScore {
			winningTeam = "Red"
		} else if blueScore > redScore {
			winningTeam = "Blue"
		} else {
			winningTeam = "Draw"
		}

		currentCharacter = apiResponse.Data[i].Stats.Character.Name

		if apiResponse.Data[i].Meta.Mode == "Competitive" && winningTeam == myTeam {
			if _, exists := lookupCharacter[currentCharacter]; exists {
				lookupCharacter[currentCharacter]++
			}
		}
	}

	for character, wins := range lookupCharacter {
		if wins > maxWins {
			maxWins = wins
			bestCharacter = character
		}
	}

	averageKills := float64(totalKills) / float64(numMatches)
	averageDeaths := float64(totalDeaths) / float64(numMatches)
	averageAssists := float64(totalAssists) / float64(numMatches)

	averages := []float64{averageKills, averageDeaths, averageAssists}

	intInfo := []int{lastMatch.Stats.Kills, lastMatch.Stats.Deaths, lastMatch.Stats.Assists}

	mapInfo := make(map[string]string)
	mapInfo["ID"] = lastMatch.Meta.Map.ID
	mapInfo["Name"] = lastMatch.Meta.Map.Name

	return intInfo, mapInfo, averages, bestCharacter, nil

}
