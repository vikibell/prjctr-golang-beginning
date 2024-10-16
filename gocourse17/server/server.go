package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/vikibell/prjctr-golang-beginning/gocourse17/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"

	pb "github.com/vikibell/prjctr-golang-beginning/gocourse17/grpcapi"
)

type server struct {
	pb.UnimplementedReviewServer

	history service.ReviewHistory
}

func (s *server) GetHistory(_ context.Context, request *pb.GetHistoryRequest) (*pb.GetHistoryResponse, error) {
	driverId := int(request.GetDriverId())

	if driverId <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid driver id")
	}

	history, _ := s.history.GerReviews(driverId)

	var response []*pb.ReviewData
	for _, review := range history {
		response = append(response, &pb.ReviewData{
			CargoState:       int32(review.CargoState),
			ServiceQuality:   int32(review.ServiceQuality),
			FulfillmentSpeed: int32(review.FulfillmentSpeed),
		})
	}

	return &pb.GetHistoryResponse{Reviews: response}, nil
}

func (s *server) SendReview(_ context.Context, request *pb.SendReviewRequest) (*pb.SendReviewResponse, error) {
	driverId := int(request.GetDriverId())
	if driverId <= 0 {
		return &pb.SendReviewResponse{Message: "Fail"}, status.Errorf(codes.InvalidArgument, "invalid driver id")
	}

	review := request.GetReview()
	cs := int(review.GetCargoState())
	sq := int(review.GetServiceQuality())
	fs := int(review.GetFulfillmentSpeed())
	if !service.IsValidRating(cs) || !service.IsValidRating(sq) || !service.IsValidRating(fs) {
		return &pb.SendReviewResponse{Message: "Fail"}, status.Errorf(codes.InvalidArgument, "invalid review id")
	}

	newReview := service.NewReview(cs, sq, fs)
	s.history.AddReview(driverId, newReview)

	return &pb.SendReviewResponse{Message: "Success"}, nil
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
