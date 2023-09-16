package main

import "net/http"

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello, world!"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("About"))
}


func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8080", nil)
}
