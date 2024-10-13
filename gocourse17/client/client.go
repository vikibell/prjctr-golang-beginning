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

func sendReview(client pb.ReviewClient, reviewRequest *pb.ReviewRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	response, err := client.SendReview(ctx, reviewRequest)
	if err != nil {
		fmt.Printf("client.SendReview failed: %v\n", err)
	}
	fmt.Println(response)
}

func getHistory(client pb.ReviewClient, historyRequest *pb.HistoryRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	response, err := client.GetHistory(ctx, historyRequest)
	if err != nil {
		fmt.Printf("client.GetHistory failed: %v\n", err)
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

	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	driverFirstId := int32(1)
	driverSecondId := int32(2)

	for range 3 {
		sendReview(client,
			&pb.ReviewRequest{
				DriverId: driverFirstId,
				Review: &pb.ReviewData{
					CargoState:       rand.Int32N(5),
					ServiceQuality:   rand.Int32N(5),
					FulfillmentSpeed: rand.Int32N(5),
				},
			},
		)
	}

	for range 2 {
		sendReview(client,
			&pb.ReviewRequest{
				DriverId: driverSecondId,
				Review: &pb.ReviewData{
					CargoState:       rand.Int32N(5),
					ServiceQuality:   rand.Int32N(5),
					FulfillmentSpeed: rand.Int32N(5),
				},
			},
		)
	}

	getHistory(client, &pb.HistoryRequest{DriverId: driverFirstId})
	getHistory(client, &pb.HistoryRequest{DriverId: driverSecondId})
}
