package main

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	desc "github.com/wanomir/simple-grpc/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedNoteV1Server // safety fallback for non-implemented handlers
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s) // tern off server-provided docs
	desc.RegisterNoteV1Server(s, &server{})

	log.Printf("grpc server listening on port %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Note id: %v\n", req.GetId())

	return &desc.GetResponse{
		Note: &desc.Note{
			Id: req.GetId(),
			Info: &desc.NoteInfo{
				Title:    gofakeit.BeerName(),
				Content:  gofakeit.IPv4Address(),
				Author:   gofakeit.Name(),
				IsPublic: gofakeit.Bool(),
			},
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}
