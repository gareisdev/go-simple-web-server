package main

import (
	"fmt"
	"net/http"

	"github.com/gareisdev/go-simple-web-server/pkg/handlers"
)

const portNumber = ":8000"

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/products", handlers.Products)
	http.HandleFunc("/store", handlers.Store)
	
	fmt.Printf("Starting app on port %v \n", portNumber)
	http.ListenAndServe(portNumber, nil)

}
