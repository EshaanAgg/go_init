package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRoot)
	
	server := http.NewServeMux()
	err := http.ListenAndServe(":3333", server)

	if err != nil {
		fmt.Println("Error while opening the server")
	}

	fmt.Println("The server is running on port 3333.")
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The server is up and running!"))
}