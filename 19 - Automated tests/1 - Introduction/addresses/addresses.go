package addresses

import (
	"strings"
)

func TypeOfAddress(address string) string {
	validTypes := []string{"rua", "avenida", "estrada", "rodovia"}

	addressInLowerCase := strings.ToLower(address)
	firstWordAddress := strings.Split(addressInLowerCase, " ")[0]

	addressHasValidType := false
	for _, validType := range validTypes {
		if validType == firstWordAddress {
			addressHasValidType = true
		}
	}

	if addressHasValidType {
		return strings.Title(firstWordAddress)
	}

	return "Invalid type"
}
