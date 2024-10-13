package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	pb "github.com/vikibell/prjctr-golang-beginning/gocourse17/grpcapi"
)

func main() {
	// TODO call of the server
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	pb.NewReviewClient(conn)

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
}
