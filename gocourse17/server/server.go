package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/vikibell/prjctr-golang-beginning/gocourse17/service"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/vikibell/prjctr-golang-beginning/gocourse17/grpcapi"
)

type server struct {
	pb.UnimplementedReviewServer

	history service.ReviewHistory
}

func (s *server) GetHistory(_ context.Context, request *pb.HistoryRequest) (*pb.HistoryResponse, error) {
	driverId := request.GetDriverId()
	if driverId <= 0 {
		return &pb.HistoryResponse{}, errors.New("invalid driver id")
	}

	history, exists := s.history.GerReviews(driverId)
	if !exists {
		return &pb.HistoryResponse{}, errors.New("there is no history for this driver")
	}

	var response []*pb.ReviewData
	for _, review := range history {
		response = append(response, &pb.ReviewData{
			CargoState:       review.CargoState,
			ServiceQuality:   review.ServiceQuality,
			FulfillmentSpeed: review.FulfillmentSpeed,
		})
	}

	return &pb.HistoryResponse{History: response}, nil
}

func (s *server) SendReview(_ context.Context, request *pb.ReviewRequest) (*pb.ReviewResponse, error) {
	driverId := request.GetDriverId()
	if driverId <= 0 {
		return &pb.ReviewResponse{Message: "Fail"}, errors.New("invalid driver id")
	}

	review := request.GetReview()
	//TODO check that each point from 0 to 5 and not empty
	newReview := service.NewReview(review.GetCargoState(), review.GetServiceQuality(), review.GetFulfillmentSpeed())
	s.history.AddReview(driverId, newReview)

	return &pb.ReviewResponse{Message: "Success"}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reviewHistory := service.NewReviewHistory()
	pb.RegisterReviewServer(s, &server{history: reviewHistory})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
