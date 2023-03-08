package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() failed = %v", err)
		return
	}
	fmt.Fprintf(w, "Post request successful \n")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "The Name is = %s\n", name)

	fmt.Fprintf(w, "The Address is = %s\n", address)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "Not Found 404", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "The method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hellooo!")

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at post 8080 \n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
