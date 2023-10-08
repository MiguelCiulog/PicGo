package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var completedClueColors = lipgloss.AdaptiveColor{Light: "#A5D8FF", Dark: "#154162"}
var dueClueColors = lipgloss.AdaptiveColor{Light: "#1E1E1E", Dark: "#D3D3D3"}
var blankCellColors = lipgloss.AdaptiveColor{Light: "#FFFFFF", Dark: "#121212"}
var filledCellColors = lipgloss.AdaptiveColor{Light: "#1E1E1E", Dark: "#D3D3D3"}
var crossedCellColors = lipgloss.AdaptiveColor{Light: "#FFC9C9", Dark: "#5A2C2C"}
var selectedCellColors = lipgloss.AdaptiveColor{Light: "#6E757C", Dark: "#6E757C"}

var CompletedClue = lipgloss.NewStyle().
	Background(completedClueColors).
	Border(lipgloss.ThickBorder(), true)

var DueClue = lipgloss.NewStyle().Background(dueClueColors).Border(lipgloss.ThickBorder(), true)

var BlankCell = lipgloss.NewStyle().
	Foreground(blankCellColors).
	Background(blankCellColors).
	Width(2).
	Height(1).
	Margin(0).
	Border(lipgloss.HiddenBorder(), false)

var SelectedCell = lipgloss.NewStyle().
	Foreground(blankCellColors).
	Background(selectedCellColors).
	Width(2).
	Height(1).
	Margin(0).
	Blink(true).
	AlignVertical(lipgloss.Center).
	AlignHorizontal(lipgloss.Center).
	Border(lipgloss.HiddenBorder(), false)

var FilledCell = lipgloss.NewStyle().
	Background(filledCellColors).
	Border(lipgloss.ThickBorder(), true)

var CrossedCell = lipgloss.NewStyle().
	Background(crossedCellColors).
	Border(lipgloss.ThickBorder(), true)
