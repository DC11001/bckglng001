package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	fmt.Println("Server is running on http://localhost:3000")
	server := http.ListenAndServe(":3000", router)
	log.Fatal(server)
}
