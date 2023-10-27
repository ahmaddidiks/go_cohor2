package main

import (
	"fmt"
	"sesi_6/gorm/database"
	"sesi_6/gorm/models"
)

func main() {
	database.StartDB()

	// createUser("didik@gmail.com")
	// getUserByID(1)
	createBook(1, "Buku test", "Penulis 2", 20)
	// getUserWithBook()
	// updateUSerByID(3, "ahmad@gmail.com")
	deleteBookByID(1)
}

func createUser(email string) {
	db := database.GetDB()
	if db == nil {
		panic("Error: Database connection is nil")
	}

	user := models.User{
		Email: email,
	}

	err := db.Create(&user).Error
	if err != nil {
		panic(err)
	}

	fmt.Println("New User Data:", user)
}

func getUserByID(id uint) {
	db := database.GetDB()
	if db == nil {
		panic("Error: Database connection is nil")
	}

	user := models.User{}

	err := db.First(&user, "id = ?", id).Error
	if err != nil {
		panic(err)
	}

	fmt.Printf("User data ID %d is %+v\n", id, user)
}

func createBook(userID uint, title string, author string, stock int) {
	db := database.GetDB()
	if db == nil {
		panic("Error: Database connection is nil")
	}

	book := models.Book{
		UserID: userID,
		Title:  title,
		Author: author,
		Stock:  stock,
	}

	err := db.Create(&book).Error
	if err != nil {
		panic(err)
	}

	fmt.Println("New Book Data:", book)
}

func getUserWithBook() {
	db := database.GetDB()
	if db == nil {
		panic("Error: Database connection is nil")
	}

	users := models.User{}
	err := db.Preload("Books").Find(&users).Error
	if err != nil {
		panic(err)
	}

	fmt.Println("User data with books")
	fmt.Println(users)
}

func deleteBookByID(id uint) {
	// db := database.GetDB()
	// book := models.Book{}

	// err := db.Where("id = ?", id).Delete(&book).Error
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Book with id %d has been deleted \n", id)

	db := database.GetDB()

	book := models.Book{}

	err := db.Where("id = ?", id).Delete(&book).Error
	if err != nil {
		fmt.Println("Error deleting book:", err.Error())
		return
	}

	fmt.Printf("Book with id %d has been successfully deleted", id)
}

func updateUSerByID(id int, email string) {
	db := database.GetDB()
	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Update("email", email).Error
	if err != nil {
		panic(err)
	}
	fmt.Printf("Update user's email: %+v \n", user.Email)

}
