package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type window struct {
	width  int
	height int
}

type Model struct {
	window window
	min    int
	sec    int
	done   bool
}

func initialModel() Model {
	return Model{min: 1, sec: 30, done: false}
}

type TickMsg time.Time

func Tick() tea.Cmd {
	return tea.Tick(1*time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m Model) Init() tea.Cmd {
	return Tick()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.window.width, m.window.height = msg.Width, msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	case TickMsg:
		if m.min == 0 && m.sec == 0 {
			m.done = true
		}

		if m.done {
			break
		}

		min, sec := m.min, m.sec-1
		if sec < 0 {
			min, sec = min-1, 59
		}

		m.min, m.sec = min, sec

		return m, Tick()
	}

	return m, nil
}

func (m Model) View() string {
	timerStr := fmt.Sprintf("%02d:%02d", m.min, m.sec)
	timer := ""
	for _, c := range timerStr {
		timer = lipgloss.JoinHorizontal(lipgloss.Center, timer, drawChar(c))
	}

	ui := lipgloss.Place(
		m.window.width, m.window.height,
		lipgloss.Center, lipgloss.Center,
		timer,
	)
	return ui
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
	}
}
