// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/kataras/iris/v12"
	"week-04/internal/biz"
	"week-04/internal/data"
	"week-04/internal/server"
	"week-04/internal/service"
)

// Injectors from wire.go:

// initApp init application.
func initApp() (*iris.Application, func(), error) {
	dataData, cleanup, err := data.NewData()
	if err != nil {
		return nil, nil, err
	}
	addressBookRepo := data.NewAddressBookRepo(dataData)
	addressBookUsecase := biz.NewAddressBookUsecase(addressBookRepo)
	addressBookService := service.NewAddressBookService(addressBookUsecase)
	myGrpcServer := server.NewGRPCServer(addressBookService)
	application := newApp(myGrpcServer)
	return application, func() {
		cleanup()
	}, nil
}
