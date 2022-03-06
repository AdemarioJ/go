package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Printf("Saving the user %s\n", u.name)
}

func (u user) overage() bool {
	return u.age > 18
}

func (u *user) makeOlder() {
	u.age++
}

func main() {
	fmt.Println("Methods")

	user1 := user{"Joseph", 15}
	user1.save()
	fmt.Println(user1.overage())

	user2 := user{"Davi", 30}
	user2.save()

	fmt.Println(user2.overage())
	fmt.Println(user2.age)

	user2.makeOlder()
	fmt.Println(user2.age)
}
