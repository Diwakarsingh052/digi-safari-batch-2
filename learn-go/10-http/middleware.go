package main

import (
	"fmt"
	"net/http"
)

// Entrypoint -> Mid() -> HandlerFunc -> someOtherLayers
//
//	Mid()   <-
func main() {
	// HandleFunc is a func that register endpoints, and the handler func  which is going to handle the request
	// add a logging middleware in below endpoint
	http.HandleFunc("/home", Mid(LoggingMid(HomePage))) // Mid->loggingMid->homePage
	panic(http.ListenAndServe(":8080", nil))
}

// Middleware is a special type of a func that accepts a handlerfunc and returns a handlerfunc

func LoggingMid(next http.HandlerFunc) http.HandlerFunc {
	// Print the request.Method and the url
	// if the request is not a GET request then return, don't call the homePage handler
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request method: %s, URL: %s\n", r.Method, r.URL)
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed "+r.Method, http.StatusMethodNotAllowed)
			return
		}
		next(w, r)
	}
}

func Mid(next http.HandlerFunc) http.HandlerFunc {
	//we need to return the func(w http.ResponseWriter, r *http.Request) because mid return type
	// is http.Handler func
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("m1 invoked")
		next(w, r) // closure // when an anonymous uses or access the variables outside of
		// the function the body
		fmt.Println("returning from mid")
	}
}

// HomePage is a handler function
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home page invoked")
	fmt.Fprintln(w, "this is my home")
}
