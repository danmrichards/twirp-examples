package main

import (
	"context"
	"fmt"
	"net/http"

	pb "github.com/danmrichards/twirp-examples/helloworld/rpc/helloworld"
)

func main() {
	client := pb.NewHelloWorldProtobufClient("http://localhost:8080", &http.Client{})

	resp, err := client.Hello(context.Background(), &pb.HelloReq{Subject: "Emma"})
	if err == nil {
		fmt.Println(resp.Text)
	}
}
