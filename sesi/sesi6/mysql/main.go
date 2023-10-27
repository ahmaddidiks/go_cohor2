package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	user    = "root"
	pasword = "Bismillaah123@"
	dbname  = "go-sql-cohort-2"
)

var (
	db  *sql.DB
	err error
)

type Book struct {
	ID     int
	Title  string
	Author string
	Stock  string
}

func main() {
	mysqlInfo := fmt.Sprintf("%s:%s@/%s", user, pasword, dbname)
	db, err = sql.Open("mysql", mysqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("yey db connected")

	// CreateBook("Homo Sapiens", "Yuval Noah Harari", 109)
	// GetBooks()
	UpdateBook()
	DeleteBook()
}

func CreateBook(title string, author string, stcok int) {
	var book = Book{}

	sqlStatement :=
		`INSERT INTO books (title, author, stock)
	VALUE (?, ?, ?)
	`
	result, err := db.Exec(sqlStatement, title, author, stcok)
	if err != nil {
		panic(err)
	}
	lastInserId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	// retrieve inserted row
	sqlRetrieve :=
		`SELECT * FROM books WHERE id = ?`

	err = db.QueryRow(sqlRetrieve, lastInserId).Scan(&book.ID, &book.Title, &book.Author, &book.Stock)
	if err != nil {
		panic(err)
	}

	fmt.Print(book)
}

func GetBooks() {
	var results = []Book{}

	sqlRetrieve := `SELECT * FROM books`
	rows, err := db.Query(sqlRetrieve)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var book = Book{}

		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Stock)
		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}

	fmt.Print(results)
}

func UpdateBook() {
	sqlStatement := `UPDATE books SET title = ?, author = ?, stock = ? WHERE id = ?;`
	result, err := db.Exec(sqlStatement, "Laskar Pelangi Update", "Anrea Hiara Update", 100, 1)
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("updated data amount: ", count)
}

func DeleteBook() {
	sqlStatement := `DELETE from books WHERE id = ?;`
	result, err := db.Exec(sqlStatement, 2)
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("deleted data amount: ", count)
}
