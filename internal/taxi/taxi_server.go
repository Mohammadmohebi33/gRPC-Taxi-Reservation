package taxi

import (
	"context"
	bookingpb "gRPC-Taxi-Reservation/gen/booking"
	taxipb "gRPC-Taxi-Reservation/gen/taxi"
	"google.golang.org/grpc"
)

type Server struct {
	taxipb.UnimplementedTaxiServiceServer
	bookingClient bookingpb.BookingServiceClient
}

func NewTaxiServer(bookingConn *grpc.ClientConn) *Server {
	return &Server{
		bookingClient: bookingpb.NewBookingServiceClient(bookingConn),
	}
}

func (s *Server) AssignTaxi(ctx context.Context, req *taxipb.AssignTaxiRequest) (*taxipb.AssignTaxiResponse, error) {
	_, err := s.bookingClient.UpdateBooking(ctx, &bookingpb.UpdateBookingRequest{
		BookingId: req.BookingId,
		Status:    "taxi reserved",
	})
	if err != nil {
		return nil, err
	}

	return &taxipb.AssignTaxiResponse{
		TaxiId:     1,
		DriverName: "mohammad",
		CarModel:   "Peugeot 206",
		Status:     "taxi reserved",
	}, nil
}
