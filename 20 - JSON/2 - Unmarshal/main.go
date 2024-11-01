package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   uint   `json:"age"`
}

func main() {
	dogInJson := `{"name":"Rex","breed":"Dalmata", "age": 3}`

	var c dog

	if erro := json.Unmarshal([]byte(dogInJson), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	dog2InJson := `{"name":"Toby","breed":"Poodle"}`

	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(dog2InJson), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)
}
