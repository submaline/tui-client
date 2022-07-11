package tui_client

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	//greetv1 "tui-client/gen/greet/v1"
	"github.com/charmbracelet/bubbles/viewport"
)

// Model モデル
type Model struct {
	viewport  viewport.Model
	responses []ReceiveNotice
}

func InitializeModel() Model {
	vp := viewport.New(30, 10)

	return Model{
		viewport:  vp,
		responses: []ReceiveNotice{},
	}
}

// Init 初期化
func (m Model) Init() tea.Cmd {
	return nil
}

// Update Msgによって画面を更新するんだと思う
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var vpCmd tea.Cmd

	m.viewport, vpCmd = m.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	case ReceiveNotice:
		// 受け取ったらmodelの中に保存してあげる
		m.responses = append(m.responses, msg)
		// viewportに出力する内容
		s := ""
		// responsesは最終的に巨大な配列になる可能性があるので
		// 全てを描写するのではなく、最小を描写する
		// 配列の最終
		rangeE := len(m.responses) - 1
		// 配列の最初
		var rangeS int
		if len(m.responses) <= 10 {
			rangeS = 0
		} else {
			// 最初ではなく、最後から、vpで表示できる最大数取り出す。
			rangeS = len(m.responses) - 10
		}
		for _, resp := range m.responses[rangeS:rangeE] {
			s += fmt.Sprintf("(%v) %v\n", resp.Id, resp.Content)
		}
		m.viewport.SetContent(s)
		m.viewport.GotoBottom()
	}

	return m, tea.Batch(vpCmd)
}

// View 描写
func (m Model) View() string {

	var view string

	view += m.viewport.View()
	view += "\n\n"

	return view
}
