package main

import (
	"goweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HandlerIndex)   //membuat routng index
	mux.HandleFunc("/about", handler.AboutPage) //membuat routng index
	mux.HandleFunc("/profil", handler.ProfilPage)
	mux.HandleFunc("/product", handler.ProductPage)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Println("starting port :8888") //memberikan pesan di dalam terminal
	err := http.ListenAndServe(":8888", mux)
	log.Fatal(err)
}
