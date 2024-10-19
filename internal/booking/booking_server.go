package booking

import (
	"context"

	"sync"

	pb "gRPC-Taxi-Reservation/gen/booking"
)

type Server struct {
	pb.UnimplementedBookingServiceServer
	bookings map[uint64]string
	mu       sync.Mutex
}

func NewBookingServer() *Server {
	return &Server{
		bookings: make(map[uint64]string),
	}
}

func (s *Server) CreateBooking(ctx context.Context, req *pb.CreateBookingRequest) (*pb.CreateBookingResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	bookingID := uint64(len(s.bookings) + 1)
	s.bookings[bookingID] = "loading"

	return &pb.CreateBookingResponse{
		BookingId: bookingID,
		Status:    "loading",
	}, nil
}

func (s *Server) UpdateBooking(ctx context.Context, req *pb.UpdateBookingRequest) (*pb.UpdateBookingResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.bookings[req.BookingId] = req.Status
	return &pb.UpdateBookingResponse{
		BookingId: req.BookingId,
		Status:    req.Status,
	}, nil
}
