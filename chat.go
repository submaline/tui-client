package tui_client

import (
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	tea "github.com/charmbracelet/bubbletea"
	greetv1 "github.com/submaline/tui-client/gen/greet/v1"
	"log"
)

func chatView(m Model) string {
	var view string
	view += "[RESP]\n"
	view += m.viewport.View()
	return view
}

func chatUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	if m.readyToStream {
		m.readyToStream = false
		go streaming(m)
	}

	var vpCmd tea.Cmd

	m.viewport, vpCmd = m.viewport.Update(msg)
	switch msg := msg.(type) {
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
		return m, tea.Batch(vpCmd, ReceiveNotifier(m.ReceiveNoticeCh))
	}

	return m, tea.Batch(vpCmd)

}

func streaming(m Model) {
	stream, err := (*m.greetV1Client).GreetStream(context.Background(),
		connect.NewRequest(&greetv1.GreetStreamRequest{Name: "1"}))
	if err != nil {
		log.Fatalln(err)
	}
	for stream.Receive() {
		msg := stream.Msg()
		m.ReceiveNoticeCh <- ReceiveNotice{
			Id:      msg.Id,
			Content: msg.Text,
		}
	}
}

func ReceiveNotifier(ch chan ReceiveNotice) tea.Cmd {
	return func() tea.Msg {
		return <-ch
	}
}
