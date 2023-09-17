package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)
	router.HandleFunc("/faq", faq)

	addr := ":8080"
	log.Printf("Starting server on %s", addr)
	http.ListenAndServe(addr, router)

}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>We could not find the page you "+
			"were looking for :(</h1>"+
			"<p>Please email us if you keep being sent to an "+
			"invalid page.</p>")
		return
	}
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}
func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>To get in touch, please send an email "+
		"to <a href=\"mailto:support@lenslocked.com\">"+
		"support@lenslocked.com</a>.</h1>")
}

func faq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>FAQ</h1>")
}
