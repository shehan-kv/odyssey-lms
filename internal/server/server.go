package server

import (
	"fmt"
	"log"
	"net/http"
)

func RunApplication() {

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Hello From Odyssey LMS ! </h1>")
	})

	fmt.Print("Starting HTTP Server...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
