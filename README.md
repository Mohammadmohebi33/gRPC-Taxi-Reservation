how compile proto files :

protoc --go_out=gen/customer --go-grpc_out=gen/customer proto/customer/customer.proto


protoc --go_out=gen/taxi --go-grpc_out=gen/taxi proto/taxi/taxi.proto



protoc --go_out=gen/booking --go-grpc_out=gen/booking proto/booking/booking.proto


how run : 

go run cmd/booking/main.go


go run cmd/customer/main.go


go run cmd/taxi/main.go
