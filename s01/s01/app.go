package main

import (
	"fmt"
	"net/http"
	"s01/src/db"
	"s01/src/routes"
)

func main() {

	r := routes.NewRouter()
	println("Starting Bidding server")
	dbFlag := db.InitializeDB()
	if dbFlag == true {
		fmt.Println("DB initialized")
	}
	fmt.Println("Listening on Port Number", 8080)
	http.ListenAndServe(":8080", r)
}
