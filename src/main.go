package main

import (
	"fmt"
	"log"
	"net/http"

	"dragonir/src/router"
)

func main() {
	r, err := router.NewRouter()
	if err != nil {
		log.Fatal("Failed to initialize router:", err)
	}

	http.Handle("/", r)
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
