package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	x1 := &messageHandler{"Mesaj 1"}
	x2 := &messageHandler{"Mesaj 2"}

	mux.Handle("/m1", x1)
	mux.Handle("/m2", x2)

	log.Println("Listening Server...")

	http.ListenAndServe(":8000", mux)
}

type messageHandler struct {
	message string
}

func (x *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, x.message)
}
