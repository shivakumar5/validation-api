package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//LowerCaseLetters - Get Lower case letters from a-z
func LowerCaseLetters() []string {
	var letters []string
	// ASCII values from 97 to 122 i.e. a to z
	for i := 97; i < 123; i++ {
		letters = append(letters, string(i))
	}
	return letters
}

// validateString method to validate the given string and return false if it doesn't satisfy the condition
func validateString(inputstr string) bool {

	validstring := true

	// convert input string to lower case
	inputstr = strings.ToLower(inputstr)

	for _, letters := range LowerCaseLetters() {
		// check whether input string contains all the letters
		validstring = validstring && strings.ContainsAny(inputstr, letters)
	}

	return validstring
}

// handler function ...
func handler(w http.ResponseWriter, r *http.Request) {

	keystr := r.URL.Query().Get("inputstring")
	if keystr == "" {
		w.WriteHeader(http.StatusNoContent)
		log.Println("URL Parameter 'inputstr' is missing")
		return
	}

	// check the input string
	if validateString(keystr) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(true)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(false)
	}
}

func main() {

	newRouter := mux.NewRouter()

	newRouter.HandleFunc("/validate", handler).Methods("GET")
	log.Println("Listening on port : 8080")
	log.Fatal(http.ListenAndServe(":8080", newRouter))

} // End of Main
