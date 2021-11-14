package main

import (
	"context"
	"log"
	v1 "week-04/api/address_book/v1"

	"github.com/kataras/iris/v12"
	"google.golang.org/grpc"
)

var client v1.AddressBookServiceClient

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug") //debug
	app.Handle("GET", "/address_book", getAddressBook)
	app.Run(iris.Addr("127.0.0.1:8000"))
}

func getAddressBook(ctx iris.Context) {
	params := v1.GetAddressBookRequest{}
	params.Id = 12
	res, err := client.GetAddressBook(context.Background(), &params)
	if err != nil {
		log.Fatalf("client.GetAddressBook err: %v", err)
	}
	ctx.JSON(res)
}

func init() {
	connect, err := grpc.Dial("127.0.0.1:9527", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	client = v1.NewAddressBookServiceClient(connect)

}
