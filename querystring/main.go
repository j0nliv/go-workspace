package main

import "net/http"

func main() {
	http.HandleFunc("/search", search)
	http.ListenAndServe(":8000", nil)
}

func search(w http.ResponseWriter, r *http.Request) {
	h1 := r.FormValue("h1")
	source := r.FormValue("source")
	ei := r.FormValue("ei")
	q := r.FormValue("q")

	data := "/search?h1=" + h1 + "&source=" + source + "&ei=" + ei + "&q=" + q

	w.Write([]byte(data))

}
