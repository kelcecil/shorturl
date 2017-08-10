package main

// AddRequest is the expected structure
// when unmarshaling a request to add a URL.
type AddRequest struct {
	URL string `json:"url"`
}

// AddResponse is the expected structure
// when marshaling a response after adding
// a URL
type AddResponse struct {
	ID string `json:"id"`
}

// MakeAddResponse is a convenience method to
// create the AddResponse model.
func MakeAddResponse(ID string) AddResponse {
	return AddResponse{
		ID: ID,
	}
}
