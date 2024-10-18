package main

import (
	"context"
	pb "github.com/vikibell/prjctr-golang-beginning/gocourse17/grpcapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"log/slog"
	"math/rand/v2"
	"time"
)

func sendReview(client pb.ReviewerClient, reviewRequest *pb.SendReviewRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	response, err := client.SendReview(ctx, reviewRequest)
	if err != nil {
		slog.Error("client.SendReview failed", "error", err)
		return
	}
	slog.Info("client.SendReview", "response", response)
}

func getHistory(client pb.ReviewerClient, historyRequest *pb.GetHistoryRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.GetHistory(ctx, historyRequest)
	if err != nil {
		slog.Error("client.GetHistory failed", "error", err)
		return
	}
	slog.Info("client.GetHistory", "response", response)
}

func main() {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewReviewerClient(conn)

	driverFirstID := int32(1)
	driverSecondID := int32(2)

	for range 3 {
		sendReview(client,
			&pb.SendReviewRequest{
				DriverId: driverFirstID,
				Review: &pb.Review{
					CargoState:       pb.Rating(rand.IntN(5)),
					ServiceQuality:   pb.Rating(rand.Int32N(5)),
					FulfillmentSpeed: pb.Rating(rand.Int32N(5)),
				},
			},
		)
	}

	for range 2 {
		sendReview(client,
			&pb.SendReviewRequest{
				DriverId: driverSecondID,
				Review: &pb.Review{
					CargoState:       pb.Rating(rand.Int32N(5)),
					ServiceQuality:   pb.Rating(rand.Int32N(5)),
					FulfillmentSpeed: pb.Rating(rand.Int32N(5)),
				},
			},
		)
	}

	getHistory(client, &pb.GetHistoryRequest{DriverId: driverFirstID})
	getHistory(client, &pb.GetHistoryRequest{DriverId: driverSecondID})
}
