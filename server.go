package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/lib/pq"
)

type User struct {
	Id        int    `json:"id"`
	Age       int    `json:"age"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func main() {

	fmt.Println("Rest API - Mux Routers and Postgres")
	handleRequests()

}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("users/all", allUsers)
	log.Fatal(http.ListenAndServe(":8080", myRouter))

}

func allUsers(w http.ResponseWriter, _ *http.Request) {

	db:= connect()

	rows := executeQuery(db, "select * from users")

	user := &User{}

	for rows.Next() {

		var id int
		var age int
		var firstName string
		var lastName string
		var email string

		err := rows.Scan(&id, &age, &firstName, &lastName, &email)

		checkErr(err)

		user = &User{
			Id:        id,
			Age:       age,
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
		}

	}

	json.NewEncoder(w).Encode(user)

}
func executeQuery(db *sql.DB, query string) (*sql.Rows) {

	rows, e := db.Query(query)
	checkErr(e)

	return rows

}

func connect() (*sql.DB) {

	db, err := sql.Open("postgres", "postgres://rqwybmmu:bwR6Z9xXSt1CC5LyJPDHKXXDKTGVEBo_@baasu.db.elephantsql.com:5432/rqwybmmu")
	checkErr(err)
	fmt.Println("Connected")
	defer db.Close()

	return db

}

func checkErr(err error) {

	if err != nil {

		panic(err)

	}

}
