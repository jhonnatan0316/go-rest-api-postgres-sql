package routes

import "net/http"

func HomeIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome To my Api Rest example wit Data Base!"))
}
