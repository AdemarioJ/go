package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	city  string
	state string
}

func main() {
	fmt.Println("File structs")

	var u user
	u.name = "David"
	u.age = 30
	fmt.Println(u)

	addressExemple := address{"New York", "NY"}

	u2 := user{"David", 30, addressExemple}
	fmt.Println(u2)

	u3 := user{name: "David"}
	fmt.Println(u3)

}
