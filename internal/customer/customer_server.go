package customer

import (
	"context"
	bookingpb "gRPC-Taxi-Reservation/gen/booking"
	customerpb "gRPC-Taxi-Reservation/gen/customer"
	"google.golang.org/grpc"
)

type Server struct {
	customerpb.UnimplementedCustomerServiceServer
	bookingClient bookingpb.BookingServiceClient
}

func NewCustomerServer(bookingConn *grpc.ClientConn) *Server {
	return &Server{
		bookingClient: bookingpb.NewBookingServiceClient(bookingConn),
	}
}

func (s *Server) ReserveTaxi(ctx context.Context, req *customerpb.RequestTaxi) (*customerpb.ResponseTaxi, error) {
	bookingRes, err := s.bookingClient.CreateBooking(ctx, &bookingpb.CreateBookingRequest{
		UserId:         req.UserId,
		PickupLocation: req.PickupLocation,
		DropLocation:   req.DropLocation,
	})
	if err != nil {
		return nil, err
	}

	return &customerpb.ResponseTaxi{
		BookingId: bookingRes.BookingId,
		Status:    bookingRes.Status,
	}, nil
}
