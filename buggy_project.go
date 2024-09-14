package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB // why here, any reason?
var wg sync.WaitGroup

func main() {
	db, _ = sql.Open("postgres", "user=postgres dbname=test sslmode=disable") // handle error and will app work without DB?
	// also add defer db.Close()

	// making use of gin server would be better and efficient
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/create", createUser)

	log.Fatal(http.ListenAndServe(":8080", nil)) // hardcoded port
}

func getUsers(w http.ResponseWriter, r *http.Request) { // No structure separate file would be preferred, no docs
	nameCH := make(chan string, 50) // cap should be of specific limit
	go func(ch chan<- string) {     // what's the point? if block at line 45
		defer close(ch)

		// bad structure DB, routes and init part of single file
		rows, _ := db.Query("SELECT name FROM users") // 1. handle error 2. GetUser should have pagination and respective DB query as well. 3. use GORM
		//3. if number of user very high make sure to have max limit and take care of Partitioning if involved
		defer rows.Close()

		for rows.Next() { // 1. this can make use of goroutines but rows is not thread safe, not recommended unless absolutely required

			var name string
			rows.Scan(&name)
			ch <- name                         // this can be better than what was done previously
			fmt.Fprintf(w, "User: %s\n", name) // if absolutely  required use log
		}
	}(nameCH)

	// use name CH and log the data if required
	for name := range nameCH {
		log.Println(name)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	wg.Add(1)

	go func() { // no use, as request would in anycase wait for the execution to complete.  even though it is a long database op, still single op which blocks the request until db op completes
		defer wg.Done()

		time.Sleep(5 * time.Second) // Simulate a long database operation

		username := r.URL.Query().Get("name")
		//_, err := db.Exec("INSERT INTO users (name) VALUES ('" + username + "')") // Sql injection alert :). if unaware important do spend 5-10 mins reading about it

		stmt := `INSERT INTO users (username) VALUES ($1)`
		_, err := db.Exec(stmt, username) // handle same username case. make sure that Username is UNIQUE. if not add migration and add the Unique and NOT NULL constraints
		if err != nil {
			fmt.Fprintf(w, "Failed to create user: %v", err) // send error as part of w
			return
		}

		fmt.Fprintf(w, "User %s created successfully", username)
	}()

	wg.Wait()
}
