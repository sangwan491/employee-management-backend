package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/sangwan491/backend-assignments/employee-management/backend/routes"
)

func main() {
	r := router.SetupRouter()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server started on port 8080")
}
