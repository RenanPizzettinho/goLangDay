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

	Id int `json:"id"`
	Age int `json:"age"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`

}

func main() {
 
	fmt.Println("Rest API - Mux Routers and Postgres")
	handleRequests()

}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/sql", returnSql)
	log.Fatal(http.ListenAndServe(":8080", myRouter))

}

func returnSql(w http.ResponseWriter, r *http.Request){

	db, err := sql.Open("postgres", "postgres://rqwybmmu:bwR6Z9xXSt1CC5LyJPDHKXXDKTGVEBo_@baasu.db.elephantsql.com:5432/rqwybmmu")
	
	checkErr(err)

	fmt.Println("Connected")

	defer db.Close()

	rows, err := db.Query("select * from users")

	checkErr(err)

	user := &User{}

	for (rows.Next()) {
		
		var id int
		var age int
		var first_name string
		var last_name string
		var email string

		err = rows.Scan(&id, &age, &first_name, &last_name, &email)
		
		checkErr(err)
		
		user = &User{
			Id: id,
			Age: age,
			FirstName: first_name,
			LastName: last_name,
			Email: email,
		}

	}

	json.NewEncoder(w).Encode(user)

}

func checkErr(err error) {

	if (err != nil) {

		panic(err)

	}

}
