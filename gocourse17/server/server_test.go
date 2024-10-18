package main

import (
	"context"
	pb "github.com/vikibell/prjctr-golang-beginning/gocourse17/grpcapi"
	"github.com/vikibell/prjctr-golang-beginning/gocourse17/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"reflect"
	"testing"
	"time"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func createClient(t *testing.T) pb.ReviewerClient {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	reviewHistory := service.NewReviewHistory()
	reviewHistory.AddReview(1, service.NewReview(1, 2, 3))
	pb.RegisterReviewerServer(s, &server{history: reviewHistory})

	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()

	conn, err := grpc.Dial("bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to create client connection: %v", err)
	}

	t.Cleanup(func() { conn.Close() })

	return pb.NewReviewerClient(conn)
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGetHistory(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := createClient(t)

	t.Run("Success", func(t *testing.T) {
		req := &pb.GetHistoryRequest{DriverId: int32(1)}
		history, _ := client.GetHistory(ctx, req)
		got := history.GetReviews()
		want := []*pb.Review{{CargoState: pb.Rating(1), ServiceQuality: pb.Rating(2), FulfillmentSpeed: pb.Rating(3)}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetHistory(): got = %v, want %v", got, want)
		}
	})
	t.Run("Fail", func(t *testing.T) {
		req := &pb.GetHistoryRequest{DriverId: int32(-2)}
		got, err := client.GetHistory(ctx, req)
		if got != nil {
			t.Errorf("GetHistory(): got=%+v, want=%+v", got, nil)
		}
		if err == nil {
			t.Errorf("GetHistory(): should return error but got <nil>")
		}
		code, _ := status.FromError(err)
		want := codes.InvalidArgument

		if code.Code() != want {
			t.Errorf("GetHistory() unexpected error code: got=%+v, want=%+v", code, want)
		}
	})
}

func TestSendReview(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := createClient(t)
	req := &pb.SendReviewRequest{DriverId: 1, Review: &pb.Review{CargoState: pb.Rating(1), ServiceQuality: pb.Rating(2), FulfillmentSpeed: pb.Rating(3)}}
	_, err := client.SendReview(ctx, req)
	if err != nil {
		t.Fatalf("SendReview failed: %v", err)
	}
}
