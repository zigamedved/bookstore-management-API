package models

import (
	"fmt"

	"github.com/zigamedved/bookstore-management-API/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title       string `gorm:"title" json:"title"`
	Author      string `gorm:"author" json:"author"`
	Publication string `gorm:"publication" json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
	fmt.Println("database connected")
}

func CreateBook(b *Book) *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBookById(Id int64) {
	db.Delete(&Book{}, Id)
}

func UpdateBookById(Id int64, book *Book) (*Book, *gorm.DB) {
	var updateBook Book
	db := db.Where("ID=?", Id).Updates(&book)
	return &updateBook, db
}
