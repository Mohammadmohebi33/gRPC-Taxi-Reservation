package main

import (
	pb "gRPC-Taxi-Reservation/gen/booking"
	"gRPC-Taxi-Reservation/internal/booking"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBookingServiceServer(grpcServer, booking.NewBookingServer())

	log.Println("BookingService is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
