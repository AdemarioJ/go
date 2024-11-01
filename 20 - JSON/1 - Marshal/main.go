package main

import (
	"bytes"
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
	dog := dog{Name: "Bobby", Age: 2, Breed: "Lab"}

	dogInJSON, error := json.Marshal(dog)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(string(dogInJSON))

	dog2 := map[string]string{
		"name":  "Nina",
		"breed": "Podle",
		"age":   "5",
	}

	dogInJSON2, error := json.Marshal(dog2)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(dogInJSON2)
	fmt.Println(bytes.NewBuffer(dogInJSON2))
}
