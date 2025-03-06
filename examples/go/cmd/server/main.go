package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/c-wire/fastrtb"
	"google.golang.org/grpc"
)

const port = 8080

// Server is the gRPC server implementation.
type Server struct {
	fastrtb.UnimplementedBidServiceServer
}

// GetBids recieves a bid requests and repsponse with the bids.
func (s *Server) GetBids(ctx context.Context, in *fastrtb.BidRequest) (*fastrtb.BidResponse, error) {
	// for now just echo, your bidder logic goes here.
	return &fastrtb.BidResponse{
		RequestId: in.RequestId,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		slog.Error("failed to listen", slog.Any("errorMsg", err))
	}

	grpcServer := grpc.NewServer()
	fastrtb.RegisterBidServiceServer(grpcServer, &Server{})

	slog.Info("grpc server started", slog.Any("port", port))
	if err := grpcServer.Serve(listener); err != nil {
		slog.Error("failed to serve", slog.Any("errorMsg", err))
	}
}
