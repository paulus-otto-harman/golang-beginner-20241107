package main

import (
	"20241107/class/1/router"
	"20241107/database"
	"log"
	"net/http"
)

func main() {
	db := database.DbOpen()
	defer db.Close()
	r := router.NewRouter(db)

	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))

	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)
}
