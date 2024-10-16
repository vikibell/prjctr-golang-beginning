package main

import (
	"context"
	"fmt"
	pb "github.com/vikibell/prjctr-golang-beginning/gocourse17/grpcapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand/v2"
	"time"
)

func sendReview(client pb.ReviewClient, reviewRequest *pb.SendReviewRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	response, err := client.SendReview(ctx, reviewRequest)
	if err != nil {
		fmt.Printf("client.SendReview failed: %v\n", err)
		return
	}
	fmt.Println(response)
}

func getHistory(client pb.ReviewClient, historyRequest *pb.GetHistoryRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.GetHistory(ctx, historyRequest)
	if err != nil {
		fmt.Printf("client.GetHistory failed: %v\n", err)
		return
	}
	fmt.Println(response)
}

func main() {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewReviewClient(conn)

	driverFirstID := int32(1)
	driverSecondID := int32(2)

	for range 3 {
		sendReview(client,
			&pb.SendReviewRequest{
				DriverId: driverFirstID,
				Review: &pb.ReviewData{
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
				Review: &pb.ReviewData{
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
