package main

import (
	"github.com/nats-io/nats"
	"fmt"
	"time"
	"context"
)

type GreeterRequest struct {
	Name string `json:"name"`
}

type GreeterResponse struct {
	Greet string `json:"greet"`
}

func main() {
	rawConn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}

	conn, err := nats.NewEncodedConn(rawConn, "json")
	if err != nil {
		panic(err)
	}

	response := GreeterResponse{}
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	err = conn.RequestWithContext(ctx, "service.greeter.sayHello", &GreeterRequest{Name: "World"}, &response)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Greet)
}
