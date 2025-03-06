package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/c-wire/fastrtb"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = 8080

func main() {
	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := fastrtb.NewBidServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id, err := uuid.NewV7()
	if err != nil {
		slog.Error("Failed to generate uuid", slog.Any("errorMsg", err))
		return
	}

	res, err := client.GetBids(ctx, &fastrtb.BidRequest{
		RequestId: id.String(),
	})
	if err != nil {
		slog.Error("Failed to get bids", slog.Any("errorMsg", err))
		return
	}

	slog.Info("Got bids", slog.Any("response", res))
}
