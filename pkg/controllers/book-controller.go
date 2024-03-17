package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ShreyanshKeshav33/go-bookstore/pkg/models"
	"github.com/ShreyanshKeshav33/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book // Declare a variable NewBook of type models.Book, which refers to the Book struct in the models file

func GetBook(w http.ResponseWriter, r *http.Request) { // Define a function named GetBook that takes a ResponseWriter and a Request as parameters
	newBooks := models.GetAllBooks()                   // Call the GetAllBooks function from the models package and store the result in the newBooks variable
	res, _ := json.Marshal(newBooks)                   // Marshal the newBooks variable into JSON format and store the result in the res variable
	w.Header().Set("Content-Type", "pkglication/json") // Set the Content-Type header of the response to indicate that it contains JSON data
	w.WriteHeader(http.StatusOK)                       // Set the HTTP status code of the response to 200 (OK)
	w.Write(res)                                       // Write the JSON-encoded response to the ResponseWriter
}

func GetBookById(w http.ResponseWriter, r *http.Request) { // Define a function named GetBookById that takes a ResponseWriter and a Request as parameters
	vars := mux.Vars(r)      // Extract the variables from the request URL using mux.Vars and store them in the vars variable
	bookId := vars["bookId"] // Retrieve the value of the "bookId" variable from the URL and store it in the bookId variable

	ID, err := strconv.ParseInt(bookId, 0, 0) // Convert the bookId variable from string to int64 using strconv.ParseInt, and store the result in the ID variable
	if err != nil {                           // Check if there was an error during the conversion
		fmt.Println("error while parsing") // Print an error message to the console if there was an error
	}

	bookDetails, _ := models.GetBookById(ID) // Call the GetBookById function from the models package with the ID variable as the argument and store the result in the bookDetails variable

	res, _ := json.Marshal(bookDetails) // Marshal the bookDetails variable into JSON format and store the result in the res variable

	w.Header().Set("Content-Type", "pkglication/json") // Set the Content-Type header of the response to indicate that it contains JSON data
	w.WriteHeader(http.StatusOK)                       // Set the HTTP status code of the response to 200 (OK)
	w.Write(res)                                       // Write the JSON-encoded response to the ResponseWriter
}

func CreateBook(w http.ResponseWriter, r *http.Request) { // Define a function named CreateBook that takes a ResponseWriter and a Request as parameters
	createBook := &models.Book{}   // Create a new instance of the models.Book struct and store its address in the createBook variable
	utils.ParseBody(r, createBook) // Parse the request body and populate the createBook variable with the parsed data using the ParseBody function from the utils package
	b := createBook.CreateBook()   // Call the CreateBook method on the createBook variable to create a new book record and store the result in the b variable

	res, _ := json.Marshal(b) // Marshal the b variable into JSON format and store the result in the res variable

	w.WriteHeader(http.StatusOK) // Set the HTTP status code of the response to 200 (OK)
	w.Write(res)                 // Write the JSON-encoded response to the ResponseWriter
}

func DeleteBook(w http.ResponseWriter, r *http.Request) { // Define a function named DeleteBook that takes a ResponseWriter and a Request as parameters
	vars := mux.Vars(r)      // Extract the variables from the request URL using mux.Vars and store them in the vars variable
	bookId := vars["bookId"] // Retrieve the value of the "bookId" variable from the URL and store it in the bookId variable

	ID, err := strconv.ParseInt(bookId, 0, 0) // Convert the bookId variable from string to int64 using strconv.ParseInt, and store the result in the ID variable
	if err != nil {                           // Check if there was an error during the conversion
		fmt.Println("error while parsing") // Print an error message to the console if there was an error
	}

	book := models.DeleteBook(ID) // Call the DeleteBook function from the models package with the ID variable as the argument and store the result in the book variable

	res, _ := json.Marshal(book) // Marshal the book variable into JSON format and store the result in the res variable

	w.Header().Set("Content-Type", "pkglication/json") // Set the Content-Type header of the response to indicate that it contains JSON data
	w.WriteHeader(http.StatusOK)                       // Set the HTTP status code of the response to 200 (OK)
	w.Write(res)                                       // Write the JSON-encoded response to the ResponseWriter
}

func UpdateBook(w http.ResponseWriter, r *http.Request) { // Define a function named UpdateBook that takes a ResponseWriter and a Request as parameters
	var updateBook = &models.Book{} // Create a new instance of the models.Book struct and store its address in the updateBook variable
	utils.ParseBody(r, updateBook)  // Parse the request body and populate the updateBook variable with the parsed data using the ParseBody function from the utils package
	vars := mux.Vars(r)             // Extract the variables from the request URL using mux.Vars and store them in the vars variable
	bookId := vars["bookId"]        // Retrieve the value of the "bookId" variable from the URL and store it in the bookId variable

	ID, err := strconv.ParseInt(bookId, 0, 0) // Convert the bookId variable from string to int64 using strconv.ParseInt, and store the result in the ID variable
	if err != nil {                           // Check if there was an error during the conversion
		fmt.Println("Error while parsing") // Print an error message to the console if there was an error
	}

	// Get the book details from the database using the ID
	bookDetails, db := models.GetBookById(ID)

	// Update book details if corresponding fields in the updateBook are not empty
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)                              // Save the updated book details to the database
	res, _ := json.Marshal(bookDetails)                // Marshal the bookDetails variable into JSON format and store the result in the res variable
	w.Header().Set("Content-Type", "pkglication/json") // Set the Content-Type header of the response to indicate that it contains JSON data
	w.WriteHeader(http.StatusOK)                       // Set the HTTP status code of the response to 200 (OK)
	w.Write(res)                                       // Write the JSON-encoded response to the ResponseWriter
}
