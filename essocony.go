package main

import (
	"fmt"
	"github.com/agent-pink/esso"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.Handle("/", esso.App())
	fmt.Println("Listening on port:", port)
	http.ListenAndServe(":"+port, nil)
}
