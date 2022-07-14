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

	if err := p.Start(); err != nil {
		log.Fatalln(err)
	}

}
