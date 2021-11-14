package biz

import (
	"context"
)

type AddressBook struct {
	Hello string
}

type AddressBookRepo interface {
	CreateAddressBook(context.Context, *AddressBook) error
	UpdateAddressBook(context.Context, *AddressBook) error
}

type AddressBookUsecase struct {
	repo AddressBookRepo
}

func NewAddressBookUsecase(repo AddressBookRepo) *AddressBookUsecase {
	return &AddressBookUsecase{repo: repo}
}

func (uc *AddressBookUsecase) Create(ctx context.Context, g *AddressBook) error {
	return uc.repo.CreateAddressBook(ctx, g)
}

func (uc *AddressBookUsecase) Update(ctx context.Context, g *AddressBook) error {
	return uc.repo.UpdateAddressBook(ctx, g)
}
