package address

import "testing"

func TestAddressType(t *testing.T) {
	addressStr := "Avenida Paulista"
	expectedAddressType := "Avenida"
	addressType := AddressType(addressStr)

	if addressType != expectedAddressType {
		t.Error("Address type received is defferent from expected address type")
	}
}
