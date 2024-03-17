package models

import (
	"github.com/ShreyanshKeshav33/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// we create a struct, structs are basically based on models and models gives strcutures that are store something in database
type Book struct {
	gorm.Model         //gorm.Model" is a struct that you embed in your model structs to automatically add fields like "ID", "CreatedAt", "UpdatedAt", and "DeletedAt", etc
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// initializing the database
func init() {
	config.Connect()
	db = config.GetDB()     //config folder file->app.go
	db.AutoMigrate(&Book{}) //Book struct
	//This is a method provided by GORM that automatically migrates the schema of the given struct to the corresponding table in the database
}

// Method to create a new book record
func (b *Book) CreateBook() *Book {
	db.NewRecord(b) // Declare a new record in the database
	db.Create(&b)   // Create a new book record in the database
	return b        // Return the created book object
}

// Function to get all books from the database
func GetAllBooks() []Book {
	var Books []Book // Declare a slice to hold the retrieved books
	db.Find(&Books)  // Find all books in the database and store them in the Books slice
	return Books     // Return the slice containing all books
}

// Function to get a book by its ID from the database
func GetBookById(Id int64) (*Book, *gorm.DB) { //*gorm.DB type in GORM represents a database handle. It's a pointer to a struct that contains methods for interacting with the database
	var getBook Book                          // Declare a variable to hold the retrieved book
	db := db.Where("ID=?", Id).Find(&getBook) // Find the book with the specified ID and store it in getBook
	return &getBook, db                       // Return a pointer to the retrieved book and the database query result
}

// Function to delete a book by its ID from the database
func DeleteBook(ID int64) Book {
	var book Book                     // Declare a variable to hold the deleted book
	db.Where("ID=?", ID).Delete(book) // Delete the book with the specified ID from the database
	return book                       // Return the deleted book
}
