package cmd

import (
	"fmt"

	"github.com/Julia-Marcal/valorant-cmd/fetch"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

const (
	name = iota
	tag
)

type ValorantModel struct {
	inputs  []textinput.Model
	focused int
	err     error
}

func (p ValorantModel) Init() tea.Cmd {
	return textinput.Blink
}

func NewCmdValorant() ValorantModel {
	var inputs []textinput.Model = make([]textinput.Model, 2)

	inputs[name] = textinput.New()
	inputs[name].Placeholder = "Username"
	inputs[name].Focus()
	inputs[name].CharLimit = 16
	inputs[name].Width = 30

	inputs[tag] = textinput.New()
	inputs[tag].Placeholder = "Tag"
	inputs[tag].Focus()
	inputs[tag].CharLimit = 5
	inputs[tag].Width = 15

	return ValorantModel{
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}

}

func (p ValorantModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(p.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if p.focused == len(p.inputs)-1 {
				return p, tea.Quit
			}
			p.nextInput()
		case tea.KeyCtrlC, tea.KeyEsc:
			return p, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP:
			p.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			p.nextInput()
		}
		for i := range p.inputs {
			p.inputs[i].Blur()
		}
		p.inputs[p.focused].Focus()

	case errMsg:
		p.err = msg
		return p, nil
	}

	for i := range p.inputs {
		p.inputs[i], cmds[i] = p.inputs[i].Update(msg)
	}
	return p, tea.Batch(cmds...)
}

func (p ValorantModel) View() string {
	return fmt.Sprintf(
		`
		%s
		%s
		
		`,
		p.inputs[name].View(),
		p.inputs[tag].View(),
	) + "\n"
}

func (p *ValorantModel) nextInput() {
	p.focused = (p.focused + 1) % len(p.inputs)
}

func (p *ValorantModel) prevInput() {
	p.focused--
	if p.focused < 0 {
		p.focused = len(p.inputs) - 1
	}
}

func FetchAccount(p *ValorantModel) (fetch.AccountInfo, error) {
	accountInfo, err := fetch.AccountInformation(p.inputs[name].Value(), p.inputs[tag].Value())
	if err != nil {
		return fetch.AccountInfo{}, err
	}
	return *accountInfo, nil
}
