package main

import (
	"net/http"

	"github.com/kelcecil/shorturl/storage"
)

func main() {

	localStorage := storage.MakeMapURLStorage()

	addHandler := MakeURLHandler(localStorage)
	http.Handle("/", addHandler)
	print("Starting server")
	http.ListenAndServe(":8080", nil)
}
