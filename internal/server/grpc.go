package server

import (
	v1 "week-04/api/address_book/v1"
	"week-04/internal/service"
	"week-04/pkg/myGrpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(addressBook *service.AddressBookService) *myGrpc.Server {
	srv := myGrpc.NewServer()
	v1.RegisterAddressBookServiceServer(srv, addressBook)
	return srv
}
