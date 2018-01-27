package main

import (
	"context"
	"math/rand"
	"net/http"

	pb "github.com/danmrichards/twirp-examples/haberdasher/rpc/haberdasher"
	"github.com/twitchtv/twirp"
)

// Server implements the Haberdasher service
type Server struct{}

func (s *Server) MakeHat(ctx context.Context, size *pb.Size) (*pb.Hat, error) {
	if size.Inches <= 0 {
		return nil, twirp.InvalidArgumentError("inches", "I can't make a hat that small!")
	}
	return &pb.Hat{
		Inches: size.Inches,
		Color:  []string{"white", "black", "brown", "red", "blue"}[rand.Intn(4)],
		Name:   []string{"bowler", "baseball cap", "top hat", "derby"}[rand.Intn(3)],
	}, nil
}

func main() {
	twirpHandler := pb.NewHaberdasherServer(&Server{}, nil)
	http.ListenAndServe(":8080", twirpHandler)
}
