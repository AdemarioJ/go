package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"name":     "Davi",
		"lastname": "Silva",
	}

	fmt.Println(user)

	user2 := map[string]map[string]string{
		"name": {
			"first": "Davi",
			"last":  "Silva",
		},
		"address": {
			"street": "Rua dos Bobos",
			"number": "123",
		},
	}

	fmt.Println(user2)
	delete(user2, "address")
	fmt.Println(user2)

	user2["address"] = map[string]string{
		"street": "Rua dos Bobos",
		"number": "123",
	}

	fmt.Println(user2)
}
