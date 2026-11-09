package main

import (
	"fmt"

	"github.com/davi1985/go-fundaments/tests/address"
)

func main() {
	addressType := address.AddressType("Avenida Paulista")

	fmt.Print(addressType)
}
