package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"unicode"

	"github.com/gorilla/mux"
)

// validateString method to validate the given string and return false if it doesn't satisfy the condition
func validateString(key string) bool {

	// Varibales to check the condition
	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, c := range key {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		default:
			return false
		}

	}

	if !hasLower || !hasUpper || !hasNumber || !hasSpecial {
		return false
	}

	return true
}

// handler function ...
func handler(w http.ResponseWriter, r *http.Request) {

	keystr := r.URL.Query().Get("inputstring")
	if keystr == "" {
		w.WriteHeader(http.StatusNoContent)
		log.Println("URL Parameter 'inputstr' is missing")
		return
	}

	if validErrs := validateString(keystr); !validErrs {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(false)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(true)
	}

}

func main() {

	newRouter := mux.NewRouter()

	newRouter.HandleFunc("/validate", handler).Methods("GET")
	fmt.Println("Listening on port : 8080")
	log.Fatal(http.ListenAndServe(":8080", newRouter))

} // End of Main
