package tui_client

import tea "github.com/charmbracelet/bubbletea"

func loginView(m Model) string {
	return "PLS ENTER"
}

func loginUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.isLoggedIn = true
			m.readyToStream = true
		}

	}
	return m, nil
}
