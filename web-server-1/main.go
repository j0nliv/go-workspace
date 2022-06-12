package main

import "net/http"

func indexHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Home Page!\n"))
	identy := r.URL.Path[1:]
	rw.Write([]byte(identy))
}

func infoHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("My Information!"))
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/info", infoHandler)
	http.ListenAndServe(":8000", nil)
}
