/* main.go */

package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// Define an application struct to hold the application-wide dependencies for the web app.
type application struct {
	books *BookModel
}

func main() {
	db, err := sql.Open("sqlite3", "./data/database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	/* Initialize a new instance of our application struct. Establish a new BookModel struct
	   in our main() function and then inject it as a dependency via the application struct. 
	   This will allow us to make the BookModel object available to our handlers.*/
	app := &application{
		books: &BookModel{DB: db},
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", app.indexHandler)
	http.HandleFunc("/view", app.viewHandler)
	http.HandleFunc("/create", app.createHandler)
	http.HandleFunc("/edit", app.editHandler)
	http.HandleFunc("/delete", app.deleteHandler)
	http.HandleFunc("/about", app.aboutHandler)
	
	log.Println("Serving HTTP on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", logRequest(http.DefaultServeMux)))
}
