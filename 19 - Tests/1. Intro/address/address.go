package address

import "strings"

// AddressType
func AddressType(address string) string {
	validAddressType := []string{"rua", "avenida", "estrada", "rodovia"}
	firstWordFromAddress, _, _ := strings.Cut(strings.ToLower(address), " ")

	hasValideAddressType := false

	for _, value := range validAddressType {
		if value == firstWordFromAddress {
			hasValideAddressType = true
		}
	}

	if hasValideAddressType {
		return firstWordFromAddress
	}

	return "Invalid address type"
}
