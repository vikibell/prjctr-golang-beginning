package main

import (
	"context"
	"github.com/google/go-cmp/cmp"
	pb "github.com/vikibell/prjctr-golang-beginning/gocourse17/grpcapi"
	"github.com/vikibell/prjctr-golang-beginning/gocourse17/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"reflect"
	"testing"
	"time"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	reviewHistory := service.NewReviewHistory()
	reviewHistory.AddReview(int32(1), service.NewReview(int32(1), int32(2), int32(3)))
	pb.RegisterReviewServer(s, &server{history: reviewHistory})

	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGetHistory(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.Dial("bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to create client connection: %v", err)
	}
	defer conn.Close()

	client := pb.NewReviewClient(conn)

	req := &pb.GetHistoryRequest{DriverId: int32(1)}
	history, err := client.GetHistory(ctx, req)
	if err != nil {
		t.Fatalf("GetHistory failed: %v", err)
	}

	got := history.GetReviews()
	want := []*pb.ReviewData{{CargoState: int32(1), ServiceQuality: int32(2), FulfillmentSpeed: int32(3)}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetHistory(): got = %v, want %v", got, want)
	}
}

func TestSendReview(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.Dial("bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to create client connection: %v", err)
	}
	defer conn.Close()

	client := pb.NewReviewClient(conn)

	req := &pb.SendReviewRequest{DriverId: int32(1), Review: &pb.ReviewData{CargoState: int32(1), ServiceQuality: int32(2), FulfillmentSpeed: int32(3)}}
	resp, err := client.SendReview(ctx, req)
	if err != nil {
		t.Fatalf("SendReview failed: %v", err)
	}

	got := resp.GetMessage()
	want := "Success"
	if !cmp.Equal(got, want) {
		t.Errorf("SendReview(): got = %v, want %v", got, want)
	}
}
