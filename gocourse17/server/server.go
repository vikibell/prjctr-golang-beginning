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
	pb.UnimplementedReviewerServer

	history service.ReviewHistory
}

func (s *server) GetHistory(_ context.Context, request *pb.GetHistoryRequest) (*pb.GetHistoryResponse, error) {
	driverID := int(request.GetDriverId())

	if driverID <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid driver id")
	}

	history := s.history.GetReviews(driverID)

	reviews := make([]*pb.Review, 0, len(history))
	for _, review := range history {
		reviews = append(reviews, &pb.Review{
			CargoState:       pb.Rating(review.CargoState),
			ServiceQuality:   pb.Rating(review.ServiceQuality),
			FulfillmentSpeed: pb.Rating(review.FulfillmentSpeed),
		})
	}

	return &pb.GetHistoryResponse{Reviews: reviews}, nil
}

func (s *server) SendReview(_ context.Context, request *pb.SendReviewRequest) (*pb.SendReviewResponse, error) {
	driverID := int(request.GetDriverId())
	if driverID <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid driver id")
	}

	review := request.GetReview()
	cs := service.Rating(review.GetCargoState())
	sq := service.Rating(review.GetServiceQuality())
	fs := service.Rating(review.GetFulfillmentSpeed())
	if !service.IsValid(cs) || !service.IsValid(sq) || !service.IsValid(fs) {
		return nil, status.Error(codes.InvalidArgument, "invalid review id")
	}

	newReview := service.NewReview(cs, sq, fs)
	s.history.AddReview(driverID, newReview)

	return &pb.SendReviewResponse{}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reviewHistory := service.NewReviewHistory()
	pb.RegisterReviewerServer(s, &server{history: reviewHistory})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
