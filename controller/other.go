package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func OtherHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("other.html")
	if err != nil {
		http.Error(w, "Could not open requested file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", content)
}
