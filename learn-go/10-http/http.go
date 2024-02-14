package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/get", GetUser)
	//handlerFunctions -> Controllers
	http.HandleFunc("/home", Home)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	//ResponseWriter -> writing back to the client
	// Request -> everything that is coming to the server
	_, err := w.Write([]byte("hello this is our first web service"))
	if err != nil {
		// it would signal the error back to the client in the text format
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		// don't forget to return after writing the error
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Define a user with id and hobbies
	user := struct {
		Id      int
		Hobbies []string
	}{
		Id:      1,
		Hobbies: []string{"reading", "coding", "gaming"},
	}

	//// Marshal the user into JSON
	//jsonBytes, err := json.Marshal(user)
	// below line is going to encode the user to json and also write it to the client
	w.WriteHeader(http.StatusAccepted)
	err := json.NewEncoder(w).Encode(user)
	// If there was an error while marshaling
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// If there was no error, respond with 'OK' status
	// Write the JSON data to the response body
	//w.WriteHeader(http.StatusAccepted)
	//w.Write(jsonBytes)

}
