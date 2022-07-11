package tui_client

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/submaline/tui-client/gen/greet/v1/greetv1connect"

	//greetv1 "tui-client/gen/greet/v1"
	"github.com/charmbracelet/bubbles/viewport"
)

// Model モデル
type Model struct {
	viewport  viewport.Model
	responses []ReceiveNotice

	isLoggedIn      bool
	readyToStream   bool
	greetV1Client   *greetv1connect.GreetServiceClient
	ReceiveNoticeCh chan ReceiveNotice
}

func InitializeModel(ch chan ReceiveNotice, cl *greetv1connect.GreetServiceClient) Model {
	vp := viewport.New(30, 10)

	return Model{
		viewport:        vp,
		responses:       []ReceiveNotice{},
		isLoggedIn:      false,
		readyToStream:   false,
		greetV1Client:   cl,
		ReceiveNoticeCh: ch,
	}
}

// Init 初期化
func (m Model) Init() tea.Cmd {
	return ReceiveNotifier(m.ReceiveNoticeCh)
}

// Update Msgによって画面を更新するんだと思う
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// 常に同じ動きをするもの
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	if m.isLoggedIn {
		return chatUpdate(msg, m)
	} else {
		return loginUpdate(msg, m)
	}

}

// View 描写
func (m Model) View() string {
	var s string
	if m.isLoggedIn {
		s = chatView(m)
	} else {
		s = loginView(m)
	}
	return s + "\n\n"
}
