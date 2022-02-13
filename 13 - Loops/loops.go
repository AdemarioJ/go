package main

import (
	"fmt"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementing i:", i)
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println("Final value of i:", i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementing j:", j)
	// 	time.Sleep(time.Second)
	// }

	names := [3]string{"John", "Paul", "George"}

	for key, value := range names {
		fmt.Println("Index:", key, "Value:", value)
	}

	for key, letter := range "WORD" {
		fmt.Println("Index:", key, "Value:", string(letter))
	}

	users := map[string]string{
		"user1": "John",
		"user2": "Paul",
		"user3": "George",
	}

	for key, value := range users {
		fmt.Println("Key:", key, "Value:", value)
	}
}
