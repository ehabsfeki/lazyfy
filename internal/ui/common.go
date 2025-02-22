package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

var (
	primary           = termenv.ColorProfile().Color("205")
	secondary         = termenv.ColorProfile().Color("#89F0CB")
	gray              = termenv.ColorProfile().Color("#626262")
	midGray           = termenv.ColorProfile().Color("#4A4A4A")
	blue              = termenv.ColorProfile().Color("#F0FFFF")
	red               = termenv.ColorProfile().Color("#FF2D00")
	green             = termenv.ColorProfile().Color("#00FF00")
	docStyle          = lipgloss.NewStyle().Margin(1)
	titleStyle        = lipgloss.NewStyle().MarginLeft(0)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

const (
	star = string(`* `)
	plus = string(`+ `)
)

func boldPrimaryForeground(s string) string {
	return termenv.String(s).Foreground(primary).Bold().String()
}

func boldSecondaryForeground(s string) string {
	return termenv.String(s).Foreground(secondary).Bold().String()
}

func boldBlueForeground(s string) string {
	return termenv.String(s).Foreground(blue).Bold().String()
}

func boldRedForeground(s string) string {
	return termenv.String(s).Foreground(red).Bold().String()
}

func greenRedForeground(s string) string {
	return termenv.String(s).Foreground(green).Bold().String()
}

func blueForeground(s string) string {
	return termenv.String(s).Foreground(blue).String()
}

func blueFaintForeground(s string) string {
	return termenv.String(s).Foreground(blue).Faint().String()
}

func grayForeground(s string) string {
	return termenv.String(s).Foreground(gray).String()
}

func midGrayForeground(s string) string {
	return termenv.String(s).Foreground(midGray).String()
}

func faint(s string) string {
	return termenv.String(s).Faint().String()
}
