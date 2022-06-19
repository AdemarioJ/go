package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConection := "golang:golang@/devbook?charset=utf8&parseTime=true&loc=Local"

	db, err := sql.Open("mysql", stringConection)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection established!")

	lines, error := db.Query("select * from users")
	if error != nil {
		log.Fatal(error)
	}
	defer lines.Close()

	fmt.Println(lines)
}