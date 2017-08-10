package main

import (
	"io"
	"net/http"

	"io/ioutil"

	"log"

	"fmt"

	"encoding/json"

	"github.com/kelcecil/shorturl/encoding"
	"github.com/kelcecil/shorturl/storage"
)

// URLHandler ... responsibility is to hold the code
// to add a short URL to our store.
type URLHandler struct {
	storage storage.URLStorage
}

// MakeURLHandler is a factory method to create a
// URL handler.
func MakeURLHandler(storage storage.URLStorage) URLHandler {
	return URLHandler{
		storage: storage,
	}
}

func (h URLHandler) handleGET(response http.ResponseWriter, request *http.Request) {
	key := request.URL.Path[1:]
	log.Printf("Received %v", key)
	id := encoding.HashToID(key)
	url, err := h.storage.Get(id)
	if err != nil {
		message := fmt.Sprintf("Issue retrieving URL from storage. Message: %v", err.Error())
		log.Print(message)
		response.WriteHeader(500)
		io.WriteString(response, message)
		return
	}
	http.Redirect(response, request, url, http.StatusTemporaryRedirect)

}

func (h URLHandler) handlePOST(response http.ResponseWriter, request *http.Request) {
	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		message := fmt.Sprintf("Problem reading request body. Message: %v", err.Error())
		log.Print(message)
		response.WriteHeader(500)
		io.WriteString(response, message)
		return
	}

	var parsedBody AddRequest
	err = json.Unmarshal(requestBody, &parsedBody)
	if err != nil {
		message := fmt.Sprintf("Problem unmarshaling request body. Message: %v", err.Error())
		log.Print(message)
		response.WriteHeader(500)
		io.WriteString(response, message)
		return
	}

	log.Print(parsedBody.URL)

	storageIdentifier, err := h.storage.Add(parsedBody.URL)
	if err != nil {
		message := fmt.Sprintf("Problem writing new URL. Message: %v", err.Error())
		log.Print(message)
		response.WriteHeader(500)
		io.WriteString(response, message)
		return
	}

	key := encoding.IDToHash(storageIdentifier)
	r := MakeAddResponse(key)

	responseBytes, err := json.Marshal(r)
	if err != nil {
		message := fmt.Sprintf("Problem marshaling response. Message: %v", err.Error())
		log.Print(message)
		response.WriteHeader(500)
		io.WriteString(response, message)
		return
	}

	response.WriteHeader(200)
	_, err = response.Write(responseBytes)
	if err != nil {
		message := fmt.Sprintf("Problem sending response. Message: %v", err.Error())
		log.Print(message)
	}
}

// ServeHTTP decides where the request should go based on HTTP action.
func (h URLHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		h.handlePOST(response, request)
	case "GET":
		h.handleGET(response, request)
	default:
		response.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(response, "Invalid method")
	}
}
