package main

import (
	"fmt"
	"modelo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquvo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("ademariocc@gmail.com")
	fmt.Println(erro)
}
