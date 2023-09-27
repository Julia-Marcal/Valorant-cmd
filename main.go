package main

import (
	"fmt"
	"os"

	cmd "github.com/Julia-Marcal/valorant-cmd/cmd"
	val "github.com/Julia-Marcal/valorant-cmd/internal"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(
		cmd.NewCmdValorant(),
	)

	m, err := p.Run()

	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	if m, ok := m.(cmd.ValorantModel); ok {
		nameValue := m.Inputs[cmd.Name].Value()
		tagValue := m.Inputs[cmd.Tag].Value()

		AccountFactory := &val.ConcreteAccountFetcherFactory{}
		AccountFetcher := AccountFactory.Create()

		reg, puuid, lvl, img, err_fetch := AccountFetcher.FetchAccount(nameValue, tagValue)

		if err_fetch != nil {
			fmt.Println("Oh no:", err_fetch)
			os.Exit(1)
		}

		MatchFactory := &val.ConcreteMatchesFetcherFactory{}
		MatchFetcher := MatchFactory.Create()

		matchInfo, mapINfo, averages, bestCharacter, matchErr := MatchFetcher.FetchMatchData(reg, puuid)

		if matchErr != nil {
			fmt.Println("Oh no:", err_fetch)
			os.Exit(1)
		}

		fmt.Printf("Account region: %s\n", reg)
		fmt.Printf("Account level: %d\n", lvl)
		fmt.Printf("Account card: %s\n", img)

		fmt.Printf("\n")

		fmt.Printf("Last map played: %s\n", mapINfo["Name"])
		fmt.Printf("Last match kill: %d\n", matchInfo[0])
		fmt.Printf("Last match deaths: %d\n", matchInfo[1])
		fmt.Printf("Last match assists: %d\n", matchInfo[2])

		fmt.Printf("\n")

		fmt.Printf("Average kills from last matches: %.0f\n", averages[0])
		fmt.Printf("Average death from last matches: %.0f\n", averages[1])
		fmt.Printf("Average assists from last matches: %.0f\n", averages[2])
		fmt.Printf("Your best character is: %s\n", bestCharacter)
	}

}
