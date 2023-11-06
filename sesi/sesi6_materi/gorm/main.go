package main

import (
	"errors"
	"fmt"
	"gorm/database"
	"gorm/models"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()

	// createUser("fitri@mail.com")
	// getUserById(1)
	// updateUserById(1, "fiyu@mail.com")
	// createBook(1, "Laskar Pelangi", "Andrea Hirata", 20)
	// getUserWithBook()
	// deleteBookById(1)
}

func createUser(email string) {
	db := database.GetDB()
	if db == nil {
		fmt.Println("Error: Database connection is nil.")
		return
	}

	user := models.User{
		Email: email,
	}

	err := db.Create(&user).Error

	if err != nil {
		fmt.Println("Error creating user data: ", err)
		return
	}

	fmt.Println("New User Data", user)
}

func getUserById(id uint) {
	db := database.GetDB()

	user := models.User{}

	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("Error finding user:", err)
	}

	fmt.Printf("User data: %+v \n", user)
}

func updateUserById(id int, email string) {
	db := database.GetDB()

	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error

	if err != nil {
		fmt.Println("Error updating user data:", err)
		return
	}
	fmt.Printf("Update user's emaill: %+v \n", user.Email)
}

func createBook(userId uint, title string, author string, stock int) {
	db := database.GetDB()

	Book := models.Book{
		UserID: userId,
		Title:  title,
		Author: author,
		Stock:  stock,
	}

	err := db.Create(&Book).Error

	if err != nil {
		fmt.Println("Error creating book data:", err)
		return
	}

	fmt.Println("New Book Data", Book)
}

func getUserWithBook() {
	db := database.GetDB()

	users := models.User{}
	err := db.Preload("Books").Find(&users).Error

	if err != nil {
		fmt.Println("Error getting user data with books:", err.Error())
		return
	}

	fmt.Println("User Datas With Books")
	fmt.Printf("%+v", users)
}

func deleteBookById(id uint) {
	db := database.GetDB()

	book := models.Book{}

	err := db.Where("id = ?", id).Delete(&book).Error
	if err != nil {
		fmt.Println("Error deleting book:", err.Error())
		return
	}

	fmt.Printf("Book with id %d has been successfully deleted", id)
}
