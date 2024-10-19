package main

import (
	customerpb "gRPC-Taxi-Reservation/gen/customer"
	"gRPC-Taxi-Reservation/internal/customer"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to BookingService: %v", err)
	}
	defer conn.Close()

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	customerpb.RegisterCustomerServiceServer(grpcServer, customer.NewCustomerServer(conn))

	log.Println("CustomerService is running on port 50052...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
