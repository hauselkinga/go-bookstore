package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hauselkinga/go-bookstore/pkg/models"
	"github.com/hauselkinga/go-bookstore/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	models.Init()
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}