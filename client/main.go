package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/nats-io/go-nats"
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

	for i := 0; i < 10; i++ {
		response := GreeterResponse{}
		ctx, _ := context.WithTimeout(context.Background(), time.Second)
		err = conn.RequestWithContext(ctx, "service.greeter.sayHello", &GreeterRequest{Name: strconv.Itoa(i)}, &response)
		if err != nil {
			panic(err)
		}

		fmt.Println(response.Greet)
	}
}
