package main

import (
	cmd "github.com/Julia-Marcal/valorant-cmd/cmd"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(
		cmd.NewCmdValorant("This app is under construction"),
	)
	if err := p.Start(); err != nil {
		panic(err)
	}
}
