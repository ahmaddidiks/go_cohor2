package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = ""
	password = ""
	dbname   = "db-go-sql"
)

var (
	db  *sql.DB
	err error
)

type Book struct {
	ID     int
	Title  string
	Author string
	Stock  int
}

func main() {
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", config)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	//CreateBook()
	//GetBooks()
	//UpdateBook()
	//DeleteBook()
}

func CreateBook() {
	var book = Book{}

	sqlStatement := `
	INSERT INTO books (title, author, stock)
	VALUES ($1, $2, $3)
	Returning *
	`

	err = db.QueryRow(sqlStatement, "Judul Buku", "Pengarang Buku", 30).Scan(&book.ID, &book.Title, &book.Author, &book.Stock)
	if err != nil {
		panic(err)
	}

	fmt.Printf("New Book Data: %+v\n", book)
}

func GetBooks() {
	var results = []Book{}

	sqlStatement := `SELECT * from books`

	rows, err := db.Query(sqlStatement)
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

	fmt.Println("Book datas:", results)
}

func DeleteBook() {
	sqlStatement := `DELETE from books WHERE id = $1;`

	res, err := db.Exec(sqlStatement, 1)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted data amount:", count)
}

func UpdateBook() {
	sqlStatement := `UPDATE books SET title = $2, author = $3, stock = $4 WHERE id = $1;`
	res, err := db.Exec(sqlStatement, 1, "Laskar Pelangi", "Andrea Hirata", 20)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated data amount: ", count)
}
