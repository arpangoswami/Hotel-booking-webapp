package main

import (
	"fmt"
	"net/http"

	"github.com/arpangoswami/Hotel-booking-webapp/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println(fmt.Printf("Listening on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
