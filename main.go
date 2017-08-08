package main

import (
	"net/http"
)

func main() {

	addHandler := AddUrlHandler{}
	http.Handle("/add", addHandler)
	print("Starting server")
	http.ListenAndServe(":8080", nil)
}
