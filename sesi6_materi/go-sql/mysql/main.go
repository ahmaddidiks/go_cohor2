package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "root"
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
	mysqlInfo := fmt.Sprintf("%s:%s@/%s", user, password, dbname)

	db, err = sql.Open("mysql", mysqlInfo)
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
	GetBooks()
	//UpdateBook()
	//DeleteBook()
}

func CreateBook() {
	var book = Book{}

	sqlStatement := `
	INSERT INTO books (title, author, stock)
	VALUES (?, ?, ?)
	`

	result, err := db.Exec(sqlStatement, "Judul Buku Baru", "Pengarang Buku Baru", 30)
	if err != nil {
		panic(err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	// Retrieve the inserted row using the lastInsertID
	sqlRetrieve := `
    SELECT * FROM books WHERE id = ?
	`

	err = db.QueryRow(sqlRetrieve, lastInsertID).Scan(&book.ID, &book.Title, &book.Author, &book.Stock)
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
	sqlStatement := `DELETE from books WHERE id = ?;`

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
	sqlStatement := `UPDATE books SET title = ?, author = ?, stock = ? WHERE id = ?;`
	res, err := db.Exec(sqlStatement, "Laskar Pelangi", "Andrea Hirata", 20, 2)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated data amount: ", count)
}
