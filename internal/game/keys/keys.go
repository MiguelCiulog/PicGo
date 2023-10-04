package keys

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding

	FillSquare   key.Binding
	CrossSquare  key.Binding
	DeleteSquare key.Binding

	Help    key.Binding
	Quit    key.Binding
	Restart key.Binding
	Number  key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.Left, k.Right},
		{k.FillSquare, k.CrossSquare},
		{k.DeleteSquare},
		{k.Help, k.Quit},
		{k.Number, k.Restart},
	}
}

var Keys = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "move left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "move right"),
	),

	FillSquare: key.NewBinding(
		key.WithKeys("f"),
		key.WithHelp("f", "fill square"),
	),
	CrossSquare: key.NewBinding(
		key.WithKeys("g"),
		key.WithHelp("g", "cross square"),
	),
	DeleteSquare: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete square guess"),
	),

	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc"),
		key.WithHelp("q/esc", "quit"),
	),
	Restart: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "restart"),
	),
}
