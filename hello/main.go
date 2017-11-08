package main

import (
	"github.com/nats-io/nats"
	"fmt"
	"runtime"
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

	conn.Subscribe("service.greeter.sayHello", func(subject, reply string, r GreeterRequest) {
		conn.Publish(reply, GreeterResponse{Greet: fmt.Sprintf("Hello %s", r.Name)})
	})

	runtime.Goexit()
}
