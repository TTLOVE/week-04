package service

import (
	"context"
	"fmt"

	v1 "week-04/api/address_book/v1"
	"week-04/internal/biz"

	errors "github.com/go-kratos/kratos/v2/errors"
)

// AddressBookService is a greeter service.
type AddressBookService struct {
	v1.UnimplementedAddressBookServiceServer

	uc *biz.AddressBookUsecase
}

// NewAddressBookService new a greeter service.
func NewAddressBookService(uc *biz.AddressBookUsecase) *AddressBookService {
	return &AddressBookService{uc: uc}
}

// GetAddressBook implements address_book.AddressBookServer
func (s *AddressBookService) GetAddressBook(ctx context.Context, in *v1.GetAddressBookRequest) (*v1.AddressBookResponse, error) {
	if in.GetId() == 0 {
		return nil, errors.New(
			int(v1.ErrorReason_ADDRESS_BOOK_NOT_FOUND),
			fmt.Sprintf("addres book not found :%d", in.GetId()),
			"not found",
		)
	}

	return &v1.AddressBookResponse{
		Id: in.GetId(),
		PersonData: &v1.PersonBaseData{
			Name: "test name",
		},
	}, nil
}
