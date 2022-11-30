package font

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// DrawChar returns the argument character in bigger size.
func DrawChar(c rune) string {
	b := strings.Builder{}
	for _, row := range smallFont[c] {
		for _, char := range row {
			s := lipgloss.NewStyle().SetString(" ")
			if char == rune('#') {
				s = s.Background(lipgloss.Color("#43BF6D"))
			}
			b.WriteString(s.String())
		}
		b.WriteRune('\n')
	}

	return b.String()
}
