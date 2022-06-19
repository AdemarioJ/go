package server

import (
	"crud/database"
	"fmt"
	"strconv"

	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Add user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, error := ioutil.ReadAll(r.Body)

	if error != nil {
		w.Write([]byte("Failed to read body!"))
	}

	var user user

	if error = json.Unmarshal(bodyRequest, &user); error != nil {
		w.Write([]byte("Error parsing body!"))
		return
	}

	db, error := database.Connection()
	if error != nil {
		w.Write([]byte("Failed to connect to database"))
		return
	}
	defer db.Close()

	// PREPARE STATEMENT
	statement, error := db.Prepare("insert into users (name, email) values (?, ?)")
	if error != nil {
		w.Write([]byte("Failed to prepare statement"))
		return
	}
	defer statement.Close()

	insertion, error := statement.Exec(user.Name, user.Email)
	if error != nil {
		w.Write([]byte("Failed to insert user"))
		return
	}

	id, error := insertion.LastInsertId()
	if error != nil {
		w.Write([]byte("Failed to get last inserted id"))
		return
	}

	// STATUS CODES
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User inserted with success! ID: %d", id)))
}

// GetUsers return list of users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, error := database.Connection()
	if error != nil {
		w.Write([]byte("Failed to connect to database"))
		return
	}
	defer db.Close()

	lines, error := db.Query("select * from users")
	if error != nil {
		w.Write([]byte("Failed to query users"))
		return
	}

	defer lines.Close()

	var users []user

	for lines.Next() {
		var user user

		if error := lines.Scan(&user.ID, &user.Name, &user.Email); error != nil {
			w.Write([]byte("Error scanning user"))
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if error := json.NewEncoder(w).Encode(users); error != nil {
		w.Write([]byte("Error converting users to JSON"))
		return
	}
}

// GetUser return specific user
func ShowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, error := strconv.ParseInt(params["id"], 10, 32)
	if error != nil {
		w.Write([]byte("Failed to convert ID to int"))
		return
	}

	db, error := database.Connection()
	if error != nil {
		w.Write([]byte("Failed to connect to database"))
		return
	}
	defer db.Close()

	line, error := db.Query("select * from users where id = ?", ID)
	if error != nil {
		w.Write([]byte("Failed to query user"))
		return
	}

	defer line.Close()

	var user user
	if line.Next() {
		if error := line.Scan(&user.ID, &user.Name, &user.Email); error != nil {
			w.Write([]byte("Error scanning user"))
			return
		}
	}

	if error := json.NewEncoder(w).Encode(user); error != nil {
		w.Write([]byte("Error converting user to JSON"))
		return
	}
}
