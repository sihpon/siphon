package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"
	greetv1 "github.com/siphon/siphon/gen/siphon/v1"
	"github.com/siphon/siphon/gen/siphon/v1/greetv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type GreetServer struct{}

func (s *GreetServer) Health(
	context.Context,
	*connect.Request[greetv1.HealthRequest],
) (*connect.Response[greetv1.HealthResponse], error) {
	return connect.NewResponse(&greetv1.HealthResponse{
		Status: "OK",
	}), nil
}

var _ greetv1connect.SiphonServiceHandler = (*GreetServer)(nil)

func (s *GreetServer) Echo(
	ctx context.Context,
	req *connect.Request[greetv1.EchoRequest],
) (*connect.Response[greetv1.EchoResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&greetv1.EchoResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

func main() {
	greeter := &GreetServer{}
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewSiphonServiceHandler(greeter)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
