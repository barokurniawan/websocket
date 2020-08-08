package controller

import "net/http"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 Page not found!", http.StatusNotFound)
}
