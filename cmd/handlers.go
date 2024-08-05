/* handlers.go */

package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

/* The signature of the home handler is defined as a method against *application.
   Because the home handler is a method against the application struct, it can
   access its fields. */
func (app *application) indexHandler(w http.ResponseWriter, r *http.Request) {
	books, err := app.books.GetAllBooks()
	if err != nil {
		log.Println(err)
	}

	indexTemplate, _ := template.ParseFiles("./templates/index.html")
	indexTemplate.Execute(w, books)
}

func (app *application) viewHandler(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("id")
	id, _ := strconv.Atoi(bookId)

	book, err := app.books.GetOneBook(id)
	if err != nil {
		log.Println(err)
	}

	viewTemplate := template.Must(template.ParseFiles("./templates/view.html"))
	viewTemplate.Execute(w, book)
}

func (app *application) createHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		createTemplate := template.Must(template.ParseFiles("./templates/create.html"))
		createTemplate.Execute(w, nil)
	case "POST":
		isbn := r.FormValue("isbn")
		title := r.FormValue("title")
		author := r.FormValue("author")
		excerpt := r.FormValue("excerpt")
		price := r.FormValue("price")

		type Message struct {
			Isbn    string
			Title   string
			Author  string
			Excerpt string
			Price   string
			Errors  map[string]string
		}
		msg := &Message{
			Isbn:    r.PostFormValue("isbn"),
			Title:   r.PostFormValue("title"),
			Author:  r.PostFormValue("author"),
			Excerpt: r.PostFormValue("excerpt"),
			Price:   r.PostFormValue("price"),
			Errors:  make(map[string]string),
		}

		if strings.TrimSpace(isbn) == "" {
			msg.Errors["Isbn"] = "Isbn required"
		}
		if strings.TrimSpace(title) == "" {
			msg.Errors["Title"] = "Title required"
		}
		if strings.TrimSpace(author) == "" {
			msg.Errors["Author"] = "Author required"
		}
		if strings.TrimSpace(excerpt) == "" {
			msg.Errors["Excerpt"] = "Excerpt required"
		}
		if strings.TrimSpace(price) == "" {
			msg.Errors["Price"] = "Price required"
		}

		if msg.Errors["Isbn"] != "" || msg.Errors["Title"] != "" || msg.Errors["Author"] != "" || msg.Errors["Excerpt"] != "" || msg.Errors["Price"] != ""{
			createTemplate, _ := template.ParseFiles("./templates/create.html")
			createTemplate.Execute(w, msg)
		} else {
			err := app.books.Insert(isbn, title, author, excerpt, price)
			if err != nil {
				log.Println(err)
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (app *application) editHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		bookId := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(bookId)

		book, err := app.books.GetOneBook(id)
		if err != nil {
			log.Println(err)
		}

		type Message struct {
			Id      int
			Isbn    string
			Title   string
			Author  string
			Excerpt string
			Price   string
			Errors  map[string]string
		}
		msg := &Message{
			Id:      book.Id,
			Isbn:    book.Isbn,
			Title:   book.Title,
			Author:  book.Author,
			Excerpt: book.Excerpt,
			Price:   book.Price,
			Errors:  make(map[string]string),
		}

		editTemplate := template.Must(template.ParseFiles("./templates/edit.html"))
		editTemplate.Execute(w, msg)
	case "POST":
		bookId := r.FormValue("id")
		id, _ := strconv.Atoi(bookId)

		isbn := r.FormValue("isbn")
		title := r.FormValue("title")
		author := r.FormValue("author")
		excerpt := r.FormValue("excerpt")
		price := r.FormValue("price")

		type Message struct {
			Id      int
			Isbn    string
			Title   string
			Author  string
			Excerpt string
			Price   string
			Errors  map[string]string
		}
		msg := &Message{
			Id:     id,
			Isbn:   isbn,
			Title:  title,
			Author: author,
			Excerpt: excerpt,
			Price:  price,
			Errors: make(map[string]string),
		}

		if strings.TrimSpace(isbn) == "" {
			msg.Errors["Isbn"] = "Isbn required"
		}
		if strings.TrimSpace(title) == "" {
			msg.Errors["Title"] = "Title required"
		}
		if strings.TrimSpace(author) == "" {
			msg.Errors["Author"] = "Author required"
		}
		if strings.TrimSpace(excerpt) == "" {
			msg.Errors["Excerpt"] = "Excerpt required"
		}
		if strings.TrimSpace(price) == "" {
			msg.Errors["Price"] = "Price required"
		}
		if msg.Errors["Isbn"] != "" || msg.Errors["Title"] != "" || msg.Errors["Author"] != "" || msg.Errors["Excerpt"] != "" || msg.Errors["Price"] != "" {
			editTemplate, _ := template.ParseFiles("./templates/edit.html")
			editTemplate.Execute(w, msg)
		} else {
			err := app.books.Update(id, isbn, title, author, excerpt, price)
			if err != nil {
				log.Println(err)
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (app *application) deleteHandler(w http.ResponseWriter, r *http.Request) {
	bookId := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(bookId)

	err := app.books.Delete(id)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) aboutHandler(w http.ResponseWriter, r *http.Request) {
	aboutTemplate, _ := template.ParseFiles("./templates/about.html")
	aboutTemplate.Execute(w, nil)
}
