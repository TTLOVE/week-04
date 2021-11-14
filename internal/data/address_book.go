package data

import (
	"context"
	"week-04/internal/biz"
)

type addressBookRepo struct {
	data *Data
}

// NewAddressBookRepo .
func NewAddressBookRepo(data *Data) biz.AddressBookRepo {
	return &addressBookRepo{
		data: data,
	}
}

func (r *addressBookRepo) CreateAddressBook(ctx context.Context, g *biz.AddressBook) error {
	return nil
}

func (r *addressBookRepo) UpdateAddressBook(ctx context.Context, g *biz.AddressBook) error {
	return nil
}
