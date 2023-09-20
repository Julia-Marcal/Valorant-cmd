package cmd

import (
	"fmt"

	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

const (
	Name = iota
	Tag
)

type ValorantModel struct {
	Inputs  []textinput.Model
	focused int
	err     error
}

func (p ValorantModel) Init() tea.Cmd {
	return nil
}

func NewCmdValorant() ValorantModel {
	var Inputs []textinput.Model = make([]textinput.Model, 2)

	Inputs[Name] = textinput.New()
	Inputs[Name].Placeholder = "UserName"
	Inputs[Name].Focus()
	Inputs[Name].CharLimit = 16
	Inputs[Name].Width = 30

	Inputs[Tag] = textinput.New()
	Inputs[Tag].Placeholder = "Tag"
	Inputs[Tag].Focus()
	Inputs[Tag].CharLimit = 5
	Inputs[Tag].Width = 15

	return ValorantModel{
		Inputs:  Inputs,
		focused: 0,
		err:     nil,
	}

}

func (p ValorantModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(p.Inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if p.focused == len(p.Inputs)-1 {
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
		for i := range p.Inputs {
			p.Inputs[i].Blur()
		}
		p.Inputs[p.focused].Focus()

	case errMsg:
		p.err = msg
		return p, nil
	}

	for i := range p.Inputs {
		p.Inputs[i], cmds[i] = p.Inputs[i].Update(msg)
	}
	return p, tea.Batch(cmds...)
}

func (p ValorantModel) View() string {
	s := strings.Builder{}
	s.WriteString("What is your nickName and Tag?")
	return fmt.Sprintf(
		`
		%s
		%s
		
		`,
		p.Inputs[Name].View(),
		p.Inputs[Tag].View(),
	) + "\n"
}

func (p *ValorantModel) nextInput() {
	p.focused = (p.focused + 1) % len(p.Inputs)
}

func (p *ValorantModel) prevInput() {
	p.focused--
	if p.focused < 0 {
		p.focused = len(p.Inputs) - 1
	}
}
