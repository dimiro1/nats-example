package main

import (
	"fmt"
	"runtime"

	"github.com/nats-io/go-nats"
)

type GreeterRequest struct {
	Name string `json:"name"`
}

type GreeterResponse struct {
	Greet string `json:"greet"`
}

func main() {
	rawConn, err := nats.Connect(nats.DefaultURL, nats.Name("service.greeter"))
	if err != nil {
		panic(err)
	}

	conn, err := nats.NewEncodedConn(rawConn, "json")
	if err != nil {
		panic(err)
	}

	conn.QueueSubscribe("service.greeter.sayHello", "greeter", func(_, reply string, r GreeterRequest) {
		fmt.Printf("Received %s\n", r.Name)
		conn.Publish(reply, GreeterResponse{Greet: fmt.Sprintf("Hello %s", r.Name)})
	})

	runtime.Goexit()
}
