package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// MODEL DATA

type (
	errMsg error
)

type ValorantCmd struct {
	text      string
	inputText textinput.Model
	err       error
}

func newCmdValorant(text string) ValorantCmd {
	input := textinput.New()
	input.Placeholder = "Hello Radiant"
	input.Focus()
	input.CharLimit = 30
	input.Width = 20

	return ValorantCmd{
		text:      text,
		inputText: input,
		err:       nil,
	}

}

func (p ValorantCmd) Init() tea.Cmd {
	return textinput.Blink
}

func (p ValorantCmd) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return p, tea.Quit
		}

	case errMsg:
		p.err = msg
		return p, nil
	}

	p.inputText, cmd = p.inputText.Update(msg)
	return p, cmd
}

func (p ValorantCmd) View() string {
	return fmt.Sprintf(
		"Program still being build\n\n",
		p.inputText.View(),
		"(ctrl + c to quit)",
	) + "\n"
}
