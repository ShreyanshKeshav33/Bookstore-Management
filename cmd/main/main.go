package main

import (
	"log" //for logging errors
	"net/http"

	"github.com/ShreyanshKeshav33/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r) //from routes folder's file
	http.Handle("/", r)               //tells the http server to handle requests with our router
	log.Fatal(http.ListenAndServe("localhost:3306", r))
}
