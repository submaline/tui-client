package main

import (
	tea "github.com/charmbracelet/bubbletea"
	tuiClient "github.com/submaline/tui-client"
	"github.com/submaline/tui-client/gen/greet/v1/greetv1connect"
	"log"
	"net/http"
)

func main() {

	// streamを受信するためのclientを用意
	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)

	ch := make(chan tuiClient.ReceiveNotice)

	// bubble teaの用意
	p := tea.NewProgram(tuiClient.InitializeModel(ch, &client))

	// streamを受信するプロセスを立ち上げる
	//go func() {
	//	stream, err := client.GreetStream(context.Background(),
	//		connect.NewRequest(&greetv1.GreetStreamRequest{Name: "1"}))
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	for stream.Receive() {
	//		msg := stream.Msg()
	//		// 受け取ったらbubble teaにcmd（外部操作）として通知してあげる。
	//		p.Send(tuiClient.ReceiveNotice{
	//			Id:      msg.Id,
	//			Content: msg.Text,
	//		})
	//	}
	//}()

	if err := p.Start(); err != nil {
		log.Fatalln(err)
	}

}
