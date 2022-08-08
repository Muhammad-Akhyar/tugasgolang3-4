package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Halaman Utama")
	})
	fmt.Println("startng web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
