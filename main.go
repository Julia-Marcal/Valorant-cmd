package main

import (
	"fmt"
	"os"

	cmd "github.com/Julia-Marcal/valorant-cmd/cmd"
	val "github.com/Julia-Marcal/valorant-cmd/fetch"
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

		reg, lvl, img, err_fetch := val.FetchAccount(nameValue, tagValue)

		if err_fetch != nil {
			fmt.Println("Oh no:", err_fetch)
			os.Exit(1)
		}

		fmt.Printf("Account region: %s\n", reg)
		fmt.Printf("Account level: %d\n", lvl)
		fmt.Printf("Account card: %s\n", img)
	}

}
