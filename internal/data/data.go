package data

import (
	"fmt"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAddressBookRepo)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData() (*Data, func(), error) {
	cleanup := func() {
		fmt.Println("closing the data resources")
	}
	return &Data{}, cleanup, nil
}
