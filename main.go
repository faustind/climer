/*
climer is a simple timer that runs from the command line.

Usage of climer:

	climer -d duration

The flags are:

	-d
	    The duration of the timer as a string of unsigned
	    decimal numbers, each with optional fraction and
	    a unit suffix, such as "300ms", "1.5h" or "2h45m".
	    Expected time units are "s", "m", "h".
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/faustind/climer/font"
)

type window struct {
	width  int
	height int
}

type model struct {
	window window
	hour   int
	min    int
	sec    int
	done   bool
}

func initialModel(hour, min, sec int) model {
	return model{hour: hour, min: min, sec: sec, done: false}
}

type tickMsg time.Time

func Tick() tea.Cmd {
	return tea.Tick(1*time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Init() tea.Cmd {
	return Tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.window.width, m.window.height = msg.Width, msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	case tickMsg:
		if m.hour == 0 && m.min == 0 && m.sec == 0 {
			m.done = true
		}

		if m.done {
			break
		}

		hour, min, sec := m.hour, m.min, m.sec-1
		if sec < 0 {
			min, sec = min-1, 59
		}

		if min < 0 {
			hour, min = hour-1, 59
		}

		m.hour, m.min, m.sec = hour, min, sec

		return m, Tick()
	}

	return m, nil
}

func (m model) View() string {
	timerStr := ""
	if m.hour > 0 {
		timerStr = fmt.Sprintf("%02d:", m.hour)
	}
	timerStr = timerStr + fmt.Sprintf("%02d:%02d", m.min, m.sec)
	timer := ""
	for _, c := range timerStr {
		timer = lipgloss.JoinHorizontal(lipgloss.Center, timer, font.DrawChar(c))
	}

	ui := lipgloss.Place(
		m.window.width, m.window.height,
		lipgloss.Center, lipgloss.Center,
		timer,
	)
	return ui
}

func main() {

	var d = flag.Duration("d", 1*time.Minute+30*time.Second, "Duration of the timer.")

	flag.Parse()

	seconds := int(d.Seconds())
	if seconds < 0 {
		fmt.Print("You already have no time left hehe\n")
		os.Exit(1)
	}

	minutes := seconds / 60
	seconds = seconds % 60
	hours := minutes / 60
	minutes = minutes % 60

	p := tea.NewProgram(initialModel(hours, minutes, seconds), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	os.Exit(0)
}
