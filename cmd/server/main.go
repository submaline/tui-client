package main

import (
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	greetv1 "github.com/submaline/tui-client/gen/greet/v1"
	"github.com/submaline/tui-client/gen/greet/v1/greetv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

type GreetServer struct {
}

func (s *GreetServer) Greet(_ context.Context,
	_ *connect.Request[greetv1.GreetRequest]) (*connect.Response[greetv1.GreetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, fmt.Errorf("a"))
}

func (s *GreetServer) GreetStream(ctx context.Context,
	_ *connect.Request[greetv1.GreetStreamRequest],
	stream *connect.ServerStream[greetv1.GreetStreamResponse]) error {
	c := 0

Loop:
	for {
		select {
		case <-ctx.Done():
			if ctx.Err() != nil {
				log.Println(ctx.Err())
			}
			break Loop
		default:
			err := stream.Send(&greetv1.GreetStreamResponse{Id: int64(c), Text: fmt.Sprintf("data: %v", c)})
			if err != nil {
				log.Println(err)
			}
			c++
			//time.Sleep(time.Millisecond * time.Duration(rand.Int63n(900)+100))
			log.Println(c)
		}
	}
	return nil
}

func main() {
	greeter := &GreetServer{}
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(greeter)
	mux.Handle(path, handler)
	if err := http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatalln(err)
	}
}
