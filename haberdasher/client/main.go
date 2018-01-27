package main

import (
	"context"
	"fmt"
	"net/http"

	pb "github.com/danmrichards/twirp-examples/haberdasher/rpc/haberdasher"
	"github.com/twitchtv/twirp"
)

func main() {
	client := pb.NewHaberdasherProtobufClient("http://localhost:8080", &http.Client{})

	hat, err := client.MakeHat(context.Background(), &pb.Size{Inches: 5})
	if err != nil {
		if twerr, ok := err.(twirp.Error); ok {
			switch twerr.Code() {
			case twirp.InvalidArgument:
				fmt.Println("Oh no " + twerr.Msg())
			default:
				fmt.Println(twerr.Error())
			}
		}
		return
	}

	fmt.Printf("Here is your new %d inch %s %s", hat.Inches, hat.Color, hat.Name)
}
