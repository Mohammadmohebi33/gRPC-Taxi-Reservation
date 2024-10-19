package main

import (
	taxipb "gRPC-Taxi-Reservation/gen/taxi"
	"gRPC-Taxi-Reservation/internal/taxi"
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

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	taxipb.RegisterTaxiServiceServer(grpcServer, taxi.NewTaxiServer(conn))

	log.Println("TaxiService is running on port 50053...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
