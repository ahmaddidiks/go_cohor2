package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	db  *sql.DB
	err error
)

type Product struct {
	ID        uint
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type Variant struct {
	ID          uint
	VariantName string
	Quantity    uint
	ProductID   uint
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

func main() {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	// host := os.Getenv("HOST")
	dbname := os.Getenv("DB_NAME")
	// dbport := os.Getenv("DB_PORT")

	mysqlInfo := fmt.Sprintf("%s:%s@/%s?parseTime=true", user, password, dbname) // with time parser for *time.Time
	db, err = sql.Open("mysql", mysqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// fmt.Println("db connected")

	// create product with product name param
	// createProduct("Laptop")

	// update product name with id of product and new product name param
	// updateProduct(12, "New Laptop")

	// get product info by ID
	// getProductById(12)

	// create variant of product with product ID, variant name, and quantity param
	// createVariant(12, "Thinkpad X13", 3)

	//update variant by its varian ID, not product ID
	// updateVariantById(2, "Macbook Air M1", 2)

	// get product and all its variant
	// getProductWithVariant(12)

	// delete variant of product by variant ID not product ID
	deleteVariantById(4)
}

func createProduct(productName string) {
	var product = Product{}

	sqlStatement :=
		`INSERT INTO products(name)
		VALUE (?)
	`
	result, err := db.Exec(sqlStatement, productName)
	if err != nil {
		panic(err)
	}
	lastInserId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	// retrieve inserted row
	sqlRetrieve :=
		`SELECT * FROM products WHERE id = ?`

	err = db.QueryRow(sqlRetrieve, lastInserId).Scan(&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		panic(err)
	}

	fmt.Print(product)
}

func updateProduct(ID int, newName string) {
	sqlStatement := `UPDATE products SET name = ? WHERE id = ?;`

	result, err := db.Exec(sqlStatement, newName, ID)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("updated data amount: ", count)
}

func getProductById(ID int) {
	var product = Product{}

	// retrieve inserted row
	sqlRetrieve :=
		`SELECT * FROM products WHERE id = ?`

	err = db.QueryRow(sqlRetrieve, ID).Scan(&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		panic(err)
	}

	ValueOf := reflect.ValueOf(product)
	typeOf := reflect.TypeOf(product)

	for i := 0; i < ValueOf.NumField(); i++ {
		fmt.Printf("%+v : %+v\n", typeOf.Field(i).Name, ValueOf.Field(i))
	}
}

func createVariant(productID int, variantName string, quantity int) {
	var variant = Variant{}

	sqlStatement :=
		`INSERT INTO variants(variant_name, quantity, product_id)
		VALUE (?, ?, ?)
	`
	result, err := db.Exec(sqlStatement, variantName, quantity, productID)
	if err != nil {
		panic(err)
	}
	lastInserId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	// retrieve inserted row
	sqlRetrieve :=
		`SELECT * FROM variants WHERE id = ?`

	err = db.QueryRow(sqlRetrieve, lastInserId).Scan(&variant.ID, &variant.VariantName, &variant.Quantity, &variant.ProductID, &variant.CreatedAt, &variant.UpdatedAt)
	if err != nil {
		panic(err)
	}

	fmt.Print(variant)
}

func updateVariantById(ID int, newVariantName string, newQuantity int) {
	sqlStatement :=
		`UPDATE variants
	SET variant_name = ?, quantity = ?
	WHERE id = ?;`

	result, err := db.Exec(sqlStatement, newVariantName, newQuantity, ID)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("updated data amount: ", count)
}

func getProductWithVariant(productID int) {
	var result struct {
		product  Product
		variants []Variant
	}

	// get product
	// retrieve inserted row
	sqlRetrieve :=
		`SELECT * FROM products WHERE id = ?`

	err = db.QueryRow(sqlRetrieve, productID).Scan(&result.product.ID, &result.product.Name, &result.product.CreatedAt, &result.product.UpdatedAt)
	if err != nil {
		panic(err)
	}

	// get variants
	sqlStatement := `SELECT * from variants WHERE product_id = ?`
	rows, err := db.Query(sqlStatement, productID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var variant = Variant{}

		err = rows.Scan(&variant.ID, &variant.VariantName, &variant.ProductID, &variant.Quantity, &variant.CreatedAt, &variant.UpdatedAt)
		if err != nil {
			panic(err)
		}

		result.variants = append(result.variants, variant)
	}

	fmt.Println("Data of product ID : ", result.product.ID)

	ValueOf := reflect.ValueOf(result.product)
	typeOf := reflect.TypeOf(result.product)

	for i := 0; i < ValueOf.NumField(); i++ {
		fmt.Printf("%+v : %+v\n", typeOf.Field(i).Name, ValueOf.Field(i))
	}

	fmt.Printf("Varian of %s:\n", result.product.Name)

	for _, variant := range result.variants {
		fmt.Print("[")
		ValueOf := reflect.ValueOf(variant)
		typeOf := reflect.TypeOf(variant)

		for i := 0; i < ValueOf.NumField(); i++ {
			fmt.Printf("%+v : %+v\n", typeOf.Field(i).Name, ValueOf.Field(i))
		}
		fmt.Println("],")
	}
}

func deleteVariantById(ID int) {
	sqlStatement := `DELETE from variants WHERE id = ?;`
	result, err := db.Exec(sqlStatement, ID)
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("deleted data amount: ", count)
}
