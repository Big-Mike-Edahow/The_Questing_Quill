/* models.go */

package main

import (
	"database/sql"
	"log"
)

type Book struct {
	Id      int
	Isbn    string
	Title   string
	Author  string
	Excerpt string
	Price   string
}

type BookModel struct {
	DB *sql.DB
}

func (m *BookModel) GetOneBook(id int) (Book, error) {
	// QueryRow fetches a single row instead of multiple rows.
	row := m.DB.QueryRow("SELECT * FROM books WHERE id = ?", id)

	var book Book
	err := row.Scan(&book.Id, &book.Isbn, &book.Title, &book.Author, &book.Excerpt, &book.Price)
	if err != nil {
		log.Println(err)
	}
	return book, err
}

func (m *BookModel) GetAllBooks() ([]Book, error) {
	// Query returns all matching rows as a Rows struct your code can loop over.
	rows, err := m.DB.Query("SELECT * FROM books")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.Id, &book.Isbn, &book.Title, &book.Author, &book.Excerpt, &book.Price)
		if err != nil {
			log.Println(err)
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
	}

	return books, err
}

func (m *BookModel) Insert(isbn string, title string, author string, excerpt string, price string) error {
	stmt := "INSERT INTO books(isbn, title, author, excerpt, price) VALUES(?, ?, ?, ?, ?)"
	_, err := m.DB.Exec(stmt, isbn, title, author, excerpt, price)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (m *BookModel) Update(id int, isbn string, title string, author string, excerpt string, price string) error {
	stmt := "UPDATE books SET isbn=?, title=?, author=?, excerpt=?, price=? WHERE id=?"
	_, err := m.DB.Exec(stmt, isbn, title, author, excerpt, price, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (m *BookModel) Delete(id int) error {
	stmt := "DELETE FROM books WHERE id = ?"
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		log.Println(err)
	}
	return err
}
